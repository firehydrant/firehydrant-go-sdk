// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type CreateNuncLinkRequest struct {
	NuncConnectionID string `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
}

func (o *CreateNuncLinkRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}
