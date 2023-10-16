// Code generated by go-swagger; DO NOT EDIT.

package orgs

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

// New creates a new orgs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for orgs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddOrgUser(params *AddOrgUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddOrgUserOK, error)

	CreateOrg(params *CreateOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOrgOK, error)

	DeleteOrgByID(params *DeleteOrgByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrgByIDOK, error)

	GetOrgByID(params *GetOrgByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgByIDOK, error)

	GetOrgByName(params *GetOrgByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgByNameOK, error)

	GetOrgQuota(params *GetOrgQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgQuotaOK, error)

	GetOrgUsers(params *GetOrgUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgUsersOK, error)

	RemoveOrgUser(params *RemoveOrgUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveOrgUserOK, error)

	SearchOrgUsers(params *SearchOrgUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOrgUsersOK, error)

	SearchOrgs(params *SearchOrgsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOrgsOK, error)

	UpdateOrg(params *UpdateOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrgOK, error)

	UpdateOrgAddress(params *UpdateOrgAddressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrgAddressOK, error)

	UpdateOrgQuota(params *UpdateOrgQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrgQuotaOK, error)

	UpdateOrgUser(params *UpdateOrgUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrgUserOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	AddOrgUser adds a new user to the current organization

	Adds a global user to the current organization.

If you are running Grafana Enterprise and have Fine-grained access control enabled
you need to have a permission with action: `org.users:add` with scope `users:*`.
*/
func (a *Client) AddOrgUser(params *AddOrgUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddOrgUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOrgUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addOrgUser",
		Method:             "POST",
		PathPattern:        "/orgs/{org_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddOrgUserReader{formats: a.formats},
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
	success, ok := result.(*AddOrgUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addOrgUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateOrg creates organization

Only works if [users.allow_org_create](https://grafana.com/docs/grafana/latest/administration/configuration/#allow_org_create) is set.
*/
func (a *Client) CreateOrg(params *CreateOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOrgOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrgParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createOrg",
		Method:             "POST",
		PathPattern:        "/orgs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateOrgReader{formats: a.formats},
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
	success, ok := result.(*CreateOrgOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createOrg: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteOrgByID deletes organization
*/
func (a *Client) DeleteOrgByID(params *DeleteOrgByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrgByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrgByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteOrgByID",
		Method:             "DELETE",
		PathPattern:        "/orgs/{org_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOrgByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteOrgByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteOrgByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrgByID gets organization by ID
*/
func (a *Client) GetOrgByID(params *GetOrgByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrgByID",
		Method:             "GET",
		PathPattern:        "/orgs/{org_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrgByIDReader{formats: a.formats},
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
	success, ok := result.(*GetOrgByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrgByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrgByName gets organization by ID
*/
func (a *Client) GetOrgByName(params *GetOrgByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrgByName",
		Method:             "GET",
		PathPattern:        "/orgs/name/{org_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrgByNameReader{formats: a.formats},
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
	success, ok := result.(*GetOrgByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrgByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrgQuota fetches organization quota

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `orgs.quotas:read` and scope `org:id:1` (orgIDScope).
*/
func (a *Client) GetOrgQuota(params *GetOrgQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrgQuota",
		Method:             "GET",
		PathPattern:        "/orgs/{org_id}/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrgQuotaReader{formats: a.formats},
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
	success, ok := result.(*GetOrgQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrgQuota: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetOrgUsers gets users in organization

	If you are running Grafana Enterprise and have Fine-grained access control enabled

you need to have a permission with action: `org.users:read` with scope `users:*`.
*/
func (a *Client) GetOrgUsers(params *GetOrgUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrgUsers",
		Method:             "GET",
		PathPattern:        "/orgs/{org_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrgUsersReader{formats: a.formats},
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
	success, ok := result.(*GetOrgUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrgUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	RemoveOrgUser deletes user in current organization

	If you are running Grafana Enterprise and have Fine-grained access control enabled

you need to have a permission with action: `org.users:remove` with scope `users:*`.
*/
func (a *Client) RemoveOrgUser(params *RemoveOrgUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveOrgUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveOrgUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeOrgUser",
		Method:             "DELETE",
		PathPattern:        "/orgs/{org_id}/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveOrgUserReader{formats: a.formats},
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
	success, ok := result.(*RemoveOrgUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeOrgUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SearchOrgUsers searches users in organization

	If you are running Grafana Enterprise and have Fine-grained access control enabled

you need to have a permission with action: `org.users:read` with scope `users:*`.
*/
func (a *Client) SearchOrgUsers(params *SearchOrgUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOrgUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchOrgUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchOrgUsers",
		Method:             "GET",
		PathPattern:        "/orgs/{org_id}/users/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchOrgUsersReader{formats: a.formats},
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
	success, ok := result.(*SearchOrgUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchOrgUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchOrgs searches all organizations
*/
func (a *Client) SearchOrgs(params *SearchOrgsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOrgsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchOrgsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchOrgs",
		Method:             "GET",
		PathPattern:        "/orgs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchOrgsReader{formats: a.formats},
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
	success, ok := result.(*SearchOrgsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchOrgs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrg updates organization
*/
func (a *Client) UpdateOrg(params *UpdateOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrgOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrgParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrg",
		Method:             "PUT",
		PathPattern:        "/orgs/{org_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrgReader{formats: a.formats},
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
	success, ok := result.(*UpdateOrgOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrg: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrgAddress updates organization s address
*/
func (a *Client) UpdateOrgAddress(params *UpdateOrgAddressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrgAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrgAddressParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrgAddress",
		Method:             "PUT",
		PathPattern:        "/orgs/{org_id}/address",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrgAddressReader{formats: a.formats},
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
	success, ok := result.(*UpdateOrgAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrgAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrgQuota updates user quota

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `orgs.quotas:write` and scope `org:id:1` (orgIDScope).
*/
func (a *Client) UpdateOrgQuota(params *UpdateOrgQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrgQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrgQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrgQuota",
		Method:             "PUT",
		PathPattern:        "/orgs/{org_id}/quotas/{quota_target}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrgQuotaReader{formats: a.formats},
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
	success, ok := result.(*UpdateOrgQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrgQuota: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	UpdateOrgUser updates users in organization

	If you are running Grafana Enterprise and have Fine-grained access control enabled

you need to have a permission with action: `org.users.role:update` with scope `users:*`.
*/
func (a *Client) UpdateOrgUser(params *UpdateOrgUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrgUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrgUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrgUser",
		Method:             "PATCH",
		PathPattern:        "/orgs/{org_id}/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrgUserReader{formats: a.formats},
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
	success, ok := result.(*UpdateOrgUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrgUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
