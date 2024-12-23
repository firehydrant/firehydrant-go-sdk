// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteSeverityMatrixConditionRequest struct {
	ConditionID string `pathParam:"style=simple,explode=false,name=condition_id"`
}

func (o *DeleteSeverityMatrixConditionRequest) GetConditionID() string {
	if o == nil {
		return ""
	}
	return o.ConditionID
}

type DeleteSeverityMatrixConditionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Delete a specific condition
	SeverityMatrixConditionEntity *components.SeverityMatrixConditionEntity
}

func (o *DeleteSeverityMatrixConditionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteSeverityMatrixConditionResponse) GetSeverityMatrixConditionEntity() *components.SeverityMatrixConditionEntity {
	if o == nil {
		return nil
	}
	return o.SeverityMatrixConditionEntity
}
