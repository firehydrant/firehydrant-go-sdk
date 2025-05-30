// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DeleteTicketingProjectConfigRequest struct {
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
	ConfigID           string `pathParam:"style=simple,explode=false,name=config_id"`
}

func (o *DeleteTicketingProjectConfigRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

func (o *DeleteTicketingProjectConfigRequest) GetConfigID() string {
	if o == nil {
		return ""
	}
	return o.ConfigID
}
