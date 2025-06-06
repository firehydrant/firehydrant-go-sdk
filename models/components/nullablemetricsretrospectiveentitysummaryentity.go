// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type NullableMetricsRetrospectiveEntitySummaryEntity struct {
	Completed            *int     `json:"completed,omitempty"`
	Total                *int     `json:"total,omitempty"`
	Incomplete           *int     `json:"incomplete,omitempty"`
	Mean                 *float32 `json:"mean,omitempty"`
	Shortest             *float32 `json:"shortest,omitempty"`
	Longest              *float32 `json:"longest,omitempty"`
	CompletionPercentage *float32 `json:"completion_percentage,omitempty"`
}

func (o *NullableMetricsRetrospectiveEntitySummaryEntity) GetCompleted() *int {
	if o == nil {
		return nil
	}
	return o.Completed
}

func (o *NullableMetricsRetrospectiveEntitySummaryEntity) GetTotal() *int {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *NullableMetricsRetrospectiveEntitySummaryEntity) GetIncomplete() *int {
	if o == nil {
		return nil
	}
	return o.Incomplete
}

func (o *NullableMetricsRetrospectiveEntitySummaryEntity) GetMean() *float32 {
	if o == nil {
		return nil
	}
	return o.Mean
}

func (o *NullableMetricsRetrospectiveEntitySummaryEntity) GetShortest() *float32 {
	if o == nil {
		return nil
	}
	return o.Shortest
}

func (o *NullableMetricsRetrospectiveEntitySummaryEntity) GetLongest() *float32 {
	if o == nil {
		return nil
	}
	return o.Longest
}

func (o *NullableMetricsRetrospectiveEntitySummaryEntity) GetCompletionPercentage() *float32 {
	if o == nil {
		return nil
	}
	return o.CompletionPercentage
}
