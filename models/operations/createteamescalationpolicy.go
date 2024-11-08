// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateTeamEscalationPolicyRequest struct {
	TeamID                              string                                         `pathParam:"style=simple,explode=false,name=team_id"`
	PostV1TeamsTeamIDEscalationPolicies components.PostV1TeamsTeamIDEscalationPolicies `request:"mediaType=application/json"`
}

func (o *CreateTeamEscalationPolicyRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *CreateTeamEscalationPolicyRequest) GetPostV1TeamsTeamIDEscalationPolicies() components.PostV1TeamsTeamIDEscalationPolicies {
	if o == nil {
		return components.PostV1TeamsTeamIDEscalationPolicies{}
	}
	return o.PostV1TeamsTeamIDEscalationPolicies
}

type CreateTeamEscalationPolicyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *CreateTeamEscalationPolicyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}