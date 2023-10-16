// Code generated by go-swagger; DO NOT EDIT.

package datasources

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

// NewAddDataSourceParams creates a new AddDataSourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDataSourceParams() *AddDataSourceParams {
	return &AddDataSourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDataSourceParamsWithTimeout creates a new AddDataSourceParams object
// with the ability to set a timeout on a request.
func NewAddDataSourceParamsWithTimeout(timeout time.Duration) *AddDataSourceParams {
	return &AddDataSourceParams{
		timeout: timeout,
	}
}

// NewAddDataSourceParamsWithContext creates a new AddDataSourceParams object
// with the ability to set a context for a request.
func NewAddDataSourceParamsWithContext(ctx context.Context) *AddDataSourceParams {
	return &AddDataSourceParams{
		Context: ctx,
	}
}

// NewAddDataSourceParamsWithHTTPClient creates a new AddDataSourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDataSourceParamsWithHTTPClient(client *http.Client) *AddDataSourceParams {
	return &AddDataSourceParams{
		HTTPClient: client,
	}
}

/*
AddDataSourceParams contains all the parameters to send to the API endpoint

	for the add data source operation.

	Typically these are written to a http.Request.
*/
type AddDataSourceParams struct {

	// Body.
	Body *models.AddDataSourceCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add data source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDataSourceParams) WithDefaults() *AddDataSourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add data source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDataSourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add data source params
func (o *AddDataSourceParams) WithTimeout(timeout time.Duration) *AddDataSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add data source params
func (o *AddDataSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add data source params
func (o *AddDataSourceParams) WithContext(ctx context.Context) *AddDataSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add data source params
func (o *AddDataSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add data source params
func (o *AddDataSourceParams) WithHTTPClient(client *http.Client) *AddDataSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add data source params
func (o *AddDataSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add data source params
func (o *AddDataSourceParams) WithBody(body *models.AddDataSourceCommand) *AddDataSourceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add data source params
func (o *AddDataSourceParams) SetBody(body *models.AddDataSourceCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddDataSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
