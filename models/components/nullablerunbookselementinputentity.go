// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type NullableRunbooksElementInputEntity struct {
	Label        *string `json:"label,omitempty"`
	Placeholder  *string `json:"placeholder,omitempty"`
	DefaultValue *string `json:"default_value,omitempty"`
	Required     *bool   `json:"required,omitempty"`
}

func (o *NullableRunbooksElementInputEntity) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *NullableRunbooksElementInputEntity) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *NullableRunbooksElementInputEntity) GetDefaultValue() *string {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *NullableRunbooksElementInputEntity) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}
