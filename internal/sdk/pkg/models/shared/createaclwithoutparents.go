// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateACLWithoutParents struct {
	Group *string  `json:"group,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

func (o *CreateACLWithoutParents) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *CreateACLWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
