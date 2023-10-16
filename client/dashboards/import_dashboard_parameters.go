// Code generated by go-swagger; DO NOT EDIT.

package dashboards

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

// NewImportDashboardParams creates a new ImportDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportDashboardParams() *ImportDashboardParams {
	return &ImportDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportDashboardParamsWithTimeout creates a new ImportDashboardParams object
// with the ability to set a timeout on a request.
func NewImportDashboardParamsWithTimeout(timeout time.Duration) *ImportDashboardParams {
	return &ImportDashboardParams{
		timeout: timeout,
	}
}

// NewImportDashboardParamsWithContext creates a new ImportDashboardParams object
// with the ability to set a context for a request.
func NewImportDashboardParamsWithContext(ctx context.Context) *ImportDashboardParams {
	return &ImportDashboardParams{
		Context: ctx,
	}
}

// NewImportDashboardParamsWithHTTPClient creates a new ImportDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportDashboardParamsWithHTTPClient(client *http.Client) *ImportDashboardParams {
	return &ImportDashboardParams{
		HTTPClient: client,
	}
}

/*
ImportDashboardParams contains all the parameters to send to the API endpoint

	for the import dashboard operation.

	Typically these are written to a http.Request.
*/
type ImportDashboardParams struct {

	// Body.
	Body *models.ImportDashboardRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportDashboardParams) WithDefaults() *ImportDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the import dashboard params
func (o *ImportDashboardParams) WithTimeout(timeout time.Duration) *ImportDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import dashboard params
func (o *ImportDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import dashboard params
func (o *ImportDashboardParams) WithContext(ctx context.Context) *ImportDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import dashboard params
func (o *ImportDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import dashboard params
func (o *ImportDashboardParams) WithHTTPClient(client *http.Client) *ImportDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import dashboard params
func (o *ImportDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the import dashboard params
func (o *ImportDashboardParams) WithBody(body *models.ImportDashboardRequest) *ImportDashboardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the import dashboard params
func (o *ImportDashboardParams) SetBody(body *models.ImportDashboardRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ImportDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
