package filemiddleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/pkg/s3wrapper"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// Responder that will get list of files and stream them from s3 or any other s3wrapper.API option and tar them while streaming
func NewTarResponder(ctx context.Context, next middleware.Responder, fname string, fileNames []string, client s3wrapper.API) middleware.Responder {
	return &tarResponder{
		ctx:       ctx,
		next:      next,
		fileName:  fname,
		fileNames: fileNames,
		client:    client,
	}
}

type tarResponder struct {
	ctx       context.Context
	next      middleware.Responder
	fileName  string
	client    s3wrapper.API
	fileNames []string
}

// Can't return content length cause files will be read only while streaming, so we can't predict their length.
// Even if we will get objects details it will be very hard to predict tar header length
// for now didn't find the way to do it
func (f *tarResponder) WriteResponse(rw http.ResponseWriter, r runtime.Producer) {
	rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", f.fileName))
	_ = common.CreateTar(f.ctx, rw, f.fileNames, f.client, true)
}
