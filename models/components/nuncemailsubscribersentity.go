// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"firehydrant/internal/utils"
	"time"
)

// NuncEmailSubscribersEntity model
type NuncEmailSubscribersEntity struct {
	// UUID of the subscriber
	ID *string `json:"id,omitempty"`
	// Email of the subscriber
	Email *string `json:"email,omitempty"`
	// The time the subscriber was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

func (n NuncEmailSubscribersEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NuncEmailSubscribersEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *NuncEmailSubscribersEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NuncEmailSubscribersEntity) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *NuncEmailSubscribersEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}