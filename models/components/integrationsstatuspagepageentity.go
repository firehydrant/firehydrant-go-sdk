// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IntegrationsStatuspagePageEntity - Integrations_Statuspage_PageEntity model
type IntegrationsStatuspagePageEntity struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *IntegrationsStatuspagePageEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IntegrationsStatuspagePageEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
