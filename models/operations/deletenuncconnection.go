// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DeleteNuncConnectionRequest struct {
	NuncConnectionID string `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
}

func (o *DeleteNuncConnectionRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}
