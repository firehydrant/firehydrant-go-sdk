// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1NuncConnectionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Lists the information displayed as part of your FireHydrant hosted status pages.
	NuncConnectionEntityPaginated *components.NuncConnectionEntityPaginated
}

func (o *GetV1NuncConnectionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1NuncConnectionsResponse) GetNuncConnectionEntityPaginated() *components.NuncConnectionEntityPaginated {
	if o == nil {
		return nil
	}
	return o.NuncConnectionEntityPaginated
}