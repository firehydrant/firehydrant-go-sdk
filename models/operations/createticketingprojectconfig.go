// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type CreateTicketingProjectConfigRequest struct {
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
}

func (o *CreateTicketingProjectConfigRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}
