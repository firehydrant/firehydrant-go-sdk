// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateStatusPageComponentGroupRequestBody struct {
	Name             *string `json:"name,omitempty"`
	ComponentGroupID *string `json:"component_group_id,omitempty"`
	Position         *int    `json:"position,omitempty"`
}

func (o *UpdateStatusPageComponentGroupRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateStatusPageComponentGroupRequestBody) GetComponentGroupID() *string {
	if o == nil {
		return nil
	}
	return o.ComponentGroupID
}

func (o *UpdateStatusPageComponentGroupRequestBody) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}

type UpdateStatusPageComponentGroupRequest struct {
	NuncConnectionID string                                     `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
	GroupID          string                                     `pathParam:"style=simple,explode=false,name=group_id"`
	RequestBody      *UpdateStatusPageComponentGroupRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateStatusPageComponentGroupRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}

func (o *UpdateStatusPageComponentGroupRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

func (o *UpdateStatusPageComponentGroupRequest) GetRequestBody() *UpdateStatusPageComponentGroupRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateStatusPageComponentGroupResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateStatusPageComponentGroupResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}