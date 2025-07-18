// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

type CreateServiceChecklistResponseRequest struct {
	ServiceID                      string                                    `pathParam:"style=simple,explode=false,name=service_id"`
	ChecklistID                    string                                    `pathParam:"style=simple,explode=false,name=checklist_id"`
	CreateServiceChecklistResponse components.CreateServiceChecklistResponse `request:"mediaType=application/json"`
}

func (o *CreateServiceChecklistResponseRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateServiceChecklistResponseRequest) GetChecklistID() string {
	if o == nil {
		return ""
	}
	return o.ChecklistID
}

func (o *CreateServiceChecklistResponseRequest) GetCreateServiceChecklistResponse() components.CreateServiceChecklistResponse {
	if o == nil {
		return components.CreateServiceChecklistResponse{}
	}
	return o.CreateServiceChecklistResponse
}
