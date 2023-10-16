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

// CorrelationConfig correlation config
//
// swagger:model CorrelationConfig
type CorrelationConfig struct {

	// Field used to attach the correlation link
	// Example: message
	// Required: true
	Field *string `json:"field"`

	// Target data query
	// Example: {"expr":"job=app"}
	// Required: true
	Target interface{} `json:"target"`

	// transformations
	Transformations Transformations `json:"transformations,omitempty"`

	// type
	// Required: true
	Type *CorrelationConfigType `json:"type"`
}

// Validate validates this correlation config
func (m *CorrelationConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateField(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransformations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CorrelationConfig) validateField(formats strfmt.Registry) error {

	if err := validate.Required("field", "body", m.Field); err != nil {
		return err
	}

	return nil
}

func (m *CorrelationConfig) validateTarget(formats strfmt.Registry) error {

	if m.Target == nil {
		return errors.Required("target", "body", nil)
	}

	return nil
}

func (m *CorrelationConfig) validateTransformations(formats strfmt.Registry) error {
	if swag.IsZero(m.Transformations) { // not required
		return nil
	}

	if err := m.Transformations.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transformations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transformations")
		}
		return err
	}

	return nil
}

func (m *CorrelationConfig) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this correlation config based on the context it is used
func (m *CorrelationConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransformations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CorrelationConfig) contextValidateTransformations(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Transformations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transformations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transformations")
		}
		return err
	}

	return nil
}

func (m *CorrelationConfig) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CorrelationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CorrelationConfig) UnmarshalBinary(b []byte) error {
	var res CorrelationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
