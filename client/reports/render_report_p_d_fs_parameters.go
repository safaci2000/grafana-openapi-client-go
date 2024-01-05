// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// NewRenderReportPDFsParams creates a new RenderReportPDFsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRenderReportPDFsParams() *RenderReportPDFsParams {
	return &RenderReportPDFsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRenderReportPDFsParamsWithTimeout creates a new RenderReportPDFsParams object
// with the ability to set a timeout on a request.
func NewRenderReportPDFsParamsWithTimeout(timeout time.Duration) *RenderReportPDFsParams {
	return &RenderReportPDFsParams{
		timeout: timeout,
	}
}

// NewRenderReportPDFsParamsWithContext creates a new RenderReportPDFsParams object
// with the ability to set a context for a request.
func NewRenderReportPDFsParamsWithContext(ctx context.Context) *RenderReportPDFsParams {
	return &RenderReportPDFsParams{
		Context: ctx,
	}
}

// NewRenderReportPDFsParamsWithHTTPClient creates a new RenderReportPDFsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRenderReportPDFsParamsWithHTTPClient(client *http.Client) *RenderReportPDFsParams {
	return &RenderReportPDFsParams{
		HTTPClient: client,
	}
}

/*
RenderReportPDFsParams contains all the parameters to send to the API endpoint

	for the render report p d fs operation.

	Typically these are written to a http.Request.
*/
type RenderReportPDFsParams struct {

	// DashboardID.
	DashboardID *string

	// IncludeTables.
	IncludeTables *string

	// Layout.
	Layout *string

	// Orientation.
	Orientation *string

	// ScaleFactor.
	ScaleFactor *string

	// Title.
	Title *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the render report p d fs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenderReportPDFsParams) WithDefaults() *RenderReportPDFsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the render report p d fs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenderReportPDFsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the render report p d fs params
func (o *RenderReportPDFsParams) WithTimeout(timeout time.Duration) *RenderReportPDFsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the render report p d fs params
func (o *RenderReportPDFsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the render report p d fs params
func (o *RenderReportPDFsParams) WithContext(ctx context.Context) *RenderReportPDFsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the render report p d fs params
func (o *RenderReportPDFsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the render report p d fs params
func (o *RenderReportPDFsParams) WithHTTPClient(client *http.Client) *RenderReportPDFsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the render report p d fs params
func (o *RenderReportPDFsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardID adds the dashboardID to the render report p d fs params
func (o *RenderReportPDFsParams) WithDashboardID(dashboardID *string) *RenderReportPDFsParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the render report p d fs params
func (o *RenderReportPDFsParams) SetDashboardID(dashboardID *string) {
	o.DashboardID = dashboardID
}

// WithIncludeTables adds the includeTables to the render report p d fs params
func (o *RenderReportPDFsParams) WithIncludeTables(includeTables *string) *RenderReportPDFsParams {
	o.SetIncludeTables(includeTables)
	return o
}

// SetIncludeTables adds the includeTables to the render report p d fs params
func (o *RenderReportPDFsParams) SetIncludeTables(includeTables *string) {
	o.IncludeTables = includeTables
}

// WithLayout adds the layout to the render report p d fs params
func (o *RenderReportPDFsParams) WithLayout(layout *string) *RenderReportPDFsParams {
	o.SetLayout(layout)
	return o
}

// SetLayout adds the layout to the render report p d fs params
func (o *RenderReportPDFsParams) SetLayout(layout *string) {
	o.Layout = layout
}

// WithOrientation adds the orientation to the render report p d fs params
func (o *RenderReportPDFsParams) WithOrientation(orientation *string) *RenderReportPDFsParams {
	o.SetOrientation(orientation)
	return o
}

// SetOrientation adds the orientation to the render report p d fs params
func (o *RenderReportPDFsParams) SetOrientation(orientation *string) {
	o.Orientation = orientation
}

// WithScaleFactor adds the scaleFactor to the render report p d fs params
func (o *RenderReportPDFsParams) WithScaleFactor(scaleFactor *string) *RenderReportPDFsParams {
	o.SetScaleFactor(scaleFactor)
	return o
}

// SetScaleFactor adds the scaleFactor to the render report p d fs params
func (o *RenderReportPDFsParams) SetScaleFactor(scaleFactor *string) {
	o.ScaleFactor = scaleFactor
}

// WithTitle adds the title to the render report p d fs params
func (o *RenderReportPDFsParams) WithTitle(title *string) *RenderReportPDFsParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the render report p d fs params
func (o *RenderReportPDFsParams) SetTitle(title *string) {
	o.Title = title
}

// WriteToRequest writes these params to a swagger request
func (o *RenderReportPDFsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DashboardID != nil {

		// query param dashboardID
		var qrDashboardID string

		if o.DashboardID != nil {
			qrDashboardID = *o.DashboardID
		}
		qDashboardID := qrDashboardID
		if qDashboardID != "" {

			if err := r.SetQueryParam("dashboardID", qDashboardID); err != nil {
				return err
			}
		}
	}

	if o.IncludeTables != nil {

		// query param includeTables
		var qrIncludeTables string

		if o.IncludeTables != nil {
			qrIncludeTables = *o.IncludeTables
		}
		qIncludeTables := qrIncludeTables
		if qIncludeTables != "" {

			if err := r.SetQueryParam("includeTables", qIncludeTables); err != nil {
				return err
			}
		}
	}

	if o.Layout != nil {

		// query param layout
		var qrLayout string

		if o.Layout != nil {
			qrLayout = *o.Layout
		}
		qLayout := qrLayout
		if qLayout != "" {

			if err := r.SetQueryParam("layout", qLayout); err != nil {
				return err
			}
		}
	}

	if o.Orientation != nil {

		// query param orientation
		var qrOrientation string

		if o.Orientation != nil {
			qrOrientation = *o.Orientation
		}
		qOrientation := qrOrientation
		if qOrientation != "" {

			if err := r.SetQueryParam("orientation", qOrientation); err != nil {
				return err
			}
		}
	}

	if o.ScaleFactor != nil {

		// query param scaleFactor
		var qrScaleFactor string

		if o.ScaleFactor != nil {
			qrScaleFactor = *o.ScaleFactor
		}
		qScaleFactor := qrScaleFactor
		if qScaleFactor != "" {

			if err := r.SetQueryParam("scaleFactor", qScaleFactor); err != nil {
				return err
			}
		}
	}

	if o.Title != nil {

		// query param title
		var qrTitle string

		if o.Title != nil {
			qrTitle = *o.Title
		}
		qTitle := qrTitle
		if qTitle != "" {

			if err := r.SetQueryParam("title", qTitle); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
