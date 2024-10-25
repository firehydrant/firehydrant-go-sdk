// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PatchV1PostMortemsReportsReportIDReasonsReasonIDRequest struct {
	ReportID                                         string                                                      `pathParam:"style=simple,explode=false,name=report_id"`
	ReasonID                                         string                                                      `pathParam:"style=simple,explode=false,name=reason_id"`
	PatchV1PostMortemsReportsReportIDReasonsReasonID components.PatchV1PostMortemsReportsReportIDReasonsReasonID `request:"mediaType=application/json"`
}

func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDRequest) GetReasonID() string {
	if o == nil {
		return ""
	}
	return o.ReasonID
}

func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDRequest) GetPatchV1PostMortemsReportsReportIDReasonsReasonID() components.PatchV1PostMortemsReportsReportIDReasonsReasonID {
	if o == nil {
		return components.PatchV1PostMortemsReportsReportIDReasonsReasonID{}
	}
	return o.PatchV1PostMortemsReportsReportIDReasonsReasonID
}

type PatchV1PostMortemsReportsReportIDReasonsReasonIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a contributing factor
	PostMortemsReasonEntity *components.PostMortemsReasonEntity
}

func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDResponse) GetPostMortemsReasonEntity() *components.PostMortemsReasonEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsReasonEntity
}