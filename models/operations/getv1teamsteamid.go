// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1TeamsTeamIDRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=team_id"`
	// Boolean to determine whether to return a slimified version of the teams object
	Lite *bool `queryParam:"style=form,explode=true,name=lite"`
}

func (o *GetV1TeamsTeamIDRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetV1TeamsTeamIDRequest) GetLite() *bool {
	if o == nil {
		return nil
	}
	return o.Lite
}

type GetV1TeamsTeamIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a single team from its ID
	TeamEntity *components.TeamEntity
}

func (o *GetV1TeamsTeamIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1TeamsTeamIDResponse) GetTeamEntity() *components.TeamEntity {
	if o == nil {
		return nil
	}
	return o.TeamEntity
}