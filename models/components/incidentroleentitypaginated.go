// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IncidentRoleEntityPaginated model
type IncidentRoleEntityPaginated struct {
	Data       []IncidentRoleEntity      `json:"data,omitempty"`
	Pagination *NullablePaginationEntity `json:"pagination,omitempty"`
}

func (o *IncidentRoleEntityPaginated) GetData() []IncidentRoleEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *IncidentRoleEntityPaginated) GetPagination() *NullablePaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
