// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MetricsMilestonesFunnelEntityGroupingsEntity struct {
	// The bucket size for the data
	BucketSize *string `json:"bucket_size,omitempty"`
}

func (o *MetricsMilestonesFunnelEntityGroupingsEntity) GetBucketSize() *string {
	if o == nil {
		return nil
	}
	return o.BucketSize
}