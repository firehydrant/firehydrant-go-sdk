// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type NullableSuccinctEntity struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *NullableSuccinctEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NullableSuccinctEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
