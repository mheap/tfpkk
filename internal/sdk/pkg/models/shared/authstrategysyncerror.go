// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/utils"
)

type Details struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	Message              []string    `json:"message,omitempty"`
	Type                 *string     `json:"type,omitempty"`
}

func (d Details) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Details) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Details) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *Details) GetMessage() []string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *Details) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

type Info struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	Details              []Details   `json:"details,omitempty"`
}

func (i Info) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Info) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Info) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *Info) GetDetails() []Details {
	if o == nil {
		return nil
	}
	return o.Details
}

type Value string

const (
	ValuePluginSyncErrorComm               Value = "plugin_sync_error_comm"
	ValuePluginSyncErrorUnknown            Value = "plugin_sync_error_unknown"
	ValuePluginSyncErrorFatal              Value = "plugin_sync_error_fatal"
	ValuePluginSyncErrorUpdatingPluginRefs Value = "plugin_sync_error_updating_plugin_refs"
)

func (e Value) ToPointer() *Value {
	return &e
}

func (e *Value) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "plugin_sync_error_comm":
		fallthrough
	case "plugin_sync_error_unknown":
		fallthrough
	case "plugin_sync_error_fatal":
		fallthrough
	case "plugin_sync_error_updating_plugin_refs":
		*e = Value(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Value: %v", v)
	}
}

type AuthStrategySyncError struct {
	Info       *Info   `json:"info,omitempty"`
	Message    string  `json:"message"`
	PluginName *string `json:"plugin_name,omitempty"`
	Value      *Value  `json:"value,omitempty"`
}

func (o *AuthStrategySyncError) GetInfo() *Info {
	if o == nil {
		return nil
	}
	return o.Info
}

func (o *AuthStrategySyncError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *AuthStrategySyncError) GetPluginName() *string {
	if o == nil {
		return nil
	}
	return o.PluginName
}

func (o *AuthStrategySyncError) GetValue() *Value {
	if o == nil {
		return nil
	}
	return o.Value
}
