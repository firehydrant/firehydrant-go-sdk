// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type ListFieldMapAvailableFieldsRequest struct {
	FieldMapID string `pathParam:"style=simple,explode=false,name=field_map_id"`
}

func (o *ListFieldMapAvailableFieldsRequest) GetFieldMapID() string {
	if o == nil {
		return ""
	}
	return o.FieldMapID
}
