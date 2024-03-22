// Code generated by go-swagger; DO NOT EDIT.

package folders

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

// NewGetFoldersParams creates a new GetFoldersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFoldersParams() *GetFoldersParams {
	return &GetFoldersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFoldersParamsWithTimeout creates a new GetFoldersParams object
// with the ability to set a timeout on a request.
func NewGetFoldersParamsWithTimeout(timeout time.Duration) *GetFoldersParams {
	return &GetFoldersParams{
		timeout: timeout,
	}
}

// NewGetFoldersParamsWithContext creates a new GetFoldersParams object
// with the ability to set a context for a request.
func NewGetFoldersParamsWithContext(ctx context.Context) *GetFoldersParams {
	return &GetFoldersParams{
		Context: ctx,
	}
}

// NewGetFoldersParamsWithHTTPClient creates a new GetFoldersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFoldersParamsWithHTTPClient(client *http.Client) *GetFoldersParams {
	return &GetFoldersParams{
		HTTPClient: client,
	}
}

/*
GetFoldersParams contains all the parameters to send to the API endpoint

	for the get folders operation.

	Typically these are written to a http.Request.
*/
type GetFoldersParams struct {

	/* Limit.

	   Limit the maximum number of folders to return

	   Format: int64
	   Default: 1000
	*/
	Limit *int64

	/* Page.

	   Page index for starting fetching folders

	   Format: int64
	   Default: 1
	*/
	Page *int64

	/* ParentUID.

	   The parent folder UID
	*/
	ParentUID *string

	/* Permission.

	   Set to `Edit` to return folders that the user can edit

	   Default: "View"
	*/
	Permission *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get folders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFoldersParams) WithDefaults() *GetFoldersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get folders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFoldersParams) SetDefaults() {
	var (
		limitDefault = int64(1000)

		pageDefault = int64(1)

		permissionDefault = string("View")
	)

	val := GetFoldersParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		Permission: &permissionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get folders params
func (o *GetFoldersParams) WithTimeout(timeout time.Duration) *GetFoldersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get folders params
func (o *GetFoldersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get folders params
func (o *GetFoldersParams) WithContext(ctx context.Context) *GetFoldersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get folders params
func (o *GetFoldersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get folders params
func (o *GetFoldersParams) WithHTTPClient(client *http.Client) *GetFoldersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get folders params
func (o *GetFoldersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get folders params
func (o *GetFoldersParams) WithLimit(limit *int64) *GetFoldersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get folders params
func (o *GetFoldersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get folders params
func (o *GetFoldersParams) WithPage(page *int64) *GetFoldersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get folders params
func (o *GetFoldersParams) SetPage(page *int64) {
	o.Page = page
}

// WithParentUID adds the parentUID to the get folders params
func (o *GetFoldersParams) WithParentUID(parentUID *string) *GetFoldersParams {
	o.SetParentUID(parentUID)
	return o
}

// SetParentUID adds the parentUid to the get folders params
func (o *GetFoldersParams) SetParentUID(parentUID *string) {
	o.ParentUID = parentUID
}

// WithPermission adds the permission to the get folders params
func (o *GetFoldersParams) WithPermission(permission *string) *GetFoldersParams {
	o.SetPermission(permission)
	return o
}

// SetPermission adds the permission to the get folders params
func (o *GetFoldersParams) SetPermission(permission *string) {
	o.Permission = permission
}

// WriteToRequest writes these params to a swagger request
func (o *GetFoldersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.ParentUID != nil {

		// query param parentUid
		var qrParentUID string

		if o.ParentUID != nil {
			qrParentUID = *o.ParentUID
		}
		qParentUID := qrParentUID
		if qParentUID != "" {

			if err := r.SetQueryParam("parentUid", qParentUID); err != nil {
				return err
			}
		}
	}

	if o.Permission != nil {

		// query param permission
		var qrPermission string

		if o.Permission != nil {
			qrPermission = *o.Permission
		}
		qPermission := qrPermission
		if qPermission != "" {

			if err := r.SetQueryParam("permission", qPermission); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
