// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type ListScheduledMaintenancesRequest struct {
	// Filter scheduled_maintenances with a query on their name
	Query   *string `queryParam:"style=form,explode=true,name=query"`
	Page    *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListScheduledMaintenancesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListScheduledMaintenancesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListScheduledMaintenancesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}
