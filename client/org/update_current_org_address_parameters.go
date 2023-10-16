// Code generated by go-swagger; DO NOT EDIT.

package org

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

// NewUpdateCurrentOrgAddressParams creates a new UpdateCurrentOrgAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCurrentOrgAddressParams() *UpdateCurrentOrgAddressParams {
	return &UpdateCurrentOrgAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCurrentOrgAddressParamsWithTimeout creates a new UpdateCurrentOrgAddressParams object
// with the ability to set a timeout on a request.
func NewUpdateCurrentOrgAddressParamsWithTimeout(timeout time.Duration) *UpdateCurrentOrgAddressParams {
	return &UpdateCurrentOrgAddressParams{
		timeout: timeout,
	}
}

// NewUpdateCurrentOrgAddressParamsWithContext creates a new UpdateCurrentOrgAddressParams object
// with the ability to set a context for a request.
func NewUpdateCurrentOrgAddressParamsWithContext(ctx context.Context) *UpdateCurrentOrgAddressParams {
	return &UpdateCurrentOrgAddressParams{
		Context: ctx,
	}
}

// NewUpdateCurrentOrgAddressParamsWithHTTPClient creates a new UpdateCurrentOrgAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCurrentOrgAddressParamsWithHTTPClient(client *http.Client) *UpdateCurrentOrgAddressParams {
	return &UpdateCurrentOrgAddressParams{
		HTTPClient: client,
	}
}

/*
UpdateCurrentOrgAddressParams contains all the parameters to send to the API endpoint

	for the update current org address operation.

	Typically these are written to a http.Request.
*/
type UpdateCurrentOrgAddressParams struct {

	// Body.
	Body *models.UpdateOrgAddressForm

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update current org address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCurrentOrgAddressParams) WithDefaults() *UpdateCurrentOrgAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update current org address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCurrentOrgAddressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update current org address params
func (o *UpdateCurrentOrgAddressParams) WithTimeout(timeout time.Duration) *UpdateCurrentOrgAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update current org address params
func (o *UpdateCurrentOrgAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update current org address params
func (o *UpdateCurrentOrgAddressParams) WithContext(ctx context.Context) *UpdateCurrentOrgAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update current org address params
func (o *UpdateCurrentOrgAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update current org address params
func (o *UpdateCurrentOrgAddressParams) WithHTTPClient(client *http.Client) *UpdateCurrentOrgAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update current org address params
func (o *UpdateCurrentOrgAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update current org address params
func (o *UpdateCurrentOrgAddressParams) WithBody(body *models.UpdateOrgAddressForm) *UpdateCurrentOrgAddressParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update current org address params
func (o *UpdateCurrentOrgAddressParams) SetBody(body *models.UpdateOrgAddressForm) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCurrentOrgAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
