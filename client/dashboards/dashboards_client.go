// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new dashboards API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dashboards API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CalculateDashboardDiff(params *CalculateDashboardDiffParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CalculateDashboardDiffOK, error)

	DeleteDashboardByUID(params *DeleteDashboardByUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDashboardByUIDOK, error)

	GetDashboardByUID(params *GetDashboardByUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardByUIDOK, error)

	GetDashboardTags(params *GetDashboardTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardTagsOK, error)

	GetHomeDashboard(params *GetHomeDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHomeDashboardOK, error)

	ImportDashboard(params *ImportDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImportDashboardOK, error)

	PostDashboard(params *PostDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDashboardOK, error)

	TrimDashboard(params *TrimDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TrimDashboardOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CalculateDashboardDiff performs diff on two dashboards
*/
func (a *Client) CalculateDashboardDiff(params *CalculateDashboardDiffParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CalculateDashboardDiffOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCalculateDashboardDiffParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "calculateDashboardDiff",
		Method:             "POST",
		PathPattern:        "/dashboards/calculate-diff",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CalculateDashboardDiffReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CalculateDashboardDiffOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for calculateDashboardDiff: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDashboardByUID deletes dashboard by uid

Will delete the dashboard given the specified unique identifier (uid).
*/
func (a *Client) DeleteDashboardByUID(params *DeleteDashboardByUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDashboardByUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDashboardByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDashboardByUID",
		Method:             "DELETE",
		PathPattern:        "/dashboards/uid/{uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDashboardByUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDashboardByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDashboardByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDashboardByUID gets dashboard by uid

Will return the dashboard given the dashboard unique identifier (uid).
*/
func (a *Client) GetDashboardByUID(params *GetDashboardByUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardByUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDashboardByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardByUID",
		Method:             "GET",
		PathPattern:        "/dashboards/uid/{uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardByUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDashboardTags gets all dashboards tags of an organisation
*/
func (a *Client) GetDashboardTags(params *GetDashboardTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDashboardTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardTags",
		Method:             "GET",
		PathPattern:        "/dashboards/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHomeDashboard gets home dashboard
*/
func (a *Client) GetHomeDashboard(params *GetHomeDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHomeDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHomeDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHomeDashboard",
		Method:             "GET",
		PathPattern:        "/dashboards/home",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetHomeDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHomeDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHomeDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImportDashboard imports dashboard
*/
func (a *Client) ImportDashboard(params *ImportDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImportDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "importDashboard",
		Method:             "POST",
		PathPattern:        "/dashboards/import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImportDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImportDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostDashboard creates update dashboard

Creates a new dashboard or updates an existing dashboard.
*/
func (a *Client) PostDashboard(params *PostDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postDashboard",
		Method:             "POST",
		PathPattern:        "/dashboards/db",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TrimDashboard trims defaults from dashboard
*/
func (a *Client) TrimDashboard(params *TrimDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TrimDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTrimDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "trimDashboard",
		Method:             "POST",
		PathPattern:        "/dashboards/trim",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TrimDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TrimDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for trimDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
