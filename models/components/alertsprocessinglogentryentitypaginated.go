// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AlertsProcessingLogEntryEntityPaginated - Alerts_ProcessingLogEntryEntityPaginated model
type AlertsProcessingLogEntryEntityPaginated struct {
	Data       []AlertsProcessingLogEntryEntity `json:"data,omitempty"`
	Pagination *NullablePaginationEntity        `json:"pagination,omitempty"`
}

func (o *AlertsProcessingLogEntryEntityPaginated) GetData() []AlertsProcessingLogEntryEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AlertsProcessingLogEntryEntityPaginated) GetPagination() *NullablePaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
