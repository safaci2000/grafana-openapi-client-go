// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProvisionedAlertRule provisioned alert rule
//
// swagger:model ProvisionedAlertRule
type ProvisionedAlertRule struct {

	// annotations
	// Example: {"runbook_url":"https://supercoolrunbook.com/page/13"}
	Annotations map[string]string `json:"annotations,omitempty"`

	// condition
	// Example: A
	// Required: true
	Condition *string `json:"condition"`

	// data
	// Example: [{"datasourceUid":"__expr__","model":{"conditions":[{"evaluator":{"params":[0,0],"type":"gt"},"operator":{"type":"and"},"query":{"params":[]},"reducer":{"params":[],"type":"avg"},"type":"query"}],"datasource":{"type":"__expr__","uid":"__expr__"},"expression":"1 == 1","hide":false,"intervalMs":1000,"maxDataPoints":43200,"refId":"A","type":"math"},"queryType":"","refId":"A","relativeTimeRange":{"from":0,"to":0}}]
	// Required: true
	Data []*AlertQuery `json:"data"`

	// exec err state
	// Required: true
	// Enum: [OK Alerting Error]
	ExecErrState *string `json:"execErrState"`

	// folder UID
	// Example: project_x
	// Required: true
	FolderUID *string `json:"folderUID"`

	// for
	// Required: true
	For *Duration `json:"for"`

	// id
	ID int64 `json:"id,omitempty"`

	// is paused
	// Example: false
	IsPaused bool `json:"isPaused,omitempty"`

	// labels
	// Example: {"team":"sre-team-1"}
	Labels map[string]string `json:"labels,omitempty"`

	// no data state
	// Required: true
	// Enum: [Alerting NoData OK]
	NoDataState *string `json:"noDataState"`

	// org ID
	// Required: true
	OrgID *int64 `json:"orgID"`

	// provenance
	Provenance Provenance `json:"provenance,omitempty"`

	// rule group
	// Example: eval_group_1
	// Required: true
	// Max Length: 190
	// Min Length: 1
	RuleGroup *string `json:"ruleGroup"`

	// title
	// Example: Always firing
	// Required: true
	// Max Length: 190
	// Min Length: 1
	Title *string `json:"title"`

	// uid
	UID string `json:"uid,omitempty"`

	// updated
	// Read Only: true
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this provisioned alert rule
func (m *ProvisionedAlertRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecErrState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFolderUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoDataState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrgID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvenance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvisionedAlertRule) validateCondition(formats strfmt.Registry) error {

	if err := validate.Required("condition", "body", m.Condition); err != nil {
		return err
	}

	return nil
}

func (m *ProvisionedAlertRule) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var provisionedAlertRuleTypeExecErrStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OK","Alerting","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		provisionedAlertRuleTypeExecErrStatePropEnum = append(provisionedAlertRuleTypeExecErrStatePropEnum, v)
	}
}

const (

	// ProvisionedAlertRuleExecErrStateOK captures enum value "OK"
	ProvisionedAlertRuleExecErrStateOK string = "OK"

	// ProvisionedAlertRuleExecErrStateAlerting captures enum value "Alerting"
	ProvisionedAlertRuleExecErrStateAlerting string = "Alerting"

	// ProvisionedAlertRuleExecErrStateError captures enum value "Error"
	ProvisionedAlertRuleExecErrStateError string = "Error"
)

// prop value enum
func (m *ProvisionedAlertRule) validateExecErrStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, provisionedAlertRuleTypeExecErrStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProvisionedAlertRule) validateExecErrState(formats strfmt.Registry) error {

	if err := validate.Required("execErrState", "body", m.ExecErrState); err != nil {
		return err
	}

	// value enum
	if err := m.validateExecErrStateEnum("execErrState", "body", *m.ExecErrState); err != nil {
		return err
	}

	return nil
}

func (m *ProvisionedAlertRule) validateFolderUID(formats strfmt.Registry) error {

	if err := validate.Required("folderUID", "body", m.FolderUID); err != nil {
		return err
	}

	return nil
}

func (m *ProvisionedAlertRule) validateFor(formats strfmt.Registry) error {

	if err := validate.Required("for", "body", m.For); err != nil {
		return err
	}

	if err := validate.Required("for", "body", m.For); err != nil {
		return err
	}

	if m.For != nil {
		if err := m.For.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("for")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("for")
			}
			return err
		}
	}

	return nil
}

var provisionedAlertRuleTypeNoDataStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Alerting","NoData","OK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		provisionedAlertRuleTypeNoDataStatePropEnum = append(provisionedAlertRuleTypeNoDataStatePropEnum, v)
	}
}

const (

	// ProvisionedAlertRuleNoDataStateAlerting captures enum value "Alerting"
	ProvisionedAlertRuleNoDataStateAlerting string = "Alerting"

	// ProvisionedAlertRuleNoDataStateNoData captures enum value "NoData"
	ProvisionedAlertRuleNoDataStateNoData string = "NoData"

	// ProvisionedAlertRuleNoDataStateOK captures enum value "OK"
	ProvisionedAlertRuleNoDataStateOK string = "OK"
)

// prop value enum
func (m *ProvisionedAlertRule) validateNoDataStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, provisionedAlertRuleTypeNoDataStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProvisionedAlertRule) validateNoDataState(formats strfmt.Registry) error {

	if err := validate.Required("noDataState", "body", m.NoDataState); err != nil {
		return err
	}

	// value enum
	if err := m.validateNoDataStateEnum("noDataState", "body", *m.NoDataState); err != nil {
		return err
	}

	return nil
}

func (m *ProvisionedAlertRule) validateOrgID(formats strfmt.Registry) error {

	if err := validate.Required("orgID", "body", m.OrgID); err != nil {
		return err
	}

	return nil
}

func (m *ProvisionedAlertRule) validateProvenance(formats strfmt.Registry) error {
	if swag.IsZero(m.Provenance) { // not required
		return nil
	}

	if err := m.Provenance.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provenance")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provenance")
		}
		return err
	}

	return nil
}

func (m *ProvisionedAlertRule) validateRuleGroup(formats strfmt.Registry) error {

	if err := validate.Required("ruleGroup", "body", m.RuleGroup); err != nil {
		return err
	}

	if err := validate.MinLength("ruleGroup", "body", *m.RuleGroup, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("ruleGroup", "body", *m.RuleGroup, 190); err != nil {
		return err
	}

	return nil
}

func (m *ProvisionedAlertRule) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 190); err != nil {
		return err
	}

	return nil
}

func (m *ProvisionedAlertRule) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this provisioned alert rule based on the context it is used
func (m *ProvisionedAlertRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvenance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvisionedAlertRule) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {

			if swag.IsZero(m.Data[i]) { // not required
				return nil
			}

			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProvisionedAlertRule) contextValidateFor(ctx context.Context, formats strfmt.Registry) error {

	if m.For != nil {

		if err := m.For.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("for")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("for")
			}
			return err
		}
	}

	return nil
}

func (m *ProvisionedAlertRule) contextValidateProvenance(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Provenance) { // not required
		return nil
	}

	if err := m.Provenance.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provenance")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provenance")
		}
		return err
	}

	return nil
}

func (m *ProvisionedAlertRule) contextValidateUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated", "body", strfmt.DateTime(m.Updated)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvisionedAlertRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvisionedAlertRule) UnmarshalBinary(b []byte) error {
	var res ProvisionedAlertRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
