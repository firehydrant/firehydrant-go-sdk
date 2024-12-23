// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateAwsCloudTrailBatchRequest struct {
	ID                                        string                                               `pathParam:"style=simple,explode=false,name=id"`
	PatchV1IntegrationsAwsCloudtrailBatchesID components.PatchV1IntegrationsAwsCloudtrailBatchesID `request:"mediaType=application/json"`
}

func (o *UpdateAwsCloudTrailBatchRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateAwsCloudTrailBatchRequest) GetPatchV1IntegrationsAwsCloudtrailBatchesID() components.PatchV1IntegrationsAwsCloudtrailBatchesID {
	if o == nil {
		return components.PatchV1IntegrationsAwsCloudtrailBatchesID{}
	}
	return o.PatchV1IntegrationsAwsCloudtrailBatchesID
}

type UpdateAwsCloudTrailBatchResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a CloudTrail batch with new information.
	IntegrationsAwsCloudtrailBatchEntity *components.IntegrationsAwsCloudtrailBatchEntity
}

func (o *UpdateAwsCloudTrailBatchResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateAwsCloudTrailBatchResponse) GetIntegrationsAwsCloudtrailBatchEntity() *components.IntegrationsAwsCloudtrailBatchEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsAwsCloudtrailBatchEntity
}
