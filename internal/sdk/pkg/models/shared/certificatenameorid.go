// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CertificateNameOrID - The unique identifier or the `name` attribute of the Certificate whose SNIs are to be retrieved. When using this endpoint, only SNIs associated to the specified Certificate will be listed.
type CertificateNameOrID string

const (
	CertificateNameOrIDA3ad71a866854b03A101980a953544f6 CertificateNameOrID = "a3ad71a8-6685-4b03-a101-980a953544f6"
	CertificateNameOrIDName                             CertificateNameOrID = "name"
)

func (e CertificateNameOrID) ToPointer() *CertificateNameOrID {
	return &e
}

func (e *CertificateNameOrID) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "a3ad71a8-6685-4b03-a101-980a953544f6":
		fallthrough
	case "name":
		*e = CertificateNameOrID(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CertificateNameOrID: %v", v)
	}
}
