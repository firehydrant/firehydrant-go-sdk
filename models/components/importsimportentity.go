// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
	"time"
)

type ImportsImportEntityState string

const (
	ImportsImportEntityStatePreprocessing  ImportsImportEntityState = "preprocessing"
	ImportsImportEntityStateReadyForImport ImportsImportEntityState = "ready_for_import"
	ImportsImportEntityStateImporting      ImportsImportEntityState = "importing"
	ImportsImportEntityStateCompleted      ImportsImportEntityState = "completed"
	ImportsImportEntityStateFailed         ImportsImportEntityState = "failed"
)

func (e ImportsImportEntityState) ToPointer() *ImportsImportEntityState {
	return &e
}
func (e *ImportsImportEntityState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "preprocessing":
		fallthrough
	case "ready_for_import":
		fallthrough
	case "importing":
		fallthrough
	case "completed":
		fallthrough
	case "failed":
		*e = ImportsImportEntityState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ImportsImportEntityState: %v", v)
	}
}

// ImportsImportEntity - Imports_ImportEntity model
type ImportsImportEntity struct {
	State     *ImportsImportEntityState `json:"state,omitempty"`
	UpdatedAt *time.Time                `json:"updated_at,omitempty"`
}

func (i ImportsImportEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *ImportsImportEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ImportsImportEntity) GetState() *ImportsImportEntityState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *ImportsImportEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
