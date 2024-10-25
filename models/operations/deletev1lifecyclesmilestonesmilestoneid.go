// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteV1LifecyclesMilestonesMilestoneIDRequest struct {
	MilestoneID string `pathParam:"style=simple,explode=false,name=milestone_id"`
}

func (o *DeleteV1LifecyclesMilestonesMilestoneIDRequest) GetMilestoneID() string {
	if o == nil {
		return ""
	}
	return o.MilestoneID
}

type DeleteV1LifecyclesMilestonesMilestoneIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Delete a milestone
	LifecyclesPhaseEntity *components.LifecyclesPhaseEntity
}

func (o *DeleteV1LifecyclesMilestonesMilestoneIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1LifecyclesMilestonesMilestoneIDResponse) GetLifecyclesPhaseEntity() *components.LifecyclesPhaseEntity {
	if o == nil {
		return nil
	}
	return o.LifecyclesPhaseEntity
}