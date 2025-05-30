// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type UpdateTicketingFieldMapRequest struct {
	MapID              string `pathParam:"style=simple,explode=false,name=map_id"`
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
}

func (o *UpdateTicketingFieldMapRequest) GetMapID() string {
	if o == nil {
		return ""
	}
	return o.MapID
}

func (o *UpdateTicketingFieldMapRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}
