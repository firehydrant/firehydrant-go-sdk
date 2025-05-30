// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DeleteTranscriptEntryRequest struct {
	TranscriptID string `pathParam:"style=simple,explode=false,name=transcript_id"`
	IncidentID   string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *DeleteTranscriptEntryRequest) GetTranscriptID() string {
	if o == nil {
		return ""
	}
	return o.TranscriptID
}

func (o *DeleteTranscriptEntryRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}
