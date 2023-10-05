// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"konnect/internal/sdk/pkg/models/shared"
	"time"
)

func (r *APIProductDataSourceModel) RefreshFromGetResponse(resp *shared.APIProduct) {
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
	if resp.Description != nil {
		r.Description = types.StringValue(*resp.Description)
	} else {
		r.Description = types.StringNull()
	}
	r.ID = types.StringValue(resp.ID)
	if r.Labels == nil && len(resp.Labels) > 0 {
		r.Labels = make(map[string]types.String)
		for key, value := range resp.Labels {
			r.Labels[key] = types.StringValue(value)
		}
	}
	r.Name = types.StringValue(resp.Name)
	r.PortalIds = nil
	for _, v := range resp.PortalIds {
		r.PortalIds = append(r.PortalIds, types.StringValue(v))
	}
	r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
}