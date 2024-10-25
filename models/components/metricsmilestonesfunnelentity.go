// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// MetricsMilestonesFunnelEntity - Metrics_MilestonesFunnelEntity model
type MetricsMilestonesFunnelEntity struct {
	Data      []MetricsMilestonesFunnelEntityDataBucketEntity `json:"data,omitempty"`
	Columns   []MetricsMilestonesFunnelEntityColumnEntity     `json:"columns,omitempty"`
	Groupings *MetricsMilestonesFunnelEntityGroupingsEntity   `json:"groupings,omitempty"`
	Meta      *MetricsMilestonesFunnelEntityMetaEntity        `json:"meta,omitempty"`
}

func (o *MetricsMilestonesFunnelEntity) GetData() []MetricsMilestonesFunnelEntityDataBucketEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *MetricsMilestonesFunnelEntity) GetColumns() []MetricsMilestonesFunnelEntityColumnEntity {
	if o == nil {
		return nil
	}
	return o.Columns
}

func (o *MetricsMilestonesFunnelEntity) GetGroupings() *MetricsMilestonesFunnelEntityGroupingsEntity {
	if o == nil {
		return nil
	}
	return o.Groupings
}

func (o *MetricsMilestonesFunnelEntity) GetMeta() *MetricsMilestonesFunnelEntityMetaEntity {
	if o == nil {
		return nil
	}
	return o.Meta
}