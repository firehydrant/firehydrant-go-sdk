// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

type UpdateScimUserRequest struct {
	ID             string                    `pathParam:"style=simple,explode=false,name=id"`
	UpdateScimUser components.UpdateScimUser `request:"mediaType=application/scim+json"`
}

func (o *UpdateScimUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateScimUserRequest) GetUpdateScimUser() components.UpdateScimUser {
	if o == nil {
		return components.UpdateScimUser{}
	}
	return o.UpdateScimUser
}
