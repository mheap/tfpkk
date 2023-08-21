// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CreateCertificateRequestBody struct {
	// PEM-encoded public certificate chain of the SSL key pair.
	Cert string `json:"cert"`
	// PEM-encoded public certificate chain of the alternate SSL key pair.
	CertAlt *string `json:"cert_alt,omitempty"`
	// PEM-encoded private key of the SSL key pair.
	Key string `json:"key"`
	// PEM-encoded private key of the alternate SSL key pair.
	KeyAlt *string `json:"key_alt,omitempty"`
	// An optional set of strings associated with the Certificate for grouping and filtering.
	//
	Tags []string `json:"tags,omitempty"`
}

type CreateCertificateRequest struct {
	RequestBody *CreateCertificateRequestBody `request:"mediaType=application/json"`
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
}

// CreateCertificate400ApplicationJSON - Invalid Certificate
type CreateCertificate400ApplicationJSON struct {
}

// CreateCertificate200ApplicationJSONDataMetadata - metadata obtained from the certificate
type CreateCertificate200ApplicationJSONDataMetadata struct {
	// The expiration date of the certificate
	Expiry *string `json:"expiry,omitempty"`
	// Information about the certificate issuer
	Issuer *string `json:"issuer,omitempty"`
	// Key usage flags for the certificate
	KeyUsages []string `json:"key_usages,omitempty"`
	// Information about the certificate
	Subject *string `json:"subject,omitempty"`
}

type CreateCertificate200ApplicationJSONData struct {
	// PEM-encoded public certificate chain of the SSL key This field is referenceable and can be stored in a vault
	Cert *string `json:"cert,omitempty"`
	// Unix epoch when the resource was created
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The UUID representation of the certificate object.
	ID *string `json:"id,omitempty"`
	// PEM-encoded private key of the SSL key pair. This field is _referenceable_, which means it can be securely stored as a [secret](/gateway/latest/plan-and-deploy/security/secrets-management/getting-started) in a vault. References must follow a [specific format](/gateway/latest/plan-and-deploy/security/secrets-management/reference-format).
	Key *string `json:"key,omitempty"`
	// metadata obtained from the certificate
	Metadata *CreateCertificate200ApplicationJSONDataMetadata `json:"metadata,omitempty"`
	// An optional set of strings associated with the Certificate for grouping and filtering
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was updated
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

// CreateCertificate200ApplicationJSON - A certificate object represents a public certificate. These fields are _referenceable_.
type CreateCertificate200ApplicationJSON struct {
	// An array containing certificate information.
	Data []CreateCertificate200ApplicationJSONData `json:"data,omitempty"`
}

type CreateCertificateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A certificate object represents a public certificate. These fields are _referenceable_.
	CreateCertificate200ApplicationJSONObject *CreateCertificate200ApplicationJSON
	// Invalid Certificate
	CreateCertificate400ApplicationJSONObject *CreateCertificate400ApplicationJSON
}
