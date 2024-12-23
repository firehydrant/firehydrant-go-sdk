// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetSlackWorkspacesRequest struct {
	// Connection UUID
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *GetSlackWorkspacesRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type GetSlackWorkspacesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetSlackWorkspacesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
