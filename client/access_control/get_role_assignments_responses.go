// Code generated by go-swagger; DO NOT EDIT.

package access_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetRoleAssignmentsReader is a Reader for the GetRoleAssignments structure.
type GetRoleAssignmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleAssignmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleAssignmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRoleAssignmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoleAssignmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoleAssignmentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /access-control/roles/{roleUID}/assignments] getRoleAssignments", response, response.Code())
	}
}

// NewGetRoleAssignmentsOK creates a GetRoleAssignmentsOK with default headers values
func NewGetRoleAssignmentsOK() *GetRoleAssignmentsOK {
	return &GetRoleAssignmentsOK{}
}

/*
GetRoleAssignmentsOK describes a response with status code 200, with default header values.

(empty)
*/
type GetRoleAssignmentsOK struct {
	Payload *models.RoleAssignmentsDTO
}

// IsSuccess returns true when this get role assignments o k response has a 2xx status code
func (o *GetRoleAssignmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get role assignments o k response has a 3xx status code
func (o *GetRoleAssignmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role assignments o k response has a 4xx status code
func (o *GetRoleAssignmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get role assignments o k response has a 5xx status code
func (o *GetRoleAssignmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get role assignments o k response a status code equal to that given
func (o *GetRoleAssignmentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get role assignments o k response
func (o *GetRoleAssignmentsOK) Code() int {
	return 200
}

func (o *GetRoleAssignmentsOK) Error() string {
	return fmt.Sprintf("[GET /access-control/roles/{roleUID}/assignments][%d] getRoleAssignmentsOK  %+v", 200, o.Payload)
}

func (o *GetRoleAssignmentsOK) String() string {
	return fmt.Sprintf("[GET /access-control/roles/{roleUID}/assignments][%d] getRoleAssignmentsOK  %+v", 200, o.Payload)
}

func (o *GetRoleAssignmentsOK) GetPayload() *models.RoleAssignmentsDTO {
	return o.Payload
}

func (o *GetRoleAssignmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleAssignmentsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleAssignmentsForbidden creates a GetRoleAssignmentsForbidden with default headers values
func NewGetRoleAssignmentsForbidden() *GetRoleAssignmentsForbidden {
	return &GetRoleAssignmentsForbidden{}
}

/*
GetRoleAssignmentsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetRoleAssignmentsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get role assignments forbidden response has a 2xx status code
func (o *GetRoleAssignmentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role assignments forbidden response has a 3xx status code
func (o *GetRoleAssignmentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role assignments forbidden response has a 4xx status code
func (o *GetRoleAssignmentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get role assignments forbidden response has a 5xx status code
func (o *GetRoleAssignmentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get role assignments forbidden response a status code equal to that given
func (o *GetRoleAssignmentsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get role assignments forbidden response
func (o *GetRoleAssignmentsForbidden) Code() int {
	return 403
}

func (o *GetRoleAssignmentsForbidden) Error() string {
	return fmt.Sprintf("[GET /access-control/roles/{roleUID}/assignments][%d] getRoleAssignmentsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoleAssignmentsForbidden) String() string {
	return fmt.Sprintf("[GET /access-control/roles/{roleUID}/assignments][%d] getRoleAssignmentsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoleAssignmentsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetRoleAssignmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleAssignmentsNotFound creates a GetRoleAssignmentsNotFound with default headers values
func NewGetRoleAssignmentsNotFound() *GetRoleAssignmentsNotFound {
	return &GetRoleAssignmentsNotFound{}
}

/*
GetRoleAssignmentsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetRoleAssignmentsNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get role assignments not found response has a 2xx status code
func (o *GetRoleAssignmentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role assignments not found response has a 3xx status code
func (o *GetRoleAssignmentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role assignments not found response has a 4xx status code
func (o *GetRoleAssignmentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get role assignments not found response has a 5xx status code
func (o *GetRoleAssignmentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get role assignments not found response a status code equal to that given
func (o *GetRoleAssignmentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get role assignments not found response
func (o *GetRoleAssignmentsNotFound) Code() int {
	return 404
}

func (o *GetRoleAssignmentsNotFound) Error() string {
	return fmt.Sprintf("[GET /access-control/roles/{roleUID}/assignments][%d] getRoleAssignmentsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoleAssignmentsNotFound) String() string {
	return fmt.Sprintf("[GET /access-control/roles/{roleUID}/assignments][%d] getRoleAssignmentsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoleAssignmentsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetRoleAssignmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleAssignmentsInternalServerError creates a GetRoleAssignmentsInternalServerError with default headers values
func NewGetRoleAssignmentsInternalServerError() *GetRoleAssignmentsInternalServerError {
	return &GetRoleAssignmentsInternalServerError{}
}

/*
GetRoleAssignmentsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetRoleAssignmentsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get role assignments internal server error response has a 2xx status code
func (o *GetRoleAssignmentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role assignments internal server error response has a 3xx status code
func (o *GetRoleAssignmentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role assignments internal server error response has a 4xx status code
func (o *GetRoleAssignmentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get role assignments internal server error response has a 5xx status code
func (o *GetRoleAssignmentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get role assignments internal server error response a status code equal to that given
func (o *GetRoleAssignmentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get role assignments internal server error response
func (o *GetRoleAssignmentsInternalServerError) Code() int {
	return 500
}

func (o *GetRoleAssignmentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /access-control/roles/{roleUID}/assignments][%d] getRoleAssignmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoleAssignmentsInternalServerError) String() string {
	return fmt.Sprintf("[GET /access-control/roles/{roleUID}/assignments][%d] getRoleAssignmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoleAssignmentsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetRoleAssignmentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
