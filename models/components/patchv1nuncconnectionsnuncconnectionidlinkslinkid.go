// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PatchV1NuncConnectionsNuncConnectionIDLinksLinkID - Update a link to be displayed on a FireHydrant status page
type PatchV1NuncConnectionsNuncConnectionIDLinksLinkID struct {
	DisplayText *string `json:"display_text,omitempty"`
	IconURL     *string `json:"icon_url,omitempty"`
	HrefURL     *string `json:"href_url,omitempty"`
}

func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkID) GetDisplayText() *string {
	if o == nil {
		return nil
	}
	return o.DisplayText
}

func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkID) GetIconURL() *string {
	if o == nil {
		return nil
	}
	return o.IconURL
}

func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkID) GetHrefURL() *string {
	if o == nil {
		return nil
	}
	return o.HrefURL
}