// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreateTicket - Creates a ticket for a project
type CreateTicket struct {
	Summary string `json:"summary"`
	// Which incident this ticket is related to, in the format of 'incident/UUID'
	RelatedTo   *string `json:"related_to,omitempty"`
	ProjectID   *string `json:"project_id,omitempty"`
	Description *string `json:"description,omitempty"`
	State       *string `json:"state,omitempty"`
	Type        *string `json:"type,omitempty"`
	PriorityID  *string `json:"priority_id,omitempty"`
	// List of tags for the ticket
	TagList []string `json:"tag_list,omitempty"`
	// The remote URL for an existing ticket in a supported and configured ticketing integration
	RemoteURL *string `json:"remote_url,omitempty"`
}

func (o *CreateTicket) GetSummary() string {
	if o == nil {
		return ""
	}
	return o.Summary
}

func (o *CreateTicket) GetRelatedTo() *string {
	if o == nil {
		return nil
	}
	return o.RelatedTo
}

func (o *CreateTicket) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *CreateTicket) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateTicket) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *CreateTicket) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CreateTicket) GetPriorityID() *string {
	if o == nil {
		return nil
	}
	return o.PriorityID
}

func (o *CreateTicket) GetTagList() []string {
	if o == nil {
		return nil
	}
	return o.TagList
}

func (o *CreateTicket) GetRemoteURL() *string {
	if o == nil {
		return nil
	}
	return o.RemoteURL
}
