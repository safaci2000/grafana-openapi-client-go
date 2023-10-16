// Code generated by go-swagger; DO NOT EDIT.

package correlations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// CreateCorrelationReader is a Reader for the CreateCorrelation structure.
type CreateCorrelationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCorrelationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCorrelationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCorrelationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateCorrelationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateCorrelationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCorrelationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateCorrelationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /datasources/uid/{sourceUID}/correlations] createCorrelation", response, response.Code())
	}
}

// NewCreateCorrelationOK creates a CreateCorrelationOK with default headers values
func NewCreateCorrelationOK() *CreateCorrelationOK {
	return &CreateCorrelationOK{}
}

/*
CreateCorrelationOK describes a response with status code 200, with default header values.

(empty)
*/
type CreateCorrelationOK struct {
	Payload *models.CreateCorrelationResponseBody
}

// IsSuccess returns true when this create correlation o k response has a 2xx status code
func (o *CreateCorrelationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create correlation o k response has a 3xx status code
func (o *CreateCorrelationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create correlation o k response has a 4xx status code
func (o *CreateCorrelationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create correlation o k response has a 5xx status code
func (o *CreateCorrelationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create correlation o k response a status code equal to that given
func (o *CreateCorrelationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create correlation o k response
func (o *CreateCorrelationOK) Code() int {
	return 200
}

func (o *CreateCorrelationOK) Error() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationOK  %+v", 200, o.Payload)
}

func (o *CreateCorrelationOK) String() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationOK  %+v", 200, o.Payload)
}

func (o *CreateCorrelationOK) GetPayload() *models.CreateCorrelationResponseBody {
	return o.Payload
}

func (o *CreateCorrelationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateCorrelationResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCorrelationBadRequest creates a CreateCorrelationBadRequest with default headers values
func NewCreateCorrelationBadRequest() *CreateCorrelationBadRequest {
	return &CreateCorrelationBadRequest{}
}

/*
CreateCorrelationBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type CreateCorrelationBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create correlation bad request response has a 2xx status code
func (o *CreateCorrelationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create correlation bad request response has a 3xx status code
func (o *CreateCorrelationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create correlation bad request response has a 4xx status code
func (o *CreateCorrelationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create correlation bad request response has a 5xx status code
func (o *CreateCorrelationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create correlation bad request response a status code equal to that given
func (o *CreateCorrelationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create correlation bad request response
func (o *CreateCorrelationBadRequest) Code() int {
	return 400
}

func (o *CreateCorrelationBadRequest) Error() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCorrelationBadRequest) String() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCorrelationBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateCorrelationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCorrelationUnauthorized creates a CreateCorrelationUnauthorized with default headers values
func NewCreateCorrelationUnauthorized() *CreateCorrelationUnauthorized {
	return &CreateCorrelationUnauthorized{}
}

/*
CreateCorrelationUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CreateCorrelationUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create correlation unauthorized response has a 2xx status code
func (o *CreateCorrelationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create correlation unauthorized response has a 3xx status code
func (o *CreateCorrelationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create correlation unauthorized response has a 4xx status code
func (o *CreateCorrelationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create correlation unauthorized response has a 5xx status code
func (o *CreateCorrelationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create correlation unauthorized response a status code equal to that given
func (o *CreateCorrelationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create correlation unauthorized response
func (o *CreateCorrelationUnauthorized) Code() int {
	return 401
}

func (o *CreateCorrelationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateCorrelationUnauthorized) String() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateCorrelationUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateCorrelationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCorrelationForbidden creates a CreateCorrelationForbidden with default headers values
func NewCreateCorrelationForbidden() *CreateCorrelationForbidden {
	return &CreateCorrelationForbidden{}
}

/*
CreateCorrelationForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CreateCorrelationForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create correlation forbidden response has a 2xx status code
func (o *CreateCorrelationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create correlation forbidden response has a 3xx status code
func (o *CreateCorrelationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create correlation forbidden response has a 4xx status code
func (o *CreateCorrelationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create correlation forbidden response has a 5xx status code
func (o *CreateCorrelationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create correlation forbidden response a status code equal to that given
func (o *CreateCorrelationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create correlation forbidden response
func (o *CreateCorrelationForbidden) Code() int {
	return 403
}

func (o *CreateCorrelationForbidden) Error() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationForbidden  %+v", 403, o.Payload)
}

func (o *CreateCorrelationForbidden) String() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationForbidden  %+v", 403, o.Payload)
}

func (o *CreateCorrelationForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateCorrelationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCorrelationNotFound creates a CreateCorrelationNotFound with default headers values
func NewCreateCorrelationNotFound() *CreateCorrelationNotFound {
	return &CreateCorrelationNotFound{}
}

/*
CreateCorrelationNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type CreateCorrelationNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create correlation not found response has a 2xx status code
func (o *CreateCorrelationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create correlation not found response has a 3xx status code
func (o *CreateCorrelationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create correlation not found response has a 4xx status code
func (o *CreateCorrelationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create correlation not found response has a 5xx status code
func (o *CreateCorrelationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create correlation not found response a status code equal to that given
func (o *CreateCorrelationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create correlation not found response
func (o *CreateCorrelationNotFound) Code() int {
	return 404
}

func (o *CreateCorrelationNotFound) Error() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationNotFound  %+v", 404, o.Payload)
}

func (o *CreateCorrelationNotFound) String() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationNotFound  %+v", 404, o.Payload)
}

func (o *CreateCorrelationNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateCorrelationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCorrelationInternalServerError creates a CreateCorrelationInternalServerError with default headers values
func NewCreateCorrelationInternalServerError() *CreateCorrelationInternalServerError {
	return &CreateCorrelationInternalServerError{}
}

/*
CreateCorrelationInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CreateCorrelationInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create correlation internal server error response has a 2xx status code
func (o *CreateCorrelationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create correlation internal server error response has a 3xx status code
func (o *CreateCorrelationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create correlation internal server error response has a 4xx status code
func (o *CreateCorrelationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create correlation internal server error response has a 5xx status code
func (o *CreateCorrelationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create correlation internal server error response a status code equal to that given
func (o *CreateCorrelationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create correlation internal server error response
func (o *CreateCorrelationInternalServerError) Code() int {
	return 500
}

func (o *CreateCorrelationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCorrelationInternalServerError) String() string {
	return fmt.Sprintf("[POST /datasources/uid/{sourceUID}/correlations][%d] createCorrelationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCorrelationInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateCorrelationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
