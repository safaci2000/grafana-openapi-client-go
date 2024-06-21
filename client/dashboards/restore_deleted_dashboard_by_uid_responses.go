// Code generated by go-swagger; DO NOT EDIT.

package dashboards

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

// RestoreDeletedDashboardByUIDReader is a Reader for the RestoreDeletedDashboardByUID structure.
type RestoreDeletedDashboardByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreDeletedDashboardByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreDeletedDashboardByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRestoreDeletedDashboardByUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRestoreDeletedDashboardByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRestoreDeletedDashboardByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestoreDeletedDashboardByUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRestoreDeletedDashboardByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /dashboards/uid/{uid}/trash] restoreDeletedDashboardByUID", response, response.Code())
	}
}

// NewRestoreDeletedDashboardByUIDOK creates a RestoreDeletedDashboardByUIDOK with default headers values
func NewRestoreDeletedDashboardByUIDOK() *RestoreDeletedDashboardByUIDOK {
	return &RestoreDeletedDashboardByUIDOK{}
}

/*
RestoreDeletedDashboardByUIDOK describes a response with status code 200, with default header values.

(empty)
*/
type RestoreDeletedDashboardByUIDOK struct {
	Payload *models.RestoreDeletedDashboardByUIDOKBody
}

// IsSuccess returns true when this restore deleted dashboard by Uid Ok response has a 2xx status code
func (o *RestoreDeletedDashboardByUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restore deleted dashboard by Uid Ok response has a 3xx status code
func (o *RestoreDeletedDashboardByUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore deleted dashboard by Uid Ok response has a 4xx status code
func (o *RestoreDeletedDashboardByUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore deleted dashboard by Uid Ok response has a 5xx status code
func (o *RestoreDeletedDashboardByUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restore deleted dashboard by Uid Ok response a status code equal to that given
func (o *RestoreDeletedDashboardByUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the restore deleted dashboard by Uid Ok response
func (o *RestoreDeletedDashboardByUIDOK) Code() int {
	return 200
}

func (o *RestoreDeletedDashboardByUIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidOk %s", 200, payload)
}

func (o *RestoreDeletedDashboardByUIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidOk %s", 200, payload)
}

func (o *RestoreDeletedDashboardByUIDOK) GetPayload() *models.RestoreDeletedDashboardByUIDOKBody {
	return o.Payload
}

func (o *RestoreDeletedDashboardByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestoreDeletedDashboardByUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDeletedDashboardByUIDBadRequest creates a RestoreDeletedDashboardByUIDBadRequest with default headers values
func NewRestoreDeletedDashboardByUIDBadRequest() *RestoreDeletedDashboardByUIDBadRequest {
	return &RestoreDeletedDashboardByUIDBadRequest{}
}

/*
RestoreDeletedDashboardByUIDBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RestoreDeletedDashboardByUIDBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this restore deleted dashboard by Uid bad request response has a 2xx status code
func (o *RestoreDeletedDashboardByUIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restore deleted dashboard by Uid bad request response has a 3xx status code
func (o *RestoreDeletedDashboardByUIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore deleted dashboard by Uid bad request response has a 4xx status code
func (o *RestoreDeletedDashboardByUIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this restore deleted dashboard by Uid bad request response has a 5xx status code
func (o *RestoreDeletedDashboardByUIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this restore deleted dashboard by Uid bad request response a status code equal to that given
func (o *RestoreDeletedDashboardByUIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the restore deleted dashboard by Uid bad request response
func (o *RestoreDeletedDashboardByUIDBadRequest) Code() int {
	return 400
}

func (o *RestoreDeletedDashboardByUIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidBadRequest %s", 400, payload)
}

func (o *RestoreDeletedDashboardByUIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidBadRequest %s", 400, payload)
}

func (o *RestoreDeletedDashboardByUIDBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDeletedDashboardByUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDeletedDashboardByUIDUnauthorized creates a RestoreDeletedDashboardByUIDUnauthorized with default headers values
func NewRestoreDeletedDashboardByUIDUnauthorized() *RestoreDeletedDashboardByUIDUnauthorized {
	return &RestoreDeletedDashboardByUIDUnauthorized{}
}

/*
RestoreDeletedDashboardByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RestoreDeletedDashboardByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this restore deleted dashboard by Uid unauthorized response has a 2xx status code
func (o *RestoreDeletedDashboardByUIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restore deleted dashboard by Uid unauthorized response has a 3xx status code
func (o *RestoreDeletedDashboardByUIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore deleted dashboard by Uid unauthorized response has a 4xx status code
func (o *RestoreDeletedDashboardByUIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this restore deleted dashboard by Uid unauthorized response has a 5xx status code
func (o *RestoreDeletedDashboardByUIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this restore deleted dashboard by Uid unauthorized response a status code equal to that given
func (o *RestoreDeletedDashboardByUIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the restore deleted dashboard by Uid unauthorized response
func (o *RestoreDeletedDashboardByUIDUnauthorized) Code() int {
	return 401
}

func (o *RestoreDeletedDashboardByUIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidUnauthorized %s", 401, payload)
}

func (o *RestoreDeletedDashboardByUIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidUnauthorized %s", 401, payload)
}

func (o *RestoreDeletedDashboardByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDeletedDashboardByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDeletedDashboardByUIDForbidden creates a RestoreDeletedDashboardByUIDForbidden with default headers values
func NewRestoreDeletedDashboardByUIDForbidden() *RestoreDeletedDashboardByUIDForbidden {
	return &RestoreDeletedDashboardByUIDForbidden{}
}

/*
RestoreDeletedDashboardByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RestoreDeletedDashboardByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this restore deleted dashboard by Uid forbidden response has a 2xx status code
func (o *RestoreDeletedDashboardByUIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restore deleted dashboard by Uid forbidden response has a 3xx status code
func (o *RestoreDeletedDashboardByUIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore deleted dashboard by Uid forbidden response has a 4xx status code
func (o *RestoreDeletedDashboardByUIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this restore deleted dashboard by Uid forbidden response has a 5xx status code
func (o *RestoreDeletedDashboardByUIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this restore deleted dashboard by Uid forbidden response a status code equal to that given
func (o *RestoreDeletedDashboardByUIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the restore deleted dashboard by Uid forbidden response
func (o *RestoreDeletedDashboardByUIDForbidden) Code() int {
	return 403
}

func (o *RestoreDeletedDashboardByUIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidForbidden %s", 403, payload)
}

func (o *RestoreDeletedDashboardByUIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidForbidden %s", 403, payload)
}

func (o *RestoreDeletedDashboardByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDeletedDashboardByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDeletedDashboardByUIDNotFound creates a RestoreDeletedDashboardByUIDNotFound with default headers values
func NewRestoreDeletedDashboardByUIDNotFound() *RestoreDeletedDashboardByUIDNotFound {
	return &RestoreDeletedDashboardByUIDNotFound{}
}

/*
RestoreDeletedDashboardByUIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type RestoreDeletedDashboardByUIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this restore deleted dashboard by Uid not found response has a 2xx status code
func (o *RestoreDeletedDashboardByUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restore deleted dashboard by Uid not found response has a 3xx status code
func (o *RestoreDeletedDashboardByUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore deleted dashboard by Uid not found response has a 4xx status code
func (o *RestoreDeletedDashboardByUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this restore deleted dashboard by Uid not found response has a 5xx status code
func (o *RestoreDeletedDashboardByUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this restore deleted dashboard by Uid not found response a status code equal to that given
func (o *RestoreDeletedDashboardByUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the restore deleted dashboard by Uid not found response
func (o *RestoreDeletedDashboardByUIDNotFound) Code() int {
	return 404
}

func (o *RestoreDeletedDashboardByUIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidNotFound %s", 404, payload)
}

func (o *RestoreDeletedDashboardByUIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidNotFound %s", 404, payload)
}

func (o *RestoreDeletedDashboardByUIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDeletedDashboardByUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDeletedDashboardByUIDInternalServerError creates a RestoreDeletedDashboardByUIDInternalServerError with default headers values
func NewRestoreDeletedDashboardByUIDInternalServerError() *RestoreDeletedDashboardByUIDInternalServerError {
	return &RestoreDeletedDashboardByUIDInternalServerError{}
}

/*
RestoreDeletedDashboardByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RestoreDeletedDashboardByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this restore deleted dashboard by Uid internal server error response has a 2xx status code
func (o *RestoreDeletedDashboardByUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restore deleted dashboard by Uid internal server error response has a 3xx status code
func (o *RestoreDeletedDashboardByUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore deleted dashboard by Uid internal server error response has a 4xx status code
func (o *RestoreDeletedDashboardByUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore deleted dashboard by Uid internal server error response has a 5xx status code
func (o *RestoreDeletedDashboardByUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this restore deleted dashboard by Uid internal server error response a status code equal to that given
func (o *RestoreDeletedDashboardByUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the restore deleted dashboard by Uid internal server error response
func (o *RestoreDeletedDashboardByUIDInternalServerError) Code() int {
	return 500
}

func (o *RestoreDeletedDashboardByUIDInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidInternalServerError %s", 500, payload)
}

func (o *RestoreDeletedDashboardByUIDInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dashboards/uid/{uid}/trash][%d] restoreDeletedDashboardByUidInternalServerError %s", 500, payload)
}

func (o *RestoreDeletedDashboardByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDeletedDashboardByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
