// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type ListPrioritiesRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListPrioritiesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListPrioritiesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}
