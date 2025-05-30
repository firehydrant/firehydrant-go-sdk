// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ChangeIdentityEntityPaginated model
type ChangeIdentityEntityPaginated struct {
	Data       []ChangeIdentityEntity    `json:"data,omitempty"`
	Pagination *NullablePaginationEntity `json:"pagination,omitempty"`
}

func (o *ChangeIdentityEntityPaginated) GetData() []ChangeIdentityEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ChangeIdentityEntityPaginated) GetPagination() *NullablePaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
