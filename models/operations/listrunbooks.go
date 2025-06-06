// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
)

// ListRunbooksSort - Sort runbooks by their updated date. Accepts 'asc', 'desc'. This parameter is deprecated in favor of 'order_by' and 'order_direction'.
type ListRunbooksSort string

const (
	ListRunbooksSortAsc  ListRunbooksSort = "asc"
	ListRunbooksSortDesc ListRunbooksSort = "desc"
)

func (e ListRunbooksSort) ToPointer() *ListRunbooksSort {
	return &e
}
func (e *ListRunbooksSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = ListRunbooksSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListRunbooksSort: %v", v)
	}
}

// OrderBy - Sort runbooks by their updated date or name. Accepts 'updated_at', 'name', and 'created_at'.
type OrderBy string

const (
	OrderByUpdatedAt OrderBy = "updated_at"
	OrderByName      OrderBy = "name"
	OrderByCreatedAt OrderBy = "created_at"
)

func (e OrderBy) ToPointer() *OrderBy {
	return &e
}
func (e *OrderBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updated_at":
		fallthrough
	case "name":
		fallthrough
	case "created_at":
		*e = OrderBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrderBy: %v", v)
	}
}

// OrderDirection - Allows assigning a direction to how the specified `order_by` parameter is sorted. This parameter must be paired with `order_by` and does nothing on its own.
type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "asc"
	OrderDirectionDesc OrderDirection = "desc"
)

func (e OrderDirection) ToPointer() *OrderDirection {
	return &e
}
func (e *OrderDirection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = OrderDirection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrderDirection: %v", v)
	}
}

type ListRunbooksRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// A query to search runbooks by their name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// A query to search runbooks by their owners
	Owners *string `queryParam:"style=form,explode=true,name=owners"`
	// Sort runbooks by their updated date. Accepts 'asc', 'desc'. This parameter is deprecated in favor of 'order_by' and 'order_direction'.
	Sort *ListRunbooksSort `queryParam:"style=form,explode=true,name=sort"`
	// Sort runbooks by their updated date or name. Accepts 'updated_at', 'name', and 'created_at'.
	OrderBy *OrderBy `queryParam:"style=form,explode=true,name=order_by"`
	// Allows assigning a direction to how the specified `order_by` parameter is sorted. This parameter must be paired with `order_by` and does nothing on its own.
	OrderDirection *OrderDirection `queryParam:"style=form,explode=true,name=order_direction"`
}

func (o *ListRunbooksRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListRunbooksRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListRunbooksRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListRunbooksRequest) GetOwners() *string {
	if o == nil {
		return nil
	}
	return o.Owners
}

func (o *ListRunbooksRequest) GetSort() *ListRunbooksSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListRunbooksRequest) GetOrderBy() *OrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListRunbooksRequest) GetOrderDirection() *OrderDirection {
	if o == nil {
		return nil
	}
	return o.OrderDirection
}
