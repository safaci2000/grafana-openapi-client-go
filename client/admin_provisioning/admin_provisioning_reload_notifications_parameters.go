// Code generated by go-swagger; DO NOT EDIT.

package admin_provisioning

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
)

// NewAdminProvisioningReloadNotificationsParams creates a new AdminProvisioningReloadNotificationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminProvisioningReloadNotificationsParams() *AdminProvisioningReloadNotificationsParams {
	return &AdminProvisioningReloadNotificationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminProvisioningReloadNotificationsParamsWithTimeout creates a new AdminProvisioningReloadNotificationsParams object
// with the ability to set a timeout on a request.
func NewAdminProvisioningReloadNotificationsParamsWithTimeout(timeout time.Duration) *AdminProvisioningReloadNotificationsParams {
	return &AdminProvisioningReloadNotificationsParams{
		timeout: timeout,
	}
}

// NewAdminProvisioningReloadNotificationsParamsWithContext creates a new AdminProvisioningReloadNotificationsParams object
// with the ability to set a context for a request.
func NewAdminProvisioningReloadNotificationsParamsWithContext(ctx context.Context) *AdminProvisioningReloadNotificationsParams {
	return &AdminProvisioningReloadNotificationsParams{
		Context: ctx,
	}
}

// NewAdminProvisioningReloadNotificationsParamsWithHTTPClient creates a new AdminProvisioningReloadNotificationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminProvisioningReloadNotificationsParamsWithHTTPClient(client *http.Client) *AdminProvisioningReloadNotificationsParams {
	return &AdminProvisioningReloadNotificationsParams{
		HTTPClient: client,
	}
}

/*
AdminProvisioningReloadNotificationsParams contains all the parameters to send to the API endpoint

	for the admin provisioning reload notifications operation.

	Typically these are written to a http.Request.
*/
type AdminProvisioningReloadNotificationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin provisioning reload notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminProvisioningReloadNotificationsParams) WithDefaults() *AdminProvisioningReloadNotificationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin provisioning reload notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminProvisioningReloadNotificationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin provisioning reload notifications params
func (o *AdminProvisioningReloadNotificationsParams) WithTimeout(timeout time.Duration) *AdminProvisioningReloadNotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin provisioning reload notifications params
func (o *AdminProvisioningReloadNotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin provisioning reload notifications params
func (o *AdminProvisioningReloadNotificationsParams) WithContext(ctx context.Context) *AdminProvisioningReloadNotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin provisioning reload notifications params
func (o *AdminProvisioningReloadNotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin provisioning reload notifications params
func (o *AdminProvisioningReloadNotificationsParams) WithHTTPClient(client *http.Client) *AdminProvisioningReloadNotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin provisioning reload notifications params
func (o *AdminProvisioningReloadNotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AdminProvisioningReloadNotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
