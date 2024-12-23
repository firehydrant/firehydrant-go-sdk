// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetScimUserRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetScimUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetScimUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetScimUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
