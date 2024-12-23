// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateRunbookExecutionStepVotesRequest struct {
	ExecutionID                                          string                                                          `pathParam:"style=simple,explode=false,name=execution_id"`
	StepID                                               string                                                          `pathParam:"style=simple,explode=false,name=step_id"`
	PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotes components.PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotes `request:"mediaType=application/json"`
}

func (o *UpdateRunbookExecutionStepVotesRequest) GetExecutionID() string {
	if o == nil {
		return ""
	}
	return o.ExecutionID
}

func (o *UpdateRunbookExecutionStepVotesRequest) GetStepID() string {
	if o == nil {
		return ""
	}
	return o.StepID
}

func (o *UpdateRunbookExecutionStepVotesRequest) GetPatchV1RunbooksExecutionsExecutionIDStepsStepIDVotes() components.PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotes {
	if o == nil {
		return components.PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotes{}
	}
	return o.PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotes
}

type UpdateRunbookExecutionStepVotesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Allows for upvoting or downvoting an event
	VotesEntity *components.VotesEntity
}

func (o *UpdateRunbookExecutionStepVotesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateRunbookExecutionStepVotesResponse) GetVotesEntity() *components.VotesEntity {
	if o == nil {
		return nil
	}
	return o.VotesEntity
}
