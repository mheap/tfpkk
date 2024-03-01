// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateSNICertificate - The id (a UUID) of the certificate with which to associate the SNI hostname. The Certificate must have a valid private key associated with it to be used by the SNI object.
type CreateSNICertificate struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateSNICertificate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateSNI - An SNI object represents a many-to-one mapping of hostnames to a certificate. That is, a certificate object can have many hostnames associated with it; when Kong receives an SSL request, it uses the SNI field in the Client Hello to lookup the certificate object based on the SNI associated with the certificate.
type CreateSNI struct {
	// The id (a UUID) of the certificate with which to associate the SNI hostname. The Certificate must have a valid private key associated with it to be used by the SNI object.
	Certificate CreateSNICertificate `json:"certificate"`
	// The SNI name to associate with the given certificate.
	Name string `json:"name"`
	// An optional set of strings associated with the SNIs for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (o *CreateSNI) GetCertificate() CreateSNICertificate {
	if o == nil {
		return CreateSNICertificate{}
	}
	return o.Certificate
}

func (o *CreateSNI) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateSNI) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
