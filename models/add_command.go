// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddCommand add command
//
// swagger:model AddCommand
type AddCommand struct {

	// name
	Name string `json:"name,omitempty"`

	// role
	// Enum: [Viewer Editor Admin]
	Role string `json:"role,omitempty"`

	// seconds to live
	SecondsToLive int64 `json:"secondsToLive,omitempty"`
}

// Validate validates this add command
func (m *AddCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addCommandTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Viewer","Editor","Admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addCommandTypeRolePropEnum = append(addCommandTypeRolePropEnum, v)
	}
}

const (

	// AddCommandRoleViewer captures enum value "Viewer"
	AddCommandRoleViewer string = "Viewer"

	// AddCommandRoleEditor captures enum value "Editor"
	AddCommandRoleEditor string = "Editor"

	// AddCommandRoleAdmin captures enum value "Admin"
	AddCommandRoleAdmin string = "Admin"
)

// prop value enum
func (m *AddCommand) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addCommandTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AddCommand) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add command based on context it is used
func (m *AddCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddCommand) UnmarshalBinary(b []byte) error {
	var res AddCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
