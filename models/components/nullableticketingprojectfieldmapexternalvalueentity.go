// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type NullableTicketingProjectFieldMapExternalValueEntityType string

const (
	NullableTicketingProjectFieldMapExternalValueEntityTypeLiteral     NullableTicketingProjectFieldMapExternalValueEntityType = "literal"
	NullableTicketingProjectFieldMapExternalValueEntityTypeFhAttribute NullableTicketingProjectFieldMapExternalValueEntityType = "fh-attribute"
	NullableTicketingProjectFieldMapExternalValueEntityTypeFhType      NullableTicketingProjectFieldMapExternalValueEntityType = "fh-type"
)

func (e NullableTicketingProjectFieldMapExternalValueEntityType) ToPointer() *NullableTicketingProjectFieldMapExternalValueEntityType {
	return &e
}
func (e *NullableTicketingProjectFieldMapExternalValueEntityType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "literal":
		fallthrough
	case "fh-attribute":
		fallthrough
	case "fh-type":
		*e = NullableTicketingProjectFieldMapExternalValueEntityType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NullableTicketingProjectFieldMapExternalValueEntityType: %v", v)
	}
}

type NullableTicketingProjectFieldMapExternalValueEntity struct {
	Type         *NullableTicketingProjectFieldMapExternalValueEntityType `json:"type,omitempty"`
	Value        *string                                                  `json:"value,omitempty"`
	Attribute    *string                                                  `json:"attribute,omitempty"`
	Presentation *string                                                  `json:"presentation,omitempty"`
}

func (o *NullableTicketingProjectFieldMapExternalValueEntity) GetType() *NullableTicketingProjectFieldMapExternalValueEntityType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *NullableTicketingProjectFieldMapExternalValueEntity) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *NullableTicketingProjectFieldMapExternalValueEntity) GetAttribute() *string {
	if o == nil {
		return nil
	}
	return o.Attribute
}

func (o *NullableTicketingProjectFieldMapExternalValueEntity) GetPresentation() *string {
	if o == nil {
		return nil
	}
	return o.Presentation
}
