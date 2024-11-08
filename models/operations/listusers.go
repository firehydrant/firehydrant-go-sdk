// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListUsersRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// Text string of a query to filter users by name or email
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Text string of a query to filter users by name
	Name *string `queryParam:"style=form,explode=true,name=name"`
}

func (o *ListUsersRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListUsersRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListUsersRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListUsersRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ListUsersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a list of all users in an organization
	UserEntityPaginated *components.UserEntityPaginated
}

func (o *ListUsersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListUsersResponse) GetUserEntityPaginated() *components.UserEntityPaginated {
	if o == nil {
		return nil
	}
	return o.UserEntityPaginated
}