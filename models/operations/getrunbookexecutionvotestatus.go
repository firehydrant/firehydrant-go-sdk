// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetRunbookExecutionVoteStatusRequest struct {
	ExecutionID string `pathParam:"style=simple,explode=false,name=execution_id"`
}

func (o *GetRunbookExecutionVoteStatusRequest) GetExecutionID() string {
	if o == nil {
		return ""
	}
	return o.ExecutionID
}

type GetRunbookExecutionVoteStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns the current vote counts for an object
	VotesEntity *components.VotesEntity
}

func (o *GetRunbookExecutionVoteStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRunbookExecutionVoteStatusResponse) GetVotesEntity() *components.VotesEntity {
	if o == nil {
		return nil
	}
	return o.VotesEntity
}
