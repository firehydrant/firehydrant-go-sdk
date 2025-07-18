// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
)

type GenerateAudienceSummaryRequestBody struct {
	// Whether to force regeneration of the summary
	ForceRegenerate *bool `default:"true" json:"force_regenerate"`
}

func (g GenerateAudienceSummaryRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GenerateAudienceSummaryRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GenerateAudienceSummaryRequestBody) GetForceRegenerate() *bool {
	if o == nil {
		return nil
	}
	return o.ForceRegenerate
}

type GenerateAudienceSummaryRequest struct {
	// Unique identifier of the audience
	AudienceID string `pathParam:"style=simple,explode=false,name=audience_id"`
	// Unique identifier of the incident to summarize
	IncidentID  string                              `pathParam:"style=simple,explode=false,name=incident_id"`
	RequestBody *GenerateAudienceSummaryRequestBody `request:"mediaType=application/json"`
}

func (o *GenerateAudienceSummaryRequest) GetAudienceID() string {
	if o == nil {
		return ""
	}
	return o.AudienceID
}

func (o *GenerateAudienceSummaryRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *GenerateAudienceSummaryRequest) GetRequestBody() *GenerateAudienceSummaryRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}
