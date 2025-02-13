// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateIncidentRoleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a new incident role
	IncidentRoleEntity *components.IncidentRoleEntity
}

func (o *CreateIncidentRoleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateIncidentRoleResponse) GetIncidentRoleEntity() *components.IncidentRoleEntity {
	if o == nil {
		return nil
	}
	return o.IncidentRoleEntity
}
