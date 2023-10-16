// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// New creates a new reports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateReport(params *CreateReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateReportOK, error)

	DeleteReport(params *DeleteReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteReportOK, error)

	GetReport(params *GetReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReportOK, error)

	GetReportSettings(params *GetReportSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReportSettingsOK, error)

	GetReports(params *GetReportsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReportsOK, error)

	RenderReportPDF(params *RenderReportPDFParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RenderReportPDFOK, error)

	RenderReportPDFs(params *RenderReportPDFsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RenderReportPDFsOK, error)

	SaveReportSettings(params *SaveReportSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SaveReportSettingsOK, error)

	SendReport(params *SendReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SendReportOK, error)

	SendTestEmail(params *SendTestEmailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SendTestEmailOK, error)

	UpdateReport(params *UpdateReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateReportOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CreateReport creates a report

	Available to org admins only and with a valid license.

You need to have a permission with action `reports.admin:create`.
*/
func (a *Client) CreateReport(params *CreateReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createReport",
		Method:             "POST",
		PathPattern:        "/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateReportReader{formats: a.formats},
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
	success, ok := result.(*CreateReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	DeleteReport deletes a report

	Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports.delete` with scope `reports:id:<report ID>`.
*/
func (a *Client) DeleteReport(params *DeleteReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteReport",
		Method:             "DELETE",
		PathPattern:        "/reports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteReportReader{formats: a.formats},
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
	success, ok := result.(*DeleteReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetReport gets a report

	Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports:read` with scope `reports:id:<report ID>`.
*/
func (a *Client) GetReport(params *GetReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getReport",
		Method:             "GET",
		PathPattern:        "/reports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReportReader{formats: a.formats},
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
	success, ok := result.(*GetReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetReportSettings gets settings

	Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports.settings:read`x.
*/
func (a *Client) GetReportSettings(params *GetReportSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReportSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReportSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getReportSettings",
		Method:             "GET",
		PathPattern:        "/reports/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReportSettingsReader{formats: a.formats},
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
	success, ok := result.(*GetReportSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReportSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetReports lists reports

	Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports:read` with scope `reports:*`.
*/
func (a *Client) GetReports(params *GetReportsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReportsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getReports",
		Method:             "GET",
		PathPattern:        "/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReportsReader{formats: a.formats},
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
	success, ok := result.(*GetReportsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReports: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RenderReportPDF renders report for dashboard

Please refer to [reports enterprise](#/reports/renderReportPDFs) instead. This will be removed in Grafana 10.
*/
func (a *Client) RenderReportPDF(params *RenderReportPDFParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RenderReportPDFOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRenderReportPDFParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "renderReportPDF",
		Method:             "GET",
		PathPattern:        "/reports/render/pdf/{dashboardID}",
		ProducesMediaTypes: []string{"application/pdf"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RenderReportPDFReader{formats: a.formats},
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
	success, ok := result.(*RenderReportPDFOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for renderReportPDF: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RenderReportPDFs renders report for multiple dashboards

Available to all users and with a valid license.
*/
func (a *Client) RenderReportPDFs(params *RenderReportPDFsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RenderReportPDFsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRenderReportPDFsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "renderReportPDFs",
		Method:             "GET",
		PathPattern:        "/reports/render/pdfs",
		ProducesMediaTypes: []string{"application/pdf"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RenderReportPDFsReader{formats: a.formats},
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
	success, ok := result.(*RenderReportPDFsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for renderReportPDFs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SaveReportSettings saves settings

	Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports.settings:write`xx.
*/
func (a *Client) SaveReportSettings(params *SaveReportSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SaveReportSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveReportSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "saveReportSettings",
		Method:             "POST",
		PathPattern:        "/reports/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SaveReportSettingsReader{formats: a.formats},
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
	success, ok := result.(*SaveReportSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for saveReportSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SendReport sends a report

	Generate and send a report. This API waits for the report to be generated before returning. We recommend that you set the client’s timeout to at least 60 seconds. Available to org admins only and with a valid license.

Only available in Grafana Enterprise v7.0+.
This API endpoint is experimental and may be deprecated in a future release. On deprecation, a migration strategy will be provided and the endpoint will remain functional until the next major release of Grafana.

You need to have a permission with action `reports:send`.
*/
func (a *Client) SendReport(params *SendReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SendReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sendReport",
		Method:             "POST",
		PathPattern:        "/reports/email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SendReportReader{formats: a.formats},
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
	success, ok := result.(*SendReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SendTestEmail sends test report via email

	Available to org admins only and with a valid license.

You need to have a permission with action `reports:send`.
*/
func (a *Client) SendTestEmail(params *SendTestEmailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SendTestEmailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendTestEmailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sendTestEmail",
		Method:             "POST",
		PathPattern:        "/reports/test-email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SendTestEmailReader{formats: a.formats},
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
	success, ok := result.(*SendTestEmailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendTestEmail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	UpdateReport updates a report

	Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports.admin:write` with scope `reports:id:<report ID>`.
*/
func (a *Client) UpdateReport(params *UpdateReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateReport",
		Method:             "PUT",
		PathPattern:        "/reports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateReportReader{formats: a.formats},
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
	success, ok := result.(*UpdateReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
