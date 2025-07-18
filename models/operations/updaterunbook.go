// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

type UpdateRunbookRequest struct {
	RunbookID     string                   `pathParam:"style=simple,explode=false,name=runbook_id"`
	UpdateRunbook components.UpdateRunbook `request:"mediaType=application/json"`
}

func (o *UpdateRunbookRequest) GetRunbookID() string {
	if o == nil {
		return ""
	}
	return o.RunbookID
}

func (o *UpdateRunbookRequest) GetUpdateRunbook() components.UpdateRunbook {
	if o == nil {
		return components.UpdateRunbook{}
	}
	return o.UpdateRunbook
}
