// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetPostMortemReportRequest struct {
	ReportID string `pathParam:"style=simple,explode=false,name=report_id"`
}

func (o *GetPostMortemReportRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}
