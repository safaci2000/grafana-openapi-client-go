// Code generated by go-swagger; DO NOT EDIT.

package dashboard_permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command
// To edit this file, edit the custom template in templates/clientClient.gotmpl
// More info on custom templates can be found in the README or here: https://github.com/go-swagger/go-swagger/blob/master/docs/generate/templates.md
// The template itself can be found here: https://github.com/go-swagger/go-swagger/blob/master/generator/templates/client/client.gotmpl

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new dashboard permissions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dashboard permissions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDashboardPermissionsListByID(params *GetDashboardPermissionsListByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardPermissionsListByIDOK, error)

	GetDashboardPermissionsListByUID(params *GetDashboardPermissionsListByUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardPermissionsListByUIDOK, error)

	UpdateDashboardPermissionsByID(params *UpdateDashboardPermissionsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDashboardPermissionsByIDOK, error)

	UpdateDashboardPermissionsByUID(params *UpdateDashboardPermissionsByUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDashboardPermissionsByUIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetDashboardPermissionsListByID gets all existing permissions for the given dashboard

Please refer to [updated API](#/dashboard_permissions/getDashboardPermissionsListByUID) instead
*/
func (a *Client) GetDashboardPermissionsListByID(params *GetDashboardPermissionsListByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardPermissionsListByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDashboardPermissionsListByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardPermissionsListByID",
		Method:             "GET",
		PathPattern:        "/dashboards/id/{DashboardID}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardPermissionsListByIDReader{formats: a.formats},
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
	success, ok := result.(*GetDashboardPermissionsListByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardPermissionsListByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDashboardPermissionsListByUID gets all existing permissions for the given dashboard
*/
func (a *Client) GetDashboardPermissionsListByUID(params *GetDashboardPermissionsListByUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardPermissionsListByUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDashboardPermissionsListByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardPermissionsListByUID",
		Method:             "GET",
		PathPattern:        "/dashboards/uid/{uid}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardPermissionsListByUIDReader{formats: a.formats},
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
	success, ok := result.(*GetDashboardPermissionsListByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardPermissionsListByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	UpdateDashboardPermissionsByID updates permissions for a dashboard

	Please refer to [updated API](#/dashboard_permissions/updateDashboardPermissionsByUID) instead

This operation will remove existing permissions if they’re not included in the request.
*/
func (a *Client) UpdateDashboardPermissionsByID(params *UpdateDashboardPermissionsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDashboardPermissionsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDashboardPermissionsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDashboardPermissionsByID",
		Method:             "POST",
		PathPattern:        "/dashboards/id/{DashboardID}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateDashboardPermissionsByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateDashboardPermissionsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDashboardPermissionsByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDashboardPermissionsByUID updates permissions for a dashboard

This operation will remove existing permissions if they’re not included in the request.
*/
func (a *Client) UpdateDashboardPermissionsByUID(params *UpdateDashboardPermissionsByUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDashboardPermissionsByUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDashboardPermissionsByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDashboardPermissionsByUID",
		Method:             "POST",
		PathPattern:        "/dashboards/uid/{uid}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateDashboardPermissionsByUIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateDashboardPermissionsByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDashboardPermissionsByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
