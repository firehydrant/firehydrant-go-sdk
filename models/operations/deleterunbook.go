// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteRunbookRequest struct {
	RunbookID string `pathParam:"style=simple,explode=false,name=runbook_id"`
}

func (o *DeleteRunbookRequest) GetRunbookID() string {
	if o == nil {
		return ""
	}
	return o.RunbookID
}

type DeleteRunbookResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Delete a runbook and make it unavailable for any future incidents.
	RunbookEntity *components.RunbookEntity
}

func (o *DeleteRunbookResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteRunbookResponse) GetRunbookEntity() *components.RunbookEntity {
	if o == nil {
		return nil
	}
	return o.RunbookEntity
}
