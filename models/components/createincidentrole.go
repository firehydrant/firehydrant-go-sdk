// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreateIncidentRole - Create a new incident role
type CreateIncidentRole struct {
	Name        string  `json:"name"`
	Summary     string  `json:"summary"`
	Description *string `json:"description,omitempty"`
}

func (o *CreateIncidentRole) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateIncidentRole) GetSummary() string {
	if o == nil {
		return ""
	}
	return o.Summary
}

func (o *CreateIncidentRole) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
