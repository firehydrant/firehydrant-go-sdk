// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// RetrospectivesIndexTemplateEntityPaginated - Retrospectives_IndexTemplateEntityPaginated model
type RetrospectivesIndexTemplateEntityPaginated struct {
	Data       []RetrospectivesIndexTemplateEntity `json:"data,omitempty"`
	Pagination *NullablePaginationEntity           `json:"pagination,omitempty"`
}

func (o *RetrospectivesIndexTemplateEntityPaginated) GetData() []RetrospectivesIndexTemplateEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *RetrospectivesIndexTemplateEntityPaginated) GetPagination() *NullablePaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
