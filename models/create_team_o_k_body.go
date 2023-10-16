// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateTeamOKBody create team o k body
//
// swagger:model createTeamOKBody
type CreateTeamOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// team Id
	TeamID int64 `json:"teamId,omitempty"`
}

// Validate validates this create team o k body
func (m *CreateTeamOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create team o k body based on context it is used
func (m *CreateTeamOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateTeamOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTeamOKBody) UnmarshalBinary(b []byte) error {
	var res CreateTeamOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
