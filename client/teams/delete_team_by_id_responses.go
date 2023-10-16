// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// DeleteTeamByIDReader is a Reader for the DeleteTeamByID structure.
type DeleteTeamByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeamByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTeamByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteTeamByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTeamByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTeamByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTeamByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /teams/{team_id}] deleteTeamByID", response, response.Code())
	}
}

// NewDeleteTeamByIDOK creates a DeleteTeamByIDOK with default headers values
func NewDeleteTeamByIDOK() *DeleteTeamByIDOK {
	return &DeleteTeamByIDOK{}
}

/*
DeleteTeamByIDOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type DeleteTeamByIDOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this delete team by Id o k response has a 2xx status code
func (o *DeleteTeamByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete team by Id o k response has a 3xx status code
func (o *DeleteTeamByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team by Id o k response has a 4xx status code
func (o *DeleteTeamByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team by Id o k response has a 5xx status code
func (o *DeleteTeamByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team by Id o k response a status code equal to that given
func (o *DeleteTeamByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete team by Id o k response
func (o *DeleteTeamByIDOK) Code() int {
	return 200
}

func (o *DeleteTeamByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamByIDOK) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamByIDOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *DeleteTeamByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamByIDUnauthorized creates a DeleteTeamByIDUnauthorized with default headers values
func NewDeleteTeamByIDUnauthorized() *DeleteTeamByIDUnauthorized {
	return &DeleteTeamByIDUnauthorized{}
}

/*
DeleteTeamByIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteTeamByIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete team by Id unauthorized response has a 2xx status code
func (o *DeleteTeamByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team by Id unauthorized response has a 3xx status code
func (o *DeleteTeamByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team by Id unauthorized response has a 4xx status code
func (o *DeleteTeamByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team by Id unauthorized response has a 5xx status code
func (o *DeleteTeamByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team by Id unauthorized response a status code equal to that given
func (o *DeleteTeamByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete team by Id unauthorized response
func (o *DeleteTeamByIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteTeamByIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTeamByIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTeamByIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteTeamByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamByIDForbidden creates a DeleteTeamByIDForbidden with default headers values
func NewDeleteTeamByIDForbidden() *DeleteTeamByIDForbidden {
	return &DeleteTeamByIDForbidden{}
}

/*
DeleteTeamByIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteTeamByIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete team by Id forbidden response has a 2xx status code
func (o *DeleteTeamByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team by Id forbidden response has a 3xx status code
func (o *DeleteTeamByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team by Id forbidden response has a 4xx status code
func (o *DeleteTeamByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team by Id forbidden response has a 5xx status code
func (o *DeleteTeamByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team by Id forbidden response a status code equal to that given
func (o *DeleteTeamByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete team by Id forbidden response
func (o *DeleteTeamByIDForbidden) Code() int {
	return 403
}

func (o *DeleteTeamByIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamByIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamByIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamByIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamByIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteTeamByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamByIDNotFound creates a DeleteTeamByIDNotFound with default headers values
func NewDeleteTeamByIDNotFound() *DeleteTeamByIDNotFound {
	return &DeleteTeamByIDNotFound{}
}

/*
DeleteTeamByIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteTeamByIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete team by Id not found response has a 2xx status code
func (o *DeleteTeamByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team by Id not found response has a 3xx status code
func (o *DeleteTeamByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team by Id not found response has a 4xx status code
func (o *DeleteTeamByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team by Id not found response has a 5xx status code
func (o *DeleteTeamByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team by Id not found response a status code equal to that given
func (o *DeleteTeamByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete team by Id not found response
func (o *DeleteTeamByIDNotFound) Code() int {
	return 404
}

func (o *DeleteTeamByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamByIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamByIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamByIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamByIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteTeamByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamByIDInternalServerError creates a DeleteTeamByIDInternalServerError with default headers values
func NewDeleteTeamByIDInternalServerError() *DeleteTeamByIDInternalServerError {
	return &DeleteTeamByIDInternalServerError{}
}

/*
DeleteTeamByIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteTeamByIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete team by Id internal server error response has a 2xx status code
func (o *DeleteTeamByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team by Id internal server error response has a 3xx status code
func (o *DeleteTeamByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team by Id internal server error response has a 4xx status code
func (o *DeleteTeamByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team by Id internal server error response has a 5xx status code
func (o *DeleteTeamByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete team by Id internal server error response a status code equal to that given
func (o *DeleteTeamByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete team by Id internal server error response
func (o *DeleteTeamByIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteTeamByIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTeamByIDInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTeamByIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteTeamByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
