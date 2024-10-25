// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PostV1ServiceDependenciesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Creates a service dependency relationship between two services
	ServiceDependencyEntity *components.ServiceDependencyEntity
}

func (o *PostV1ServiceDependenciesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1ServiceDependenciesResponse) GetServiceDependencyEntity() *components.ServiceDependencyEntity {
	if o == nil {
		return nil
	}
	return o.ServiceDependencyEntity
}