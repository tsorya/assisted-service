package common

import (
	"archive/tar"
	"context"
	"io"

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
		// Read file from S3, log any errors
		rdr, objectSize, err = client.Download(ctx, file)
		if err != nil {
			if continueOnError {
				continue
			}
			return errors.Wrapf(err, "Failed to open reader for %s", file)
		}

		header := tar.Header{
			Name: file,
			Size: objectSize,
		}
		err = tarWriter.WriteHeader(&header)
		if err != nil && !continueOnError {
			return errors.Wrapf(err, "Failed to write tar header with file %s details", file)
		}
		_, err = io.Copy(tarWriter, rdr)
		if err != nil && !continueOnError {
			return errors.Wrapf(err, "Failed to write file %s to tar", file)
		}
		_ = rdr.Close()
	}

	return nil
}
