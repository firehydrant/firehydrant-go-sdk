// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type ListIncidentTagsRequest struct {
	Prefix *string `queryParam:"style=form,explode=true,name=prefix"`
}

func (o *ListIncidentTagsRequest) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}
