// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type ListInfrastructuresRequest struct {
	// A query to search infrastructures by their name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Omit for any infrastructure that is added to an incident using the format "incident/{incident_id}"
	OmitFor *string `queryParam:"style=form,explode=true,name=omit_for"`
	// Restrict infrastructure search to given type.
	Type    *string `queryParam:"style=form,explode=true,name=type"`
	Page    *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListInfrastructuresRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListInfrastructuresRequest) GetOmitFor() *string {
	if o == nil {
		return nil
	}
	return o.OmitFor
}

func (o *ListInfrastructuresRequest) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListInfrastructuresRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListInfrastructuresRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}
