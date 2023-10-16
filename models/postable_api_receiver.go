// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostableAPIReceiver postable Api receiver
//
// swagger:model PostableApiReceiver
type PostableAPIReceiver struct {

	// discord configs
	DiscordConfigs []*DiscordConfig `json:"discord_configs"`

	// email configs
	EmailConfigs []*EmailConfig `json:"email_configs"`

	// grafana managed receiver configs
	GrafanaManagedReceiverConfigs []*PostableGrafanaReceiver `json:"grafana_managed_receiver_configs"`

	// A unique identifier for this receiver.
	Name string `json:"name,omitempty"`

	// opsgenie configs
	OpsgenieConfigs []*OpsGenieConfig `json:"opsgenie_configs"`

	// pagerduty configs
	PagerdutyConfigs []*PagerdutyConfig `json:"pagerduty_configs"`

	// pushover configs
	PushoverConfigs []*PushoverConfig `json:"pushover_configs"`

	// slack configs
	SlackConfigs []*SlackConfig `json:"slack_configs"`

	// sns configs
	SnsConfigs []*SNSConfig `json:"sns_configs"`

	// telegram configs
	TelegramConfigs []*TelegramConfig `json:"telegram_configs"`

	// victorops configs
	VictoropsConfigs []*VictorOpsConfig `json:"victorops_configs"`

	// webex configs
	WebexConfigs []*WebexConfig `json:"webex_configs"`

	// webhook configs
	WebhookConfigs []*WebhookConfig `json:"webhook_configs"`

	// wechat configs
	WechatConfigs []*WechatConfig `json:"wechat_configs"`
}

// Validate validates this postable Api receiver
func (m *PostableAPIReceiver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscordConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrafanaManagedReceiverConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpsgenieConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagerdutyConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePushoverConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlackConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnsConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelegramConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVictoropsConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebexConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWechatConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostableAPIReceiver) validateDiscordConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.DiscordConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.DiscordConfigs); i++ {
		if swag.IsZero(m.DiscordConfigs[i]) { // not required
			continue
		}

		if m.DiscordConfigs[i] != nil {
			if err := m.DiscordConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discord_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("discord_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateEmailConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.EmailConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.EmailConfigs); i++ {
		if swag.IsZero(m.EmailConfigs[i]) { // not required
			continue
		}

		if m.EmailConfigs[i] != nil {
			if err := m.EmailConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("email_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("email_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateGrafanaManagedReceiverConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.GrafanaManagedReceiverConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.GrafanaManagedReceiverConfigs); i++ {
		if swag.IsZero(m.GrafanaManagedReceiverConfigs[i]) { // not required
			continue
		}

		if m.GrafanaManagedReceiverConfigs[i] != nil {
			if err := m.GrafanaManagedReceiverConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateOpsgenieConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.OpsgenieConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.OpsgenieConfigs); i++ {
		if swag.IsZero(m.OpsgenieConfigs[i]) { // not required
			continue
		}

		if m.OpsgenieConfigs[i] != nil {
			if err := m.OpsgenieConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("opsgenie_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("opsgenie_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validatePagerdutyConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.PagerdutyConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.PagerdutyConfigs); i++ {
		if swag.IsZero(m.PagerdutyConfigs[i]) { // not required
			continue
		}

		if m.PagerdutyConfigs[i] != nil {
			if err := m.PagerdutyConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pagerduty_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pagerduty_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validatePushoverConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.PushoverConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.PushoverConfigs); i++ {
		if swag.IsZero(m.PushoverConfigs[i]) { // not required
			continue
		}

		if m.PushoverConfigs[i] != nil {
			if err := m.PushoverConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pushover_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pushover_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateSlackConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.SlackConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.SlackConfigs); i++ {
		if swag.IsZero(m.SlackConfigs[i]) { // not required
			continue
		}

		if m.SlackConfigs[i] != nil {
			if err := m.SlackConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("slack_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("slack_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateSnsConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.SnsConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.SnsConfigs); i++ {
		if swag.IsZero(m.SnsConfigs[i]) { // not required
			continue
		}

		if m.SnsConfigs[i] != nil {
			if err := m.SnsConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sns_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sns_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateTelegramConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.TelegramConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.TelegramConfigs); i++ {
		if swag.IsZero(m.TelegramConfigs[i]) { // not required
			continue
		}

		if m.TelegramConfigs[i] != nil {
			if err := m.TelegramConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("telegram_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("telegram_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateVictoropsConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.VictoropsConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.VictoropsConfigs); i++ {
		if swag.IsZero(m.VictoropsConfigs[i]) { // not required
			continue
		}

		if m.VictoropsConfigs[i] != nil {
			if err := m.VictoropsConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("victorops_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("victorops_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateWebexConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.WebexConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.WebexConfigs); i++ {
		if swag.IsZero(m.WebexConfigs[i]) { // not required
			continue
		}

		if m.WebexConfigs[i] != nil {
			if err := m.WebexConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webex_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webex_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateWebhookConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.WebhookConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.WebhookConfigs); i++ {
		if swag.IsZero(m.WebhookConfigs[i]) { // not required
			continue
		}

		if m.WebhookConfigs[i] != nil {
			if err := m.WebhookConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webhook_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webhook_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateWechatConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.WechatConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.WechatConfigs); i++ {
		if swag.IsZero(m.WechatConfigs[i]) { // not required
			continue
		}

		if m.WechatConfigs[i] != nil {
			if err := m.WechatConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wechat_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wechat_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this postable Api receiver based on the context it is used
func (m *PostableAPIReceiver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiscordConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmailConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrafanaManagedReceiverConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpsgenieConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagerdutyConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePushoverConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlackConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnsConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTelegramConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVictoropsConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebexConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebhookConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWechatConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostableAPIReceiver) contextValidateDiscordConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DiscordConfigs); i++ {

		if m.DiscordConfigs[i] != nil {

			if swag.IsZero(m.DiscordConfigs[i]) { // not required
				return nil
			}

			if err := m.DiscordConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discord_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("discord_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateEmailConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EmailConfigs); i++ {

		if m.EmailConfigs[i] != nil {

			if swag.IsZero(m.EmailConfigs[i]) { // not required
				return nil
			}

			if err := m.EmailConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("email_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("email_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateGrafanaManagedReceiverConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GrafanaManagedReceiverConfigs); i++ {

		if m.GrafanaManagedReceiverConfigs[i] != nil {

			if swag.IsZero(m.GrafanaManagedReceiverConfigs[i]) { // not required
				return nil
			}

			if err := m.GrafanaManagedReceiverConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateOpsgenieConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OpsgenieConfigs); i++ {

		if m.OpsgenieConfigs[i] != nil {

			if swag.IsZero(m.OpsgenieConfigs[i]) { // not required
				return nil
			}

			if err := m.OpsgenieConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("opsgenie_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("opsgenie_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidatePagerdutyConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PagerdutyConfigs); i++ {

		if m.PagerdutyConfigs[i] != nil {

			if swag.IsZero(m.PagerdutyConfigs[i]) { // not required
				return nil
			}

			if err := m.PagerdutyConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pagerduty_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pagerduty_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidatePushoverConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PushoverConfigs); i++ {

		if m.PushoverConfigs[i] != nil {

			if swag.IsZero(m.PushoverConfigs[i]) { // not required
				return nil
			}

			if err := m.PushoverConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pushover_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pushover_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateSlackConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SlackConfigs); i++ {

		if m.SlackConfigs[i] != nil {

			if swag.IsZero(m.SlackConfigs[i]) { // not required
				return nil
			}

			if err := m.SlackConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("slack_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("slack_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateSnsConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SnsConfigs); i++ {

		if m.SnsConfigs[i] != nil {

			if swag.IsZero(m.SnsConfigs[i]) { // not required
				return nil
			}

			if err := m.SnsConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sns_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sns_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateTelegramConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TelegramConfigs); i++ {

		if m.TelegramConfigs[i] != nil {

			if swag.IsZero(m.TelegramConfigs[i]) { // not required
				return nil
			}

			if err := m.TelegramConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("telegram_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("telegram_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateVictoropsConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VictoropsConfigs); i++ {

		if m.VictoropsConfigs[i] != nil {

			if swag.IsZero(m.VictoropsConfigs[i]) { // not required
				return nil
			}

			if err := m.VictoropsConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("victorops_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("victorops_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateWebexConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WebexConfigs); i++ {

		if m.WebexConfigs[i] != nil {

			if swag.IsZero(m.WebexConfigs[i]) { // not required
				return nil
			}

			if err := m.WebexConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webex_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webex_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateWebhookConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WebhookConfigs); i++ {

		if m.WebhookConfigs[i] != nil {

			if swag.IsZero(m.WebhookConfigs[i]) { // not required
				return nil
			}

			if err := m.WebhookConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webhook_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webhook_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateWechatConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WechatConfigs); i++ {

		if m.WechatConfigs[i] != nil {

			if swag.IsZero(m.WechatConfigs[i]) { // not required
				return nil
			}

			if err := m.WechatConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wechat_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wechat_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostableAPIReceiver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostableAPIReceiver) UnmarshalBinary(b []byte) error {
	var res PostableAPIReceiver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
