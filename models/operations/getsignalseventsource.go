// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetSignalsEventSourceRequest struct {
	TransposerSlug string `pathParam:"style=simple,explode=false,name=transposer_slug"`
}

func (o *GetSignalsEventSourceRequest) GetTransposerSlug() string {
	if o == nil {
		return ""
	}
	return o.TransposerSlug
}
