// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1RunbooksExecutionsExecutionIDVotesStatusRequest struct {
	ExecutionID string `pathParam:"style=simple,explode=false,name=execution_id"`
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusRequest) GetExecutionID() string {
	if o == nil {
		return ""
	}
	return o.ExecutionID
}

type GetV1RunbooksExecutionsExecutionIDVotesStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns the current vote counts for an object
	VotesEntity *components.VotesEntity
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusResponse) GetVotesEntity() *components.VotesEntity {
	if o == nil {
		return nil
	}
	return o.VotesEntity
}