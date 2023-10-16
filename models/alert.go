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

// Alert Alert has info for an alert.
//
// swagger:model Alert
type Alert struct {

	// active at
	// Format: date-time
	ActiveAt strfmt.DateTime `json:"activeAt,omitempty"`

	// annotations
	// Required: true
	Annotations OverrideLabels `json:"annotations"`

	// labels
	// Required: true
	Labels OverrideLabels `json:"labels"`

	// state
	// Required: true
	State *string `json:"state"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this alert
func (m *Alert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Alert) validateActiveAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ActiveAt) { // not required
		return nil
	}

	if err := validate.FormatOf("activeAt", "body", "date-time", m.ActiveAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Alert) validateAnnotations(formats strfmt.Registry) error {

	if err := validate.Required("annotations", "body", m.Annotations); err != nil {
		return err
	}

	if m.Annotations != nil {
		if err := m.Annotations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("annotations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("annotations")
			}
			return err
		}
	}

	return nil
}

func (m *Alert) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("labels", "body", m.Labels); err != nil {
		return err
	}

	if m.Labels != nil {
		if err := m.Labels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

func (m *Alert) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Alert) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this alert based on the context it is used
func (m *Alert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAnnotations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Alert) contextValidateAnnotations(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Annotations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("annotations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("annotations")
		}
		return err
	}

	return nil
}

func (m *Alert) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("labels")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Alert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alert) UnmarshalBinary(b []byte) error {
	var res Alert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
