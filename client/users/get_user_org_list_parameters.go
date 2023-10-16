// Code generated by go-swagger; DO NOT EDIT.

package users

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
	"github.com/go-openapi/swag"
)

// NewGetUserOrgListParams creates a new GetUserOrgListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserOrgListParams() *GetUserOrgListParams {
	return &GetUserOrgListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserOrgListParamsWithTimeout creates a new GetUserOrgListParams object
// with the ability to set a timeout on a request.
func NewGetUserOrgListParamsWithTimeout(timeout time.Duration) *GetUserOrgListParams {
	return &GetUserOrgListParams{
		timeout: timeout,
	}
}

// NewGetUserOrgListParamsWithContext creates a new GetUserOrgListParams object
// with the ability to set a context for a request.
func NewGetUserOrgListParamsWithContext(ctx context.Context) *GetUserOrgListParams {
	return &GetUserOrgListParams{
		Context: ctx,
	}
}

// NewGetUserOrgListParamsWithHTTPClient creates a new GetUserOrgListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserOrgListParamsWithHTTPClient(client *http.Client) *GetUserOrgListParams {
	return &GetUserOrgListParams{
		HTTPClient: client,
	}
}

/*
GetUserOrgListParams contains all the parameters to send to the API endpoint

	for the get user org list operation.

	Typically these are written to a http.Request.
*/
type GetUserOrgListParams struct {

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user org list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserOrgListParams) WithDefaults() *GetUserOrgListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user org list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserOrgListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user org list params
func (o *GetUserOrgListParams) WithTimeout(timeout time.Duration) *GetUserOrgListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user org list params
func (o *GetUserOrgListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user org list params
func (o *GetUserOrgListParams) WithContext(ctx context.Context) *GetUserOrgListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user org list params
func (o *GetUserOrgListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user org list params
func (o *GetUserOrgListParams) WithHTTPClient(client *http.Client) *GetUserOrgListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user org list params
func (o *GetUserOrgListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get user org list params
func (o *GetUserOrgListParams) WithUserID(userID int64) *GetUserOrgListParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user org list params
func (o *GetUserOrgListParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserOrgListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
