package filemiddleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/pkg/s3wrapper"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func NewTarResponder(next middleware.Responder, fname string, fileNames []string, client s3wrapper.API) middleware.Responder {
	return &tarResponder{
		next:      next,
		fileName:  fname,
		fileNames: fileNames,
		client:    client,
		length:    0,
	}
}

type tarResponder struct {
	next      middleware.Responder
	fileName  string
	length    int64
	client    s3wrapper.API
	fileNames []string
}

func (f *tarResponder) WriteResponse(rw http.ResponseWriter, r runtime.Producer) {
	rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", f.fileName))
	if f.length != 0 {
		rw.Header().Set("Content-Length", strconv.FormatInt(f.length, 10))
	}
	rw.Header().Add("Content-Type", "application/zip")
	_ = common.CreateTar(context.Background(), rw, f.fileNames, f.client)
	//f.next.WriteResponse(rw, r)
}
