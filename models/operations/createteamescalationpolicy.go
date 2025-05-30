// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateTeamEscalationPolicyRequest struct {
	TeamID                     string                                `pathParam:"style=simple,explode=false,name=team_id"`
	CreateTeamEscalationPolicy components.CreateTeamEscalationPolicy `request:"mediaType=application/json"`
}

func (o *CreateTeamEscalationPolicyRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *CreateTeamEscalationPolicyRequest) GetCreateTeamEscalationPolicy() components.CreateTeamEscalationPolicy {
	if o == nil {
		return components.CreateTeamEscalationPolicy{}
	}
	return o.CreateTeamEscalationPolicy
}
