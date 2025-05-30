// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type ListPostMortemReasonsRequest struct {
	ReportID string `pathParam:"style=simple,explode=false,name=report_id"`
	Page     *int   `queryParam:"style=form,explode=true,name=page"`
	PerPage  *int   `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListPostMortemReasonsRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *ListPostMortemReasonsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListPostMortemReasonsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}
