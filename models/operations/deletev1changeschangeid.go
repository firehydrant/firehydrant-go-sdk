// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteV1ChangesChangeIDRequest struct {
	ChangeID string `pathParam:"style=simple,explode=false,name=change_id"`
}

func (o *DeleteV1ChangesChangeIDRequest) GetChangeID() string {
	if o == nil {
		return ""
	}
	return o.ChangeID
}

type DeleteV1ChangesChangeIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteV1ChangesChangeIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}