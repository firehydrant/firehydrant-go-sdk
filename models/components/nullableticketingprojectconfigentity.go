// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// NullableTicketingProjectConfigEntityDetails - A config object containing details about the project config. Can be one of: Ticketing::JiraCloud::ProjectConfigEntity, Ticketing::JiraOnprem::ProjectConfigEntity, or Ticketing::Shortcut::ProjectConfigEntity
type NullableTicketingProjectConfigEntityDetails struct {
}

// NullableTicketingProjectConfigEntity - Ticketing_ProjectConfigEntity model
type NullableTicketingProjectConfigEntity struct {
	ID                   *string `json:"id,omitempty"`
	ConnectionID         *string `json:"connection_id,omitempty"`
	ConnectionType       *string `json:"connection_type,omitempty"`
	TicketingProjectID   *string `json:"ticketing_project_id,omitempty"`
	TicketingProjectName *string `json:"ticketing_project_name,omitempty"`
	// A config object containing details about the project config. Can be one of: Ticketing::JiraCloud::ProjectConfigEntity, Ticketing::JiraOnprem::ProjectConfigEntity, or Ticketing::Shortcut::ProjectConfigEntity
	Details *NullableTicketingProjectConfigEntityDetails `json:"details,omitempty"`
}

func (o *NullableTicketingProjectConfigEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NullableTicketingProjectConfigEntity) GetConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionID
}

func (o *NullableTicketingProjectConfigEntity) GetConnectionType() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionType
}

func (o *NullableTicketingProjectConfigEntity) GetTicketingProjectID() *string {
	if o == nil {
		return nil
	}
	return o.TicketingProjectID
}

func (o *NullableTicketingProjectConfigEntity) GetTicketingProjectName() *string {
	if o == nil {
		return nil
	}
	return o.TicketingProjectName
}

func (o *NullableTicketingProjectConfigEntity) GetDetails() *NullableTicketingProjectConfigEntityDetails {
	if o == nil {
		return nil
	}
	return o.Details
}
