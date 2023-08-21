// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Konnect/internal/sdk/pkg/models/shared"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetAuditLogWebhook200ApplicationJSONLogFormat - The output format of each log messages.
type GetAuditLogWebhook200ApplicationJSONLogFormat string

const (
	GetAuditLogWebhook200ApplicationJSONLogFormatCef  GetAuditLogWebhook200ApplicationJSONLogFormat = "cef"
	GetAuditLogWebhook200ApplicationJSONLogFormatJSON GetAuditLogWebhook200ApplicationJSONLogFormat = "json"
)

func (e GetAuditLogWebhook200ApplicationJSONLogFormat) ToPointer() *GetAuditLogWebhook200ApplicationJSONLogFormat {
	return &e
}

func (e *GetAuditLogWebhook200ApplicationJSONLogFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cef":
		fallthrough
	case "json":
		*e = GetAuditLogWebhook200ApplicationJSONLogFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAuditLogWebhook200ApplicationJSONLogFormat: %v", v)
	}
}

// GetAuditLogWebhook200ApplicationJSON - Get response for audit log webhook
type GetAuditLogWebhook200ApplicationJSON struct {
	// Indicates whether audit data should be sent to the webhook.
	Enabled *bool `json:"enabled,omitempty"`
	// The endpoint that will receive audit log messages.
	Endpoint *string `json:"endpoint,omitempty"`
	// The output format of each log messages.
	LogFormat *GetAuditLogWebhook200ApplicationJSONLogFormat `json:"log_format,omitempty"`
	// Timestamp when this webhook was last updated. Initial value is 0001-01-01T00:00:0Z.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type GetAuditLogWebhookResponse struct {
	ContentType string
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	StatusCode     int
	RawResponse    *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Get response for audit log webhook
	GetAuditLogWebhook200ApplicationJSONObject *GetAuditLogWebhook200ApplicationJSON
}
