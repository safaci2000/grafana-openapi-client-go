// Code generated by go-swagger; DO NOT EDIT.

package orgs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// UpdateOrgAddressReader is a Reader for the UpdateOrgAddress structure.
type UpdateOrgAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrgAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrgAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOrgAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOrgAddressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOrgAddressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateOrgAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /orgs/{org_id}/address] updateOrgAddress", response, response.Code())
	}
}

// NewUpdateOrgAddressOK creates a UpdateOrgAddressOK with default headers values
func NewUpdateOrgAddressOK() *UpdateOrgAddressOK {
	return &UpdateOrgAddressOK{}
}

/*
UpdateOrgAddressOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateOrgAddressOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this update org address o k response has a 2xx status code
func (o *UpdateOrgAddressOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update org address o k response has a 3xx status code
func (o *UpdateOrgAddressOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org address o k response has a 4xx status code
func (o *UpdateOrgAddressOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update org address o k response has a 5xx status code
func (o *UpdateOrgAddressOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update org address o k response a status code equal to that given
func (o *UpdateOrgAddressOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update org address o k response
func (o *UpdateOrgAddressOK) Code() int {
	return 200
}

func (o *UpdateOrgAddressOK) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org_id}/address][%d] updateOrgAddressOK  %+v", 200, o.Payload)
}

func (o *UpdateOrgAddressOK) String() string {
	return fmt.Sprintf("[PUT /orgs/{org_id}/address][%d] updateOrgAddressOK  %+v", 200, o.Payload)
}

func (o *UpdateOrgAddressOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateOrgAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgAddressBadRequest creates a UpdateOrgAddressBadRequest with default headers values
func NewUpdateOrgAddressBadRequest() *UpdateOrgAddressBadRequest {
	return &UpdateOrgAddressBadRequest{}
}

/*
UpdateOrgAddressBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateOrgAddressBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update org address bad request response has a 2xx status code
func (o *UpdateOrgAddressBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update org address bad request response has a 3xx status code
func (o *UpdateOrgAddressBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org address bad request response has a 4xx status code
func (o *UpdateOrgAddressBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update org address bad request response has a 5xx status code
func (o *UpdateOrgAddressBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update org address bad request response a status code equal to that given
func (o *UpdateOrgAddressBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update org address bad request response
func (o *UpdateOrgAddressBadRequest) Code() int {
	return 400
}

func (o *UpdateOrgAddressBadRequest) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org_id}/address][%d] updateOrgAddressBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrgAddressBadRequest) String() string {
	return fmt.Sprintf("[PUT /orgs/{org_id}/address][%d] updateOrgAddressBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrgAddressBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgAddressUnauthorized creates a UpdateOrgAddressUnauthorized with default headers values
func NewUpdateOrgAddressUnauthorized() *UpdateOrgAddressUnauthorized {
	return &UpdateOrgAddressUnauthorized{}
}

/*
UpdateOrgAddressUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateOrgAddressUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update org address unauthorized response has a 2xx status code
func (o *UpdateOrgAddressUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update org address unauthorized response has a 3xx status code
func (o *UpdateOrgAddressUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org address unauthorized response has a 4xx status code
func (o *UpdateOrgAddressUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update org address unauthorized response has a 5xx status code
func (o *UpdateOrgAddressUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update org address unauthorized response a status code equal to that given
func (o *UpdateOrgAddressUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update org address unauthorized response
func (o *UpdateOrgAddressUnauthorized) Code() int {
	return 401
}

func (o *UpdateOrgAddressUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org_id}/address][%d] updateOrgAddressUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrgAddressUnauthorized) String() string {
	return fmt.Sprintf("[PUT /orgs/{org_id}/address][%d] updateOrgAddressUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrgAddressUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgAddressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgAddressForbidden creates a UpdateOrgAddressForbidden with default headers values
func NewUpdateOrgAddressForbidden() *UpdateOrgAddressForbidden {
	return &UpdateOrgAddressForbidden{}
}

/*
UpdateOrgAddressForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateOrgAddressForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update org address forbidden response has a 2xx status code
func (o *UpdateOrgAddressForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update org address forbidden response has a 3xx status code
func (o *UpdateOrgAddressForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org address forbidden response has a 4xx status code
func (o *UpdateOrgAddressForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update org address forbidden response has a 5xx status code
func (o *UpdateOrgAddressForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update org address forbidden response a status code equal to that given
func (o *UpdateOrgAddressForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update org address forbidden response
func (o *UpdateOrgAddressForbidden) Code() int {
	return 403
}

func (o *UpdateOrgAddressForbidden) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org_id}/address][%d] updateOrgAddressForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrgAddressForbidden) String() string {
	return fmt.Sprintf("[PUT /orgs/{org_id}/address][%d] updateOrgAddressForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrgAddressForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgAddressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgAddressInternalServerError creates a UpdateOrgAddressInternalServerError with default headers values
func NewUpdateOrgAddressInternalServerError() *UpdateOrgAddressInternalServerError {
	return &UpdateOrgAddressInternalServerError{}
}

/*
UpdateOrgAddressInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateOrgAddressInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update org address internal server error response has a 2xx status code
func (o *UpdateOrgAddressInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update org address internal server error response has a 3xx status code
func (o *UpdateOrgAddressInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org address internal server error response has a 4xx status code
func (o *UpdateOrgAddressInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update org address internal server error response has a 5xx status code
func (o *UpdateOrgAddressInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update org address internal server error response a status code equal to that given
func (o *UpdateOrgAddressInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update org address internal server error response
func (o *UpdateOrgAddressInternalServerError) Code() int {
	return 500
}

func (o *UpdateOrgAddressInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org_id}/address][%d] updateOrgAddressInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateOrgAddressInternalServerError) String() string {
	return fmt.Sprintf("[PUT /orgs/{org_id}/address][%d] updateOrgAddressInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateOrgAddressInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
