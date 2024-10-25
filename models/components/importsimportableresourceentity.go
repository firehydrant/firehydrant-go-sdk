// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"firehydrant/internal/utils"
	"fmt"
	"time"
)

type ImportsImportableResourceEntityState string

const (
	ImportsImportableResourceEntityStateSelected ImportsImportableResourceEntityState = "selected"
	ImportsImportableResourceEntityStateSkipped  ImportsImportableResourceEntityState = "skipped"
	ImportsImportableResourceEntityStateImported ImportsImportableResourceEntityState = "imported"
	ImportsImportableResourceEntityStateErrored  ImportsImportableResourceEntityState = "errored"
)

func (e ImportsImportableResourceEntityState) ToPointer() *ImportsImportableResourceEntityState {
	return &e
}
func (e *ImportsImportableResourceEntityState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "selected":
		fallthrough
	case "skipped":
		fallthrough
	case "imported":
		fallthrough
	case "errored":
		*e = ImportsImportableResourceEntityState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ImportsImportableResourceEntityState: %v", v)
	}
}

type ImportsImportableResourceEntity struct {
	ImportErrors []ImportsImportErrorEntity            `json:"import_errors,omitempty"`
	ImportedAt   *time.Time                            `json:"imported_at,omitempty"`
	RemoteID     *string                               `json:"remote_id,omitempty"`
	State        *ImportsImportableResourceEntityState `json:"state,omitempty"`
}

func (i ImportsImportableResourceEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *ImportsImportableResourceEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ImportsImportableResourceEntity) GetImportErrors() []ImportsImportErrorEntity {
	if o == nil {
		return nil
	}
	return o.ImportErrors
}

func (o *ImportsImportableResourceEntity) GetImportedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ImportedAt
}

func (o *ImportsImportableResourceEntity) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *ImportsImportableResourceEntity) GetState() *ImportsImportableResourceEntityState {
	if o == nil {
		return nil
	}
	return o.State
}