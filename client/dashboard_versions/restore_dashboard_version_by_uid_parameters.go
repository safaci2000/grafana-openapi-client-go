// Code generated by go-swagger; DO NOT EDIT.

package dashboard_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewRestoreDashboardVersionByUIDParams creates a new RestoreDashboardVersionByUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestoreDashboardVersionByUIDParams() *RestoreDashboardVersionByUIDParams {
	return &RestoreDashboardVersionByUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreDashboardVersionByUIDParamsWithTimeout creates a new RestoreDashboardVersionByUIDParams object
// with the ability to set a timeout on a request.
func NewRestoreDashboardVersionByUIDParamsWithTimeout(timeout time.Duration) *RestoreDashboardVersionByUIDParams {
	return &RestoreDashboardVersionByUIDParams{
		timeout: timeout,
	}
}

// NewRestoreDashboardVersionByUIDParamsWithContext creates a new RestoreDashboardVersionByUIDParams object
// with the ability to set a context for a request.
func NewRestoreDashboardVersionByUIDParamsWithContext(ctx context.Context) *RestoreDashboardVersionByUIDParams {
	return &RestoreDashboardVersionByUIDParams{
		Context: ctx,
	}
}

// NewRestoreDashboardVersionByUIDParamsWithHTTPClient creates a new RestoreDashboardVersionByUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestoreDashboardVersionByUIDParamsWithHTTPClient(client *http.Client) *RestoreDashboardVersionByUIDParams {
	return &RestoreDashboardVersionByUIDParams{
		HTTPClient: client,
	}
}

/*
RestoreDashboardVersionByUIDParams contains all the parameters to send to the API endpoint

	for the restore dashboard version by UID operation.

	Typically these are written to a http.Request.
*/
type RestoreDashboardVersionByUIDParams struct {

	// Body.
	Body *models.RestoreDashboardVersionCommand

	// UID.
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restore dashboard version by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreDashboardVersionByUIDParams) WithDefaults() *RestoreDashboardVersionByUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restore dashboard version by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreDashboardVersionByUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restore dashboard version by UID params
func (o *RestoreDashboardVersionByUIDParams) WithTimeout(timeout time.Duration) *RestoreDashboardVersionByUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore dashboard version by UID params
func (o *RestoreDashboardVersionByUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore dashboard version by UID params
func (o *RestoreDashboardVersionByUIDParams) WithContext(ctx context.Context) *RestoreDashboardVersionByUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore dashboard version by UID params
func (o *RestoreDashboardVersionByUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore dashboard version by UID params
func (o *RestoreDashboardVersionByUIDParams) WithHTTPClient(client *http.Client) *RestoreDashboardVersionByUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore dashboard version by UID params
func (o *RestoreDashboardVersionByUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the restore dashboard version by UID params
func (o *RestoreDashboardVersionByUIDParams) WithBody(body *models.RestoreDashboardVersionCommand) *RestoreDashboardVersionByUIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the restore dashboard version by UID params
func (o *RestoreDashboardVersionByUIDParams) SetBody(body *models.RestoreDashboardVersionCommand) {
	o.Body = body
}

// WithUID adds the uid to the restore dashboard version by UID params
func (o *RestoreDashboardVersionByUIDParams) WithUID(uid string) *RestoreDashboardVersionByUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the restore dashboard version by UID params
func (o *RestoreDashboardVersionByUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreDashboardVersionByUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
