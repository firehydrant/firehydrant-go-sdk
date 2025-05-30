// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DeleteIncidentAlertRequest struct {
	IncidentAlertID string `pathParam:"style=simple,explode=false,name=incident_alert_id"`
	IncidentID      string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *DeleteIncidentAlertRequest) GetIncidentAlertID() string {
	if o == nil {
		return ""
	}
	return o.IncidentAlertID
}

func (o *DeleteIncidentAlertRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}
