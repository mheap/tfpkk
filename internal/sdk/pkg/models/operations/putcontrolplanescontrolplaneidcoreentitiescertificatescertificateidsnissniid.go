// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDRequest struct {
	// ID of the Certificate to lookup
	CertificateID string `pathParam:"style=simple,explode=false,name=CertificateId"`
	// Description of the SNI
	CreateSNIWithoutParents shared.CreateSNIWithoutParents `request:"mediaType=application/json"`
	// ID of the SNI to lookup
	SNIID string `pathParam:"style=simple,explode=false,name=SNIId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDRequest) GetCertificateID() string {
	if o == nil {
		return ""
	}
	return o.CertificateID
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDRequest) GetCreateSNIWithoutParents() shared.CreateSNIWithoutParents {
	if o == nil {
		return shared.CreateSNIWithoutParents{}
	}
	return o.CreateSNIWithoutParents
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDRequest) GetSNIID() string {
	if o == nil {
		return ""
	}
	return o.SNIID
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponseBody - Invalid SNI
type PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponseBody struct {
}

type PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully upserted SNI
	Sni *shared.Sni
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid SNI
	Object *PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponseBody
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponse) GetSni() *shared.Sni {
	if o == nil {
		return nil
	}
	return o.Sni
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponse) GetObject() *PutControlPlanesControlPlaneIDCoreEntitiesCertificatesCertificateIDSnisSNIIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}