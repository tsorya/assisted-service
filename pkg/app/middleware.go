package app

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/openshift/assisted-service/pkg/s3wrapper"
	"github.com/thoas/go-funk"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
)

// WithMetricsResponderMiddleware Returns middleware which responds to /metrics endpoint with the prometheus metrics
// of the service
func WithMetricsResponderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/metrics" {
			promhttp.Handler().ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// WithHealthMiddleware returns middleware which responds to the /health endpoint
func WithHealthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/health" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetupCORSMiddleware(handler http.Handler, domains []string) http.Handler {
	corsHandler := cors.New(cors.Options{
		Debug: false,
		AllowedMethods: []string{
			http.MethodDelete,
			http.MethodGet,
			http.MethodPatch,
			http.MethodPost,
			http.MethodPut,
		},
		AllowedOrigins: domains,
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		MaxAge: int((10 * time.Minute).Seconds()),
	})
	return corsHandler.Handler(handler)
}

func WithLogs(handler http.Handler, objectHandler s3wrapper.API) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet && r.URL.Path == "/logs" {
			w.WriteHeader(http.StatusOK)
			// Start processing the response
			w.Header().Add("Content-Disposition", "attachment; filename=\""+"aaaaaaaaaaa.zip"+"\"")
			w.Header().Add("Content-Type", "application/zip")

			files, _ := objectHandler.ListObjectsByPrefix(context.Background(), fmt.Sprintf("%s/logs/", "9dddb7f1-998b-4e94-bf7e-76dc2fff686e"))

			// Loop over files, add them to the
			zipWriter := zip.NewWriter(w)
			var rdr io.ReadCloser
			for _, file := range files {

				if funk.Contains(file, "all_logs") {
					continue
				}
				// Read file from S3, log any errors
				rdr, _, _ = objectHandler.Download(context.Background(), file)

				// We have to set a special flag so zip files recognize utf file names
				// See http://stackoverflow.com/questions/30026083/creating-a-zip-archive-with-unicode-filenames-using-gos-archive-zip
				h := &zip.FileHeader{
					Name:   file,
					Method: zip.Deflate,
					Flags:  0x800,
				}
				f, _ := zipWriter.CreateHeader(h)

				_, _ = io.Copy(f, rdr)
				_ = rdr.Close()
			}
			_ = zipWriter.Close()

			zipWriter.Close()
			return
		}
		handler.ServeHTTP(w, r)
	})
}
