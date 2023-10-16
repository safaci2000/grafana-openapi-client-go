// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewGetTeamMembersParams creates a new GetTeamMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTeamMembersParams() *GetTeamMembersParams {
	return &GetTeamMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamMembersParamsWithTimeout creates a new GetTeamMembersParams object
// with the ability to set a timeout on a request.
func NewGetTeamMembersParamsWithTimeout(timeout time.Duration) *GetTeamMembersParams {
	return &GetTeamMembersParams{
		timeout: timeout,
	}
}

// NewGetTeamMembersParamsWithContext creates a new GetTeamMembersParams object
// with the ability to set a context for a request.
func NewGetTeamMembersParamsWithContext(ctx context.Context) *GetTeamMembersParams {
	return &GetTeamMembersParams{
		Context: ctx,
	}
}

// NewGetTeamMembersParamsWithHTTPClient creates a new GetTeamMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTeamMembersParamsWithHTTPClient(client *http.Client) *GetTeamMembersParams {
	return &GetTeamMembersParams{
		HTTPClient: client,
	}
}

/*
GetTeamMembersParams contains all the parameters to send to the API endpoint

	for the get team members operation.

	Typically these are written to a http.Request.
*/
type GetTeamMembersParams struct {

	// TeamID.
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get team members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamMembersParams) WithDefaults() *GetTeamMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get team members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get team members params
func (o *GetTeamMembersParams) WithTimeout(timeout time.Duration) *GetTeamMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get team members params
func (o *GetTeamMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get team members params
func (o *GetTeamMembersParams) WithContext(ctx context.Context) *GetTeamMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get team members params
func (o *GetTeamMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get team members params
func (o *GetTeamMembersParams) WithHTTPClient(client *http.Client) *GetTeamMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get team members params
func (o *GetTeamMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the get team members params
func (o *GetTeamMembersParams) WithTeamID(teamID string) *GetTeamMembersParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the get team members params
func (o *GetTeamMembersParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
