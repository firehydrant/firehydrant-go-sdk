// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateLifecycleMilestoneRequestBody struct {
	// The name of the milestone
	Name string `json:"name"`
	// A long-form description of the milestone's purpose
	Description string `json:"description"`
	// A unique identifier for the milestone. If not provided, one will be generated from the name.
	Slug *string `json:"slug,omitempty"`
	// The ID of the phase to which the milestone should belong
	PhaseID string `json:"phase_id"`
	// The position of the milestone within the phase. If not provided, the milestone will be added as the last milestone in the phase.
	Position *int `json:"position,omitempty"`
	// The ID of a later milestone that cannot be started until this milestone has a timestamp populated
	RequiredAtMilestoneID *string `json:"required_at_milestone_id,omitempty"`
}

func (o *CreateLifecycleMilestoneRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateLifecycleMilestoneRequestBody) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *CreateLifecycleMilestoneRequestBody) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateLifecycleMilestoneRequestBody) GetPhaseID() string {
	if o == nil {
		return ""
	}
	return o.PhaseID
}

func (o *CreateLifecycleMilestoneRequestBody) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *CreateLifecycleMilestoneRequestBody) GetRequiredAtMilestoneID() *string {
	if o == nil {
		return nil
	}
	return o.RequiredAtMilestoneID
}

type CreateLifecycleMilestoneResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a new milestone
	LifecyclesPhaseEntityList *components.LifecyclesPhaseEntityList
}

func (o *CreateLifecycleMilestoneResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateLifecycleMilestoneResponse) GetLifecyclesPhaseEntityList() *components.LifecyclesPhaseEntityList {
	if o == nil {
		return nil
	}
	return o.LifecyclesPhaseEntityList
}
