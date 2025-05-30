// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// UpdateIncidentLink - Update the external incident link attributes
type UpdateIncidentLink struct {
	DisplayText *string `json:"display_text,omitempty"`
	IconURL     *string `json:"icon_url,omitempty"`
	HrefURL     *string `json:"href_url,omitempty"`
}

func (o *UpdateIncidentLink) GetDisplayText() *string {
	if o == nil {
		return nil
	}
	return o.DisplayText
}

func (o *UpdateIncidentLink) GetIconURL() *string {
	if o == nil {
		return nil
	}
	return o.IconURL
}

func (o *UpdateIncidentLink) GetHrefURL() *string {
	if o == nil {
		return nil
	}
	return o.HrefURL
}
