// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreateServiceDependency - Creates a service dependency relationship between two services
type CreateServiceDependency struct {
	ServiceID          string `json:"service_id"`
	ConnectedServiceID string `json:"connected_service_id"`
	// A note to describe the service dependency relationship
	Notes *string `json:"notes,omitempty"`
}

func (o *CreateServiceDependency) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateServiceDependency) GetConnectedServiceID() string {
	if o == nil {
		return ""
	}
	return o.ConnectedServiceID
}

func (o *CreateServiceDependency) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}
