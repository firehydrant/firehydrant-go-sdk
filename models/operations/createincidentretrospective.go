// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type CreateIncidentRetrospectiveRequestBody struct {
	// The id of the retrospective template to apply.
	RetrospectiveTemplateID string `json:"retrospective_template_id"`
}

func (o *CreateIncidentRetrospectiveRequestBody) GetRetrospectiveTemplateID() string {
	if o == nil {
		return ""
	}
	return o.RetrospectiveTemplateID
}

type CreateIncidentRetrospectiveRequest struct {
	IncidentID  string                                 `pathParam:"style=simple,explode=false,name=incident_id"`
	RequestBody CreateIncidentRetrospectiveRequestBody `request:"mediaType=application/json"`
}

func (o *CreateIncidentRetrospectiveRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *CreateIncidentRetrospectiveRequest) GetRequestBody() CreateIncidentRetrospectiveRequestBody {
	if o == nil {
		return CreateIncidentRetrospectiveRequestBody{}
	}
	return o.RequestBody
}
