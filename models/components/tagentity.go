// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// TagEntity model
type TagEntity struct {
	Name *string `json:"name,omitempty"`
}

func (o *TagEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
