// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetMemberDefaultAudienceRequest struct {
	MemberID int `pathParam:"style=simple,explode=false,name=member_id"`
}

func (o *GetMemberDefaultAudienceRequest) GetMemberID() int {
	if o == nil {
		return 0
	}
	return o.MemberID
}
