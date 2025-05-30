// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type NullableRunbooksElementDynamicSelectEntity struct {
	Label        *string                                                       `json:"label,omitempty"`
	Placeholder  *string                                                       `json:"placeholder,omitempty"`
	AsyncURL     *string                                                       `json:"async_url,omitempty"`
	Required     *bool                                                         `json:"required,omitempty"`
	Clearable    *bool                                                         `json:"clearable,omitempty"`
	IsMulti      *bool                                                         `json:"is_multi,omitempty"`
	DefaultValue *NullableRunbooksElementDynamicSelectEntitySelectOptionEntity `json:"default_value,omitempty"`
	Options      []RunbooksElementDynamicSelectEntitySelectOptionEntity        `json:"options,omitempty"`
}

func (o *NullableRunbooksElementDynamicSelectEntity) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *NullableRunbooksElementDynamicSelectEntity) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *NullableRunbooksElementDynamicSelectEntity) GetAsyncURL() *string {
	if o == nil {
		return nil
	}
	return o.AsyncURL
}

func (o *NullableRunbooksElementDynamicSelectEntity) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *NullableRunbooksElementDynamicSelectEntity) GetClearable() *bool {
	if o == nil {
		return nil
	}
	return o.Clearable
}

func (o *NullableRunbooksElementDynamicSelectEntity) GetIsMulti() *bool {
	if o == nil {
		return nil
	}
	return o.IsMulti
}

func (o *NullableRunbooksElementDynamicSelectEntity) GetDefaultValue() *NullableRunbooksElementDynamicSelectEntitySelectOptionEntity {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *NullableRunbooksElementDynamicSelectEntity) GetOptions() []RunbooksElementDynamicSelectEntitySelectOptionEntity {
	if o == nil {
		return nil
	}
	return o.Options
}
