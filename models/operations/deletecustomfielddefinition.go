// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DeleteCustomFieldDefinitionRequest struct {
	FieldID string `pathParam:"style=simple,explode=false,name=field_id"`
}

func (o *DeleteCustomFieldDefinitionRequest) GetFieldID() string {
	if o == nil {
		return ""
	}
	return o.FieldID
}
