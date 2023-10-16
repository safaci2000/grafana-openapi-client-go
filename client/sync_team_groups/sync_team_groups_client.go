// Code generated by go-swagger; DO NOT EDIT.

package sync_team_groups

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

// New creates a new sync team groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sync team groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddTeamGroupAPI(params *AddTeamGroupAPIParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddTeamGroupAPIOK, error)

	GetTeamGroupsAPI(params *GetTeamGroupsAPIParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTeamGroupsAPIOK, error)

	RemoveTeamGroupAPI(params *RemoveTeamGroupAPIParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveTeamGroupAPIOK, error)

	RemoveTeamGroupAPIQuery(params *RemoveTeamGroupAPIQueryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveTeamGroupAPIQueryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddTeamGroupAPI adds external group
*/
func (a *Client) AddTeamGroupAPI(params *AddTeamGroupAPIParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddTeamGroupAPIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddTeamGroupAPIParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addTeamGroupApi",
		Method:             "POST",
		PathPattern:        "/teams/{teamId}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddTeamGroupAPIReader{formats: a.formats},
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
	success, ok := result.(*AddTeamGroupAPIOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addTeamGroupApi: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTeamGroupsAPI gets external groups
*/
func (a *Client) GetTeamGroupsAPI(params *GetTeamGroupsAPIParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTeamGroupsAPIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTeamGroupsAPIParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTeamGroupsApi",
		Method:             "GET",
		PathPattern:        "/teams/{teamId}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTeamGroupsAPIReader{formats: a.formats},
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
	success, ok := result.(*GetTeamGroupsAPIOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTeamGroupsApi: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveTeamGroupAPI removes external group
*/
func (a *Client) RemoveTeamGroupAPI(params *RemoveTeamGroupAPIParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveTeamGroupAPIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveTeamGroupAPIParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeTeamGroupApi",
		Method:             "DELETE",
		PathPattern:        "/teams/{teamId}/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveTeamGroupAPIReader{formats: a.formats},
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
	success, ok := result.(*RemoveTeamGroupAPIOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeTeamGroupApi: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveTeamGroupAPIQuery removes external group
*/
func (a *Client) RemoveTeamGroupAPIQuery(params *RemoveTeamGroupAPIQueryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveTeamGroupAPIQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveTeamGroupAPIQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeTeamGroupApiQuery",
		Method:             "DELETE",
		PathPattern:        "/teams/{teamId}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveTeamGroupAPIQueryReader{formats: a.formats},
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
	success, ok := result.(*RemoveTeamGroupAPIQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeTeamGroupApiQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
