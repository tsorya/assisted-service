package common

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/openshift/assisted-service/pkg/s3wrapper"
	"github.com/pkg/errors"
)

const MinMasterHostsNeededForInstallation = 3

// continueOnError is set when running as stream, error is doing nothing when it happens cause we in the middle of stream
// and 200 was already returned
func CreateTar(ctx context.Context, w io.Writer, files []string, client s3wrapper.API, continueOnError bool) error {
	var rdr io.ReadCloser
	tarWriter := tar.NewWriter(w)
	defer func() {
		tarWriter.Close()
		if rdr != nil {
			rdr.Close()
		}
	}()
	var err error
	var objectSize int64

	// Create tar headers from s3 files
	for _, file := range files {
		fmt.Println("5555555555555555555555555555555555")
		// Read file from S3, log any errors
		rdr, objectSize, err = client.Download(ctx, file)
		if err != nil {
			if continueOnError {
				continue
			}
			fmt.Println(err)
			return errors.Wrapf(err, "Failed to open reader for %s", file)
		}

		header := tar.Header{
			Name: file,
			Size: objectSize,
		}
		err = tarWriter.WriteHeader(&header)
		fmt.Println("888888888888888888888888888888")
		if err != nil && !continueOnError {
			fmt.Println("99999999999999999999999")
			fmt.Println(err)
			return errors.Wrapf(err, "Failed to write tar header with file %s details", file)
		}
		_, err = io.Copy(tarWriter, rdr)
		if err != nil && !continueOnError {
			fmt.Println(err)
			return errors.Wrapf(err, "Failed to write file %s to tar", file)
		}
		fmt.Println("6666666666666666666666666666666666")
		_ = rdr.Close()
		fmt.Println("7777777777777777777777777777")
	}

	return nil
}

// Tar given files in s3 bucket.
// We open pipe for reading from aws and writing archived back to it while zipping them.
// It creates stream by using io.pipe
func TarAwsFiles(ctx context.Context, tarName string, files []string, client s3wrapper.API, log logrus.FieldLogger) error {
	// Create pipe
	var err error
	pr, pw := io.Pipe()
	// Create zip.Write which will writes to pipes
	wg := sync.WaitGroup{}
	// Wait for downloader and uploader
	wg.Add(2)
	// Run 'downloader'
	go func() {
		defer func() {
			wg.Done()
			fmt.Println("CCCCCCCCCCCCCCCCCCC")
			// closing pipe will stop uploading
			pw.Close()
		}()
		fmt.Println("11111111111111111111111111111111111111")
		downloadError := CreateTar(ctx, pw, files, client, false)
		if downloadError != nil && err == nil {
			err = errors.Wrapf(downloadError, "Failed to download files while creating zip %s", tarName)
			log.Error(err)
		}
	}()
	go func() {
		defer func() {
			wg.Done()
			// if upload fails close pipe
			// will fail download too
			fmt.Println("BBBBBBBBBBBBBBBB")
			pr.Close()
		}()
		// Upload the file, body is `io.Reader` from pipe
		time.Sleep(20 *time.Second)
		uploadError := client.UploadStream(ctx, pr, tarName)
		fmt.Println("BBBBBBBBBBBBBBBB ", uploadError)
		if uploadError != nil && err == nil {
			err = errors.Wrapf(uploadError, "Failed to upload zip %s", tarName)
			log.Error(err)
		}
	}()
	wg.Wait()
	return err
}
