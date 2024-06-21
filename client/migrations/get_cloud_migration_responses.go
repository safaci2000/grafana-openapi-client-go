// Code generated by go-swagger; DO NOT EDIT.

package migrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetCloudMigrationReader is a Reader for the GetCloudMigration structure.
type GetCloudMigrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudMigrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudMigrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCloudMigrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCloudMigrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCloudMigrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloudmigration/migration/{uid}] getCloudMigration", response, response.Code())
	}
}

// NewGetCloudMigrationOK creates a GetCloudMigrationOK with default headers values
func NewGetCloudMigrationOK() *GetCloudMigrationOK {
	return &GetCloudMigrationOK{}
}

/*
GetCloudMigrationOK describes a response with status code 200, with default header values.

(empty)
*/
type GetCloudMigrationOK struct {
	Payload *models.CloudMigrationResponse
}

// IsSuccess returns true when this get cloud migration Ok response has a 2xx status code
func (o *GetCloudMigrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cloud migration Ok response has a 3xx status code
func (o *GetCloudMigrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud migration Ok response has a 4xx status code
func (o *GetCloudMigrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cloud migration Ok response has a 5xx status code
func (o *GetCloudMigrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud migration Ok response a status code equal to that given
func (o *GetCloudMigrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cloud migration Ok response
func (o *GetCloudMigrationOK) Code() int {
	return 200
}

func (o *GetCloudMigrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{uid}][%d] getCloudMigrationOk %s", 200, payload)
}

func (o *GetCloudMigrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{uid}][%d] getCloudMigrationOk %s", 200, payload)
}

func (o *GetCloudMigrationOK) GetPayload() *models.CloudMigrationResponse {
	return o.Payload
}

func (o *GetCloudMigrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudMigrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudMigrationUnauthorized creates a GetCloudMigrationUnauthorized with default headers values
func NewGetCloudMigrationUnauthorized() *GetCloudMigrationUnauthorized {
	return &GetCloudMigrationUnauthorized{}
}

/*
GetCloudMigrationUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetCloudMigrationUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get cloud migration unauthorized response has a 2xx status code
func (o *GetCloudMigrationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud migration unauthorized response has a 3xx status code
func (o *GetCloudMigrationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud migration unauthorized response has a 4xx status code
func (o *GetCloudMigrationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud migration unauthorized response has a 5xx status code
func (o *GetCloudMigrationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud migration unauthorized response a status code equal to that given
func (o *GetCloudMigrationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get cloud migration unauthorized response
func (o *GetCloudMigrationUnauthorized) Code() int {
	return 401
}

func (o *GetCloudMigrationUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{uid}][%d] getCloudMigrationUnauthorized %s", 401, payload)
}

func (o *GetCloudMigrationUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{uid}][%d] getCloudMigrationUnauthorized %s", 401, payload)
}

func (o *GetCloudMigrationUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCloudMigrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudMigrationForbidden creates a GetCloudMigrationForbidden with default headers values
func NewGetCloudMigrationForbidden() *GetCloudMigrationForbidden {
	return &GetCloudMigrationForbidden{}
}

/*
GetCloudMigrationForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetCloudMigrationForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get cloud migration forbidden response has a 2xx status code
func (o *GetCloudMigrationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud migration forbidden response has a 3xx status code
func (o *GetCloudMigrationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud migration forbidden response has a 4xx status code
func (o *GetCloudMigrationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud migration forbidden response has a 5xx status code
func (o *GetCloudMigrationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud migration forbidden response a status code equal to that given
func (o *GetCloudMigrationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cloud migration forbidden response
func (o *GetCloudMigrationForbidden) Code() int {
	return 403
}

func (o *GetCloudMigrationForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{uid}][%d] getCloudMigrationForbidden %s", 403, payload)
}

func (o *GetCloudMigrationForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{uid}][%d] getCloudMigrationForbidden %s", 403, payload)
}

func (o *GetCloudMigrationForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCloudMigrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudMigrationInternalServerError creates a GetCloudMigrationInternalServerError with default headers values
func NewGetCloudMigrationInternalServerError() *GetCloudMigrationInternalServerError {
	return &GetCloudMigrationInternalServerError{}
}

/*
GetCloudMigrationInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetCloudMigrationInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get cloud migration internal server error response has a 2xx status code
func (o *GetCloudMigrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud migration internal server error response has a 3xx status code
func (o *GetCloudMigrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud migration internal server error response has a 4xx status code
func (o *GetCloudMigrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cloud migration internal server error response has a 5xx status code
func (o *GetCloudMigrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cloud migration internal server error response a status code equal to that given
func (o *GetCloudMigrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cloud migration internal server error response
func (o *GetCloudMigrationInternalServerError) Code() int {
	return 500
}

func (o *GetCloudMigrationInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{uid}][%d] getCloudMigrationInternalServerError %s", 500, payload)
}

func (o *GetCloudMigrationInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{uid}][%d] getCloudMigrationInternalServerError %s", 500, payload)
}

func (o *GetCloudMigrationInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCloudMigrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
