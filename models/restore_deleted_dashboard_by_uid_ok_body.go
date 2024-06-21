// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestoreDeletedDashboardByUIDOKBody restore deleted dashboard by Uid Ok body
//
// swagger:model restoreDeletedDashboardByUidOkBody
type RestoreDeletedDashboardByUIDOKBody struct {

	// FolderUID The unique identifier (uid) of the folder the dashboard belongs to.
	FolderUID string `json:"folderUid,omitempty"`

	// ID The unique identifier (id) of the created/updated dashboard.
	// Example: 1
	// Required: true
	ID *int64 `json:"id"`

	// Status status of the response.
	// Example: success
	// Required: true
	Status *string `json:"status"`

	// Slug The slug of the dashboard.
	// Example: my-dashboard
	// Required: true
	Title *string `json:"title"`

	// UID The unique identifier (uid) of the created/updated dashboard.
	// Example: nHz3SXiiz
	// Required: true
	UID *string `json:"uid"`

	// URL The relative URL for accessing the created/updated dashboard.
	// Example: /d/nHz3SXiiz/my-dashboard
	// Required: true
	URL *string `json:"url"`

	// Version The version of the dashboard.
	// Example: 2
	// Required: true
	Version *int64 `json:"version"`
}

// Validate validates this restore deleted dashboard by Uid Ok body
func (m *RestoreDeletedDashboardByUIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreDeletedDashboardByUIDOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RestoreDeletedDashboardByUIDOKBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *RestoreDeletedDashboardByUIDOKBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *RestoreDeletedDashboardByUIDOKBody) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	return nil
}

func (m *RestoreDeletedDashboardByUIDOKBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *RestoreDeletedDashboardByUIDOKBody) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this restore deleted dashboard by Uid Ok body based on context it is used
func (m *RestoreDeletedDashboardByUIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoreDeletedDashboardByUIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreDeletedDashboardByUIDOKBody) UnmarshalBinary(b []byte) error {
	var res RestoreDeletedDashboardByUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
