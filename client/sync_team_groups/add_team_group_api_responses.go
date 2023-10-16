// Code generated by go-swagger; DO NOT EDIT.

package sync_team_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// AddTeamGroupAPIReader is a Reader for the AddTeamGroupAPI structure.
type AddTeamGroupAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTeamGroupAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddTeamGroupAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddTeamGroupAPIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddTeamGroupAPIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddTeamGroupAPIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddTeamGroupAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddTeamGroupAPIInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /teams/{teamId}/groups] addTeamGroupApi", response, response.Code())
	}
}

// NewAddTeamGroupAPIOK creates a AddTeamGroupAPIOK with default headers values
func NewAddTeamGroupAPIOK() *AddTeamGroupAPIOK {
	return &AddTeamGroupAPIOK{}
}

/*
AddTeamGroupAPIOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AddTeamGroupAPIOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this add team group Api o k response has a 2xx status code
func (o *AddTeamGroupAPIOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add team group Api o k response has a 3xx status code
func (o *AddTeamGroupAPIOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add team group Api o k response has a 4xx status code
func (o *AddTeamGroupAPIOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add team group Api o k response has a 5xx status code
func (o *AddTeamGroupAPIOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add team group Api o k response a status code equal to that given
func (o *AddTeamGroupAPIOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add team group Api o k response
func (o *AddTeamGroupAPIOK) Code() int {
	return 200
}

func (o *AddTeamGroupAPIOK) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiOK  %+v", 200, o.Payload)
}

func (o *AddTeamGroupAPIOK) String() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiOK  %+v", 200, o.Payload)
}

func (o *AddTeamGroupAPIOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamGroupAPIBadRequest creates a AddTeamGroupAPIBadRequest with default headers values
func NewAddTeamGroupAPIBadRequest() *AddTeamGroupAPIBadRequest {
	return &AddTeamGroupAPIBadRequest{}
}

/*
AddTeamGroupAPIBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AddTeamGroupAPIBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add team group Api bad request response has a 2xx status code
func (o *AddTeamGroupAPIBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add team group Api bad request response has a 3xx status code
func (o *AddTeamGroupAPIBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add team group Api bad request response has a 4xx status code
func (o *AddTeamGroupAPIBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add team group Api bad request response has a 5xx status code
func (o *AddTeamGroupAPIBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add team group Api bad request response a status code equal to that given
func (o *AddTeamGroupAPIBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add team group Api bad request response
func (o *AddTeamGroupAPIBadRequest) Code() int {
	return 400
}

func (o *AddTeamGroupAPIBadRequest) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiBadRequest  %+v", 400, o.Payload)
}

func (o *AddTeamGroupAPIBadRequest) String() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiBadRequest  %+v", 400, o.Payload)
}

func (o *AddTeamGroupAPIBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamGroupAPIUnauthorized creates a AddTeamGroupAPIUnauthorized with default headers values
func NewAddTeamGroupAPIUnauthorized() *AddTeamGroupAPIUnauthorized {
	return &AddTeamGroupAPIUnauthorized{}
}

/*
AddTeamGroupAPIUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AddTeamGroupAPIUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add team group Api unauthorized response has a 2xx status code
func (o *AddTeamGroupAPIUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add team group Api unauthorized response has a 3xx status code
func (o *AddTeamGroupAPIUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add team group Api unauthorized response has a 4xx status code
func (o *AddTeamGroupAPIUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add team group Api unauthorized response has a 5xx status code
func (o *AddTeamGroupAPIUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add team group Api unauthorized response a status code equal to that given
func (o *AddTeamGroupAPIUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add team group Api unauthorized response
func (o *AddTeamGroupAPIUnauthorized) Code() int {
	return 401
}

func (o *AddTeamGroupAPIUnauthorized) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiUnauthorized  %+v", 401, o.Payload)
}

func (o *AddTeamGroupAPIUnauthorized) String() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiUnauthorized  %+v", 401, o.Payload)
}

func (o *AddTeamGroupAPIUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamGroupAPIForbidden creates a AddTeamGroupAPIForbidden with default headers values
func NewAddTeamGroupAPIForbidden() *AddTeamGroupAPIForbidden {
	return &AddTeamGroupAPIForbidden{}
}

/*
AddTeamGroupAPIForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AddTeamGroupAPIForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add team group Api forbidden response has a 2xx status code
func (o *AddTeamGroupAPIForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add team group Api forbidden response has a 3xx status code
func (o *AddTeamGroupAPIForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add team group Api forbidden response has a 4xx status code
func (o *AddTeamGroupAPIForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add team group Api forbidden response has a 5xx status code
func (o *AddTeamGroupAPIForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add team group Api forbidden response a status code equal to that given
func (o *AddTeamGroupAPIForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the add team group Api forbidden response
func (o *AddTeamGroupAPIForbidden) Code() int {
	return 403
}

func (o *AddTeamGroupAPIForbidden) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiForbidden  %+v", 403, o.Payload)
}

func (o *AddTeamGroupAPIForbidden) String() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiForbidden  %+v", 403, o.Payload)
}

func (o *AddTeamGroupAPIForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamGroupAPINotFound creates a AddTeamGroupAPINotFound with default headers values
func NewAddTeamGroupAPINotFound() *AddTeamGroupAPINotFound {
	return &AddTeamGroupAPINotFound{}
}

/*
AddTeamGroupAPINotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type AddTeamGroupAPINotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add team group Api not found response has a 2xx status code
func (o *AddTeamGroupAPINotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add team group Api not found response has a 3xx status code
func (o *AddTeamGroupAPINotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add team group Api not found response has a 4xx status code
func (o *AddTeamGroupAPINotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add team group Api not found response has a 5xx status code
func (o *AddTeamGroupAPINotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add team group Api not found response a status code equal to that given
func (o *AddTeamGroupAPINotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add team group Api not found response
func (o *AddTeamGroupAPINotFound) Code() int {
	return 404
}

func (o *AddTeamGroupAPINotFound) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiNotFound  %+v", 404, o.Payload)
}

func (o *AddTeamGroupAPINotFound) String() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiNotFound  %+v", 404, o.Payload)
}

func (o *AddTeamGroupAPINotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamGroupAPIInternalServerError creates a AddTeamGroupAPIInternalServerError with default headers values
func NewAddTeamGroupAPIInternalServerError() *AddTeamGroupAPIInternalServerError {
	return &AddTeamGroupAPIInternalServerError{}
}

/*
AddTeamGroupAPIInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AddTeamGroupAPIInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add team group Api internal server error response has a 2xx status code
func (o *AddTeamGroupAPIInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add team group Api internal server error response has a 3xx status code
func (o *AddTeamGroupAPIInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add team group Api internal server error response has a 4xx status code
func (o *AddTeamGroupAPIInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add team group Api internal server error response has a 5xx status code
func (o *AddTeamGroupAPIInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add team group Api internal server error response a status code equal to that given
func (o *AddTeamGroupAPIInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add team group Api internal server error response
func (o *AddTeamGroupAPIInternalServerError) Code() int {
	return 500
}

func (o *AddTeamGroupAPIInternalServerError) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiInternalServerError  %+v", 500, o.Payload)
}

func (o *AddTeamGroupAPIInternalServerError) String() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiInternalServerError  %+v", 500, o.Payload)
}

func (o *AddTeamGroupAPIInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPIInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
