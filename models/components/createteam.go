// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreateTeamMsTeamsChannel - MS Teams channel identity for channel associated with this team
type CreateTeamMsTeamsChannel struct {
	ChannelID string `json:"channel_id"`
	MsTeamID  string `json:"ms_team_id"`
}

func (o *CreateTeamMsTeamsChannel) GetChannelID() string {
	if o == nil {
		return ""
	}
	return o.ChannelID
}

func (o *CreateTeamMsTeamsChannel) GetMsTeamID() string {
	if o == nil {
		return ""
	}
	return o.MsTeamID
}

type CreateTeamMembership struct {
	UserID     *string `json:"user_id,omitempty"`
	ScheduleID *string `json:"schedule_id,omitempty"`
	// An incident role ID that this user will automatically assigned if this team is assigned to an incident
	IncidentRoleID *string `json:"incident_role_id,omitempty"`
}

func (o *CreateTeamMembership) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *CreateTeamMembership) GetScheduleID() *string {
	if o == nil {
		return nil
	}
	return o.ScheduleID
}

func (o *CreateTeamMembership) GetIncidentRoleID() *string {
	if o == nil {
		return nil
	}
	return o.IncidentRoleID
}

// CreateTeam - Create a new team
type CreateTeam struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	// The Slack channel ID associated with this team. This may be the reference in FireHydrant's system (i.e. UUID) or the ID value from Slack (e.g. C1234567890).
	//
	SlackChannelID *string `json:"slack_channel_id,omitempty"`
	// MS Teams channel identity for channel associated with this team
	MsTeamsChannel *CreateTeamMsTeamsChannel `json:"ms_teams_channel,omitempty"`
	Memberships    []CreateTeamMembership    `json:"memberships,omitempty"`
	// A list of email addresses to invite to join the organization and automatically add to this team. If an email already has a pending invitation, the team will be added to their existing invitation.
	InviteEmails []string `json:"invite_emails,omitempty"`
}

func (o *CreateTeam) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateTeam) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateTeam) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateTeam) GetSlackChannelID() *string {
	if o == nil {
		return nil
	}
	return o.SlackChannelID
}

func (o *CreateTeam) GetMsTeamsChannel() *CreateTeamMsTeamsChannel {
	if o == nil {
		return nil
	}
	return o.MsTeamsChannel
}

func (o *CreateTeam) GetMemberships() []CreateTeamMembership {
	if o == nil {
		return nil
	}
	return o.Memberships
}

func (o *CreateTeam) GetInviteEmails() []string {
	if o == nil {
		return nil
	}
	return o.InviteEmails
}
