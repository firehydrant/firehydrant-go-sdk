// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type ListServiceAvailableUpstreamDependenciesRequest struct {
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *ListServiceAvailableUpstreamDependenciesRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}
