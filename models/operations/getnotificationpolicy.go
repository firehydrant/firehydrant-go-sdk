// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetNotificationPolicyRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetNotificationPolicyRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
