// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateBasicAuthConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateBasicAuthConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateBasicAuth struct {
	Consumer *CreateBasicAuthConsumer `json:"consumer,omitempty"`
	Password *string                  `json:"password,omitempty"`
	Tags     []string                 `json:"tags,omitempty"`
	Username *string                  `json:"username,omitempty"`
}

func (o *CreateBasicAuth) GetConsumer() *CreateBasicAuthConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateBasicAuth) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *CreateBasicAuth) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateBasicAuth) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}
