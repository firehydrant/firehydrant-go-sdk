// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type PublicAPIV1IncidentsSuccinctEntity struct {
	ID     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Number *int    `json:"number,omitempty"`
}

func (o *PublicAPIV1IncidentsSuccinctEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PublicAPIV1IncidentsSuccinctEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PublicAPIV1IncidentsSuccinctEntity) GetNumber() *int {
	if o == nil {
		return nil
	}
	return o.Number
}
