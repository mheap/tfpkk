// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateACLConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateACLConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateACL struct {
	Consumer *CreateACLConsumer `json:"consumer,omitempty"`
	Group    *string            `json:"group,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

func (o *CreateACL) GetConsumer() *CreateACLConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateACL) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *CreateACL) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
