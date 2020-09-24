// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// DownloadHostLogsReader is a Reader for the DownloadHostLogs structure.
type DownloadHostLogsReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadHostLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadHostLogsOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDownloadHostLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadHostLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadHostLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDownloadHostLogsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDownloadHostLogsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDownloadHostLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDownloadHostLogsOK creates a DownloadHostLogsOK with default headers values
func NewDownloadHostLogsOK(writer io.Writer) *DownloadHostLogsOK {
	return &DownloadHostLogsOK{
		Payload: writer,
	}
}

/*DownloadHostLogsOK handles this case with default header values.

Success.
*/
type DownloadHostLogsOK struct {
	Payload io.Writer
}

func (o *DownloadHostLogsOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/logs][%d] downloadHostLogsOK  %+v", 200, o.Payload)
}

func (o *DownloadHostLogsOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadHostLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostLogsUnauthorized creates a DownloadHostLogsUnauthorized with default headers values
func NewDownloadHostLogsUnauthorized() *DownloadHostLogsUnauthorized {
	return &DownloadHostLogsUnauthorized{}
}

/*DownloadHostLogsUnauthorized handles this case with default header values.

Unauthorized.
*/
type DownloadHostLogsUnauthorized struct {
	Payload *models.InfraError
}

func (o *DownloadHostLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/logs][%d] downloadHostLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *DownloadHostLogsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DownloadHostLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostLogsForbidden creates a DownloadHostLogsForbidden with default headers values
func NewDownloadHostLogsForbidden() *DownloadHostLogsForbidden {
	return &DownloadHostLogsForbidden{}
}

/*DownloadHostLogsForbidden handles this case with default header values.

Forbidden.
*/
type DownloadHostLogsForbidden struct {
	Payload *models.InfraError
}

func (o *DownloadHostLogsForbidden) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/logs][%d] downloadHostLogsForbidden  %+v", 403, o.Payload)
}

func (o *DownloadHostLogsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DownloadHostLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostLogsNotFound creates a DownloadHostLogsNotFound with default headers values
func NewDownloadHostLogsNotFound() *DownloadHostLogsNotFound {
	return &DownloadHostLogsNotFound{}
}

/*DownloadHostLogsNotFound handles this case with default header values.

Error.
*/
type DownloadHostLogsNotFound struct {
	Payload *models.Error
}

func (o *DownloadHostLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/logs][%d] downloadHostLogsNotFound  %+v", 404, o.Payload)
}

func (o *DownloadHostLogsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadHostLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostLogsMethodNotAllowed creates a DownloadHostLogsMethodNotAllowed with default headers values
func NewDownloadHostLogsMethodNotAllowed() *DownloadHostLogsMethodNotAllowed {
	return &DownloadHostLogsMethodNotAllowed{}
}

/*DownloadHostLogsMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type DownloadHostLogsMethodNotAllowed struct {
	Payload *models.Error
}

func (o *DownloadHostLogsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/logs][%d] downloadHostLogsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *DownloadHostLogsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadHostLogsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostLogsConflict creates a DownloadHostLogsConflict with default headers values
func NewDownloadHostLogsConflict() *DownloadHostLogsConflict {
	return &DownloadHostLogsConflict{}
}

/*DownloadHostLogsConflict handles this case with default header values.

Error.
*/
type DownloadHostLogsConflict struct {
	Payload *models.Error
}

func (o *DownloadHostLogsConflict) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/logs][%d] downloadHostLogsConflict  %+v", 409, o.Payload)
}

func (o *DownloadHostLogsConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadHostLogsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostLogsInternalServerError creates a DownloadHostLogsInternalServerError with default headers values
func NewDownloadHostLogsInternalServerError() *DownloadHostLogsInternalServerError {
	return &DownloadHostLogsInternalServerError{}
}

/*DownloadHostLogsInternalServerError handles this case with default header values.

Error.
*/
type DownloadHostLogsInternalServerError struct {
	Payload *models.Error
}

func (o *DownloadHostLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/logs][%d] downloadHostLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *DownloadHostLogsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadHostLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}