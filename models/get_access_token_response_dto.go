// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetAccessTokenResponseDTO get access token response DTO
//
// swagger:model GetAccessTokenResponseDTO
type GetAccessTokenResponseDTO struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// expires at
	ExpiresAt string `json:"expiresAt,omitempty"`

	// first used at
	FirstUsedAt string `json:"firstUsedAt,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last used at
	LastUsedAt string `json:"lastUsedAt,omitempty"`
}

// Validate validates this get access token response DTO
func (m *GetAccessTokenResponseDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get access token response DTO based on context it is used
func (m *GetAccessTokenResponseDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetAccessTokenResponseDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAccessTokenResponseDTO) UnmarshalBinary(b []byte) error {
	var res GetAccessTokenResponseDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
