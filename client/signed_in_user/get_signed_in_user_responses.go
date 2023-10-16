// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetSignedInUserReader is a Reader for the GetSignedInUser structure.
type GetSignedInUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSignedInUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSignedInUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSignedInUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSignedInUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSignedInUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSignedInUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /user] getSignedInUser", response, response.Code())
	}
}

// NewGetSignedInUserOK creates a GetSignedInUserOK with default headers values
func NewGetSignedInUserOK() *GetSignedInUserOK {
	return &GetSignedInUserOK{}
}

/*
GetSignedInUserOK describes a response with status code 200, with default header values.

(empty)
*/
type GetSignedInUserOK struct {
	Payload *models.UserProfileDTO
}

// IsSuccess returns true when this get signed in user o k response has a 2xx status code
func (o *GetSignedInUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get signed in user o k response has a 3xx status code
func (o *GetSignedInUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get signed in user o k response has a 4xx status code
func (o *GetSignedInUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get signed in user o k response has a 5xx status code
func (o *GetSignedInUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get signed in user o k response a status code equal to that given
func (o *GetSignedInUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get signed in user o k response
func (o *GetSignedInUserOK) Code() int {
	return 200
}

func (o *GetSignedInUserOK) Error() string {
	return fmt.Sprintf("[GET /user][%d] getSignedInUserOK  %+v", 200, o.Payload)
}

func (o *GetSignedInUserOK) String() string {
	return fmt.Sprintf("[GET /user][%d] getSignedInUserOK  %+v", 200, o.Payload)
}

func (o *GetSignedInUserOK) GetPayload() *models.UserProfileDTO {
	return o.Payload
}

func (o *GetSignedInUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserProfileDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSignedInUserUnauthorized creates a GetSignedInUserUnauthorized with default headers values
func NewGetSignedInUserUnauthorized() *GetSignedInUserUnauthorized {
	return &GetSignedInUserUnauthorized{}
}

/*
GetSignedInUserUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetSignedInUserUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get signed in user unauthorized response has a 2xx status code
func (o *GetSignedInUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get signed in user unauthorized response has a 3xx status code
func (o *GetSignedInUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get signed in user unauthorized response has a 4xx status code
func (o *GetSignedInUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get signed in user unauthorized response has a 5xx status code
func (o *GetSignedInUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get signed in user unauthorized response a status code equal to that given
func (o *GetSignedInUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get signed in user unauthorized response
func (o *GetSignedInUserUnauthorized) Code() int {
	return 401
}

func (o *GetSignedInUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user][%d] getSignedInUserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSignedInUserUnauthorized) String() string {
	return fmt.Sprintf("[GET /user][%d] getSignedInUserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSignedInUserUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSignedInUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSignedInUserForbidden creates a GetSignedInUserForbidden with default headers values
func NewGetSignedInUserForbidden() *GetSignedInUserForbidden {
	return &GetSignedInUserForbidden{}
}

/*
GetSignedInUserForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetSignedInUserForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get signed in user forbidden response has a 2xx status code
func (o *GetSignedInUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get signed in user forbidden response has a 3xx status code
func (o *GetSignedInUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get signed in user forbidden response has a 4xx status code
func (o *GetSignedInUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get signed in user forbidden response has a 5xx status code
func (o *GetSignedInUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get signed in user forbidden response a status code equal to that given
func (o *GetSignedInUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get signed in user forbidden response
func (o *GetSignedInUserForbidden) Code() int {
	return 403
}

func (o *GetSignedInUserForbidden) Error() string {
	return fmt.Sprintf("[GET /user][%d] getSignedInUserForbidden  %+v", 403, o.Payload)
}

func (o *GetSignedInUserForbidden) String() string {
	return fmt.Sprintf("[GET /user][%d] getSignedInUserForbidden  %+v", 403, o.Payload)
}

func (o *GetSignedInUserForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSignedInUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSignedInUserNotFound creates a GetSignedInUserNotFound with default headers values
func NewGetSignedInUserNotFound() *GetSignedInUserNotFound {
	return &GetSignedInUserNotFound{}
}

/*
GetSignedInUserNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetSignedInUserNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get signed in user not found response has a 2xx status code
func (o *GetSignedInUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get signed in user not found response has a 3xx status code
func (o *GetSignedInUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get signed in user not found response has a 4xx status code
func (o *GetSignedInUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get signed in user not found response has a 5xx status code
func (o *GetSignedInUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get signed in user not found response a status code equal to that given
func (o *GetSignedInUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get signed in user not found response
func (o *GetSignedInUserNotFound) Code() int {
	return 404
}

func (o *GetSignedInUserNotFound) Error() string {
	return fmt.Sprintf("[GET /user][%d] getSignedInUserNotFound  %+v", 404, o.Payload)
}

func (o *GetSignedInUserNotFound) String() string {
	return fmt.Sprintf("[GET /user][%d] getSignedInUserNotFound  %+v", 404, o.Payload)
}

func (o *GetSignedInUserNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSignedInUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSignedInUserInternalServerError creates a GetSignedInUserInternalServerError with default headers values
func NewGetSignedInUserInternalServerError() *GetSignedInUserInternalServerError {
	return &GetSignedInUserInternalServerError{}
}

/*
GetSignedInUserInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetSignedInUserInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get signed in user internal server error response has a 2xx status code
func (o *GetSignedInUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get signed in user internal server error response has a 3xx status code
func (o *GetSignedInUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get signed in user internal server error response has a 4xx status code
func (o *GetSignedInUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get signed in user internal server error response has a 5xx status code
func (o *GetSignedInUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get signed in user internal server error response a status code equal to that given
func (o *GetSignedInUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get signed in user internal server error response
func (o *GetSignedInUserInternalServerError) Code() int {
	return 500
}

func (o *GetSignedInUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user][%d] getSignedInUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSignedInUserInternalServerError) String() string {
	return fmt.Sprintf("[GET /user][%d] getSignedInUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSignedInUserInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSignedInUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
