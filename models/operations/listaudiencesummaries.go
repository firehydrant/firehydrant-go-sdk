// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type ListAudienceSummariesRequest struct {
	// Unique identifier of the incident to summarize
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *ListAudienceSummariesRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}
