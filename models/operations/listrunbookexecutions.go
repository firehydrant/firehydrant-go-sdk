// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListRunbookExecutionsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListRunbookExecutionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListRunbookExecutionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListRunbookExecutionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List all Runbook executions across all Runbooks
	RunbooksExecutionEntityPaginated *components.RunbooksExecutionEntityPaginated
}

func (o *ListRunbookExecutionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListRunbookExecutionsResponse) GetRunbooksExecutionEntityPaginated() *components.RunbooksExecutionEntityPaginated {
	if o == nil {
		return nil
	}
	return o.RunbooksExecutionEntityPaginated
}
