// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MetricsTicketFunnelMetricsEntityGroupingsEntity struct {
	// The bucket size for the data
	BucketSize *string `json:"bucket_size,omitempty"`
}

func (o *MetricsTicketFunnelMetricsEntityGroupingsEntity) GetBucketSize() *string {
	if o == nil {
		return nil
	}
	return o.BucketSize
}