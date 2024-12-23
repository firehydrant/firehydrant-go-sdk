// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"firehydrant/internal/utils"
	"firehydrant/models/components"
	"fmt"
)

// QueryParamSort - The order to sort the transcript entries.
type QueryParamSort string

const (
	QueryParamSortAsc  QueryParamSort = "asc"
	QueryParamSortDesc QueryParamSort = "desc"
)

func (e QueryParamSort) ToPointer() *QueryParamSort {
	return &e
}
func (e *QueryParamSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = QueryParamSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamSort: %v", v)
	}
}

type GetIncidentTranscriptRequest struct {
	// The ID of the transcript entry to start after.
	After *string `queryParam:"style=form,explode=true,name=after"`
	// The ID of the transcript entry to start before.
	Before *string `queryParam:"style=form,explode=true,name=before"`
	// The order to sort the transcript entries.
	Sort       *QueryParamSort `default:"asc" queryParam:"style=form,explode=true,name=sort"`
	IncidentID string          `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (g GetIncidentTranscriptRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetIncidentTranscriptRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetIncidentTranscriptRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *GetIncidentTranscriptRequest) GetBefore() *string {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *GetIncidentTranscriptRequest) GetSort() *QueryParamSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetIncidentTranscriptRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type GetIncidentTranscriptResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve the transcript for a specific incident
	PublicAPIV1IncidentsTranscriptEntity *components.PublicAPIV1IncidentsTranscriptEntity
}

func (o *GetIncidentTranscriptResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetIncidentTranscriptResponse) GetPublicAPIV1IncidentsTranscriptEntity() *components.PublicAPIV1IncidentsTranscriptEntity {
	if o == nil {
		return nil
	}
	return o.PublicAPIV1IncidentsTranscriptEntity
}
