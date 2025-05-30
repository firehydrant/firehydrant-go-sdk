// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IncidentsMilestoneEntityPaginated - Incidents_MilestoneEntityPaginated model
type IncidentsMilestoneEntityPaginated struct {
	Data       []IncidentsMilestoneEntity `json:"data,omitempty"`
	Pagination *NullablePaginationEntity  `json:"pagination,omitempty"`
}

func (o *IncidentsMilestoneEntityPaginated) GetData() []IncidentsMilestoneEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *IncidentsMilestoneEntityPaginated) GetPagination() *NullablePaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
