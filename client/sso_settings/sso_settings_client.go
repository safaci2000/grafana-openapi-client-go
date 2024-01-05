// Code generated by go-swagger; DO NOT EDIT.

package sso_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new sso settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sso settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	RemoveProviderSettings(key string, opts ...ClientOption) (*RemoveProviderSettingsNoContent, error)
	RemoveProviderSettingsWithParams(params *RemoveProviderSettingsParams, opts ...ClientOption) (*RemoveProviderSettingsNoContent, error)

	UpdateProviderSettings(key string, body *models.SSOSettings, opts ...ClientOption) (*UpdateProviderSettingsNoContent, error)
	UpdateProviderSettingsWithParams(params *UpdateProviderSettingsParams, opts ...ClientOption) (*UpdateProviderSettingsNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
RemoveProviderSettings removes s s o settings

Removes the SSO Settings for a provider.

You need to have a permission with action `settings:write` and scope `settings:auth.<provider>:*`.
*/
func (a *Client) RemoveProviderSettings(key string, opts ...ClientOption) (*RemoveProviderSettingsNoContent, error) {
	params := NewRemoveProviderSettingsParams().WithKey(key)
	return a.RemoveProviderSettingsWithParams(params, opts...)
}

func (a *Client) RemoveProviderSettingsWithParams(params *RemoveProviderSettingsParams, opts ...ClientOption) (*RemoveProviderSettingsNoContent, error) {
	if params == nil {
		params = NewRemoveProviderSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeProviderSettings",
		Method:             "DELETE",
		PathPattern:        "/v1/sso-settings/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveProviderSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveProviderSettingsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeProviderSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateProviderSettings updates s s o settings

Inserts or updates the SSO Settings for a provider.

You need to have a permission with action `settings:write` and scope `settings:auth.<provider>:*`.
*/
func (a *Client) UpdateProviderSettings(key string, body *models.SSOSettings, opts ...ClientOption) (*UpdateProviderSettingsNoContent, error) {
	params := NewUpdateProviderSettingsParams().WithBody(body).WithKey(key)
	return a.UpdateProviderSettingsWithParams(params, opts...)
}

func (a *Client) UpdateProviderSettingsWithParams(params *UpdateProviderSettingsParams, opts ...ClientOption) (*UpdateProviderSettingsNoContent, error) {
	if params == nil {
		params = NewUpdateProviderSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateProviderSettings",
		Method:             "PUT",
		PathPattern:        "/v1/sso-settings/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateProviderSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProviderSettingsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateProviderSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// WithAuthInfo changes the transport on the client
func WithAuthInfo(authInfo runtime.ClientAuthInfoWriter) ClientOption {
	return func(op *runtime.ClientOperation) {
		op.AuthInfo = authInfo
	}
}
