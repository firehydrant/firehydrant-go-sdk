// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type ListTeamsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// A query to search teams by their name or description
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// A query to search teams by their name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// A comma separated list of service IDs
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// Filter by teams that have or do not have members with a default incident role asssigned. Value may be 'present', 'blank', or the ID of an incident role.
	DefaultIncidentRole *string `queryParam:"style=form,explode=true,name=default_incident_role"`
	// Boolean to determine whether to return a slimified version of the teams object
	Lite *bool `queryParam:"style=form,explode=true,name=lite"`
}

func (o *ListTeamsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListTeamsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListTeamsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListTeamsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListTeamsRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *ListTeamsRequest) GetDefaultIncidentRole() *string {
	if o == nil {
		return nil
	}
	return o.DefaultIncidentRole
}

func (o *ListTeamsRequest) GetLite() *bool {
	if o == nil {
		return nil
	}
	return o.Lite
}
