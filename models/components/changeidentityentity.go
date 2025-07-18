// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
	"time"
)

// ChangeIdentityEntity model
type ChangeIdentityEntity struct {
	ID        *string    `json:"id,omitempty"`
	Type      *string    `json:"type,omitempty"`
	Value     *string    `json:"value,omitempty"`
	ChangeID  *string    `json:"change_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (c ChangeIdentityEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ChangeIdentityEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ChangeIdentityEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChangeIdentityEntity) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ChangeIdentityEntity) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *ChangeIdentityEntity) GetChangeID() *string {
	if o == nil {
		return nil
	}
	return o.ChangeID
}

func (o *ChangeIdentityEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ChangeIdentityEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
