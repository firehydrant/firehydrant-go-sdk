// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MetricsMetricsEntitySortEntity struct {
	Field     *string `json:"field,omitempty"`
	Direction *string `json:"direction,omitempty"`
	Limit     *int    `json:"limit,omitempty"`
}

func (o *MetricsMetricsEntitySortEntity) GetField() *string {
	if o == nil {
		return nil
	}
	return o.Field
}

func (o *MetricsMetricsEntitySortEntity) GetDirection() *string {
	if o == nil {
		return nil
	}
	return o.Direction
}

func (o *MetricsMetricsEntitySortEntity) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}
