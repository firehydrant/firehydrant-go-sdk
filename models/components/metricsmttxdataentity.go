// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// MetricsMttxDataEntity - Metrics_MttxDataEntity model
type MetricsMttxDataEntity struct {
	Groupings []MetricsMttxDataEntityGroupingEntity `json:"groupings,omitempty"`
	Data      []MetricsMttxGroupEntity              `json:"data,omitempty"`
}

func (o *MetricsMttxDataEntity) GetGroupings() []MetricsMttxDataEntityGroupingEntity {
	if o == nil {
		return nil
	}
	return o.Groupings
}

func (o *MetricsMttxDataEntity) GetData() []MetricsMttxGroupEntity {
	if o == nil {
		return nil
	}
	return o.Data
}
