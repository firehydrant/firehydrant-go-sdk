// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListSignalsOnCallRequest struct {
	// An optional comma separated list of team IDs to filter currently on-call users by
	TeamID *string `queryParam:"style=form,explode=true,name=team_id"`
}

func (o *ListSignalsOnCallRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type ListSignalsOnCallResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ListSignalsOnCallResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
