// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
	"time"
)

// TaskListEntity model
type TaskListEntity struct {
	ID            *string               `json:"id,omitempty"`
	Name          *string               `json:"name,omitempty"`
	Description   *string               `json:"description,omitempty"`
	CreatedAt     *time.Time            `json:"created_at,omitempty"`
	UpdatedAt     *time.Time            `json:"updated_at,omitempty"`
	CreatedBy     *NullableAuthorEntity `json:"created_by,omitempty"`
	TaskListItems []TaskListItemEntity  `json:"task_list_items,omitempty"`
}

func (t TaskListEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskListEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskListEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TaskListEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TaskListEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TaskListEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TaskListEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TaskListEntity) GetCreatedBy() *NullableAuthorEntity {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *TaskListEntity) GetTaskListItems() []TaskListItemEntity {
	if o == nil {
		return nil
	}
	return o.TaskListItems
}
