// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreateSignalsWebhookTarget - Create a Signals webhook target.
type CreateSignalsWebhookTarget struct {
	// The webhook target's name.
	Name string `json:"name"`
	// An optional detailed description of the webhook target.
	Description *string `json:"description,omitempty"`
	// The URL that the webhook target will notify.
	URL string `json:"url"`
	// An optional secret we will provide in the `FH-Signature` header
	// when sending a payload to the webhook target. This key will not be
	// shown in any response once configured.
	//
	SigningKey *string `json:"signing_key,omitempty"`
}

func (o *CreateSignalsWebhookTarget) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateSignalsWebhookTarget) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateSignalsWebhookTarget) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *CreateSignalsWebhookTarget) GetSigningKey() *string {
	if o == nil {
		return nil
	}
	return o.SigningKey
}
