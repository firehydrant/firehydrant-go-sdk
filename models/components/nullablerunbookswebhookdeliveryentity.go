// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type NullableRunbooksWebhookDeliveryEntity struct {
	Headers    *string `json:"headers,omitempty"`
	StatusCode *string `json:"status_code,omitempty"`
}

func (o *NullableRunbooksWebhookDeliveryEntity) GetHeaders() *string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *NullableRunbooksWebhookDeliveryEntity) GetStatusCode() *string {
	if o == nil {
		return nil
	}
	return o.StatusCode
}
