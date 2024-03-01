// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDRequest struct {
	// ID of the Certificate to lookup
	CertificateID string `pathParam:"style=simple,explode=false,name=CertificateId"`
	// ID of the SNI to lookup
	SNIID string `pathParam:"style=simple,explode=false,name=SNIId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDRequest) GetCertificateID() string {
	if o == nil {
		return ""
	}
	return o.CertificateID
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDRequest) GetSNIID() string {
	if o == nil {
		return ""
	}
	return o.SNIID
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully fetched SNI
	Sni *shared.Sni
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponse) GetSni() *shared.Sni {
	if o == nil {
		return nil
	}
	return o.Sni
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}