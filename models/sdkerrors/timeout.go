// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// Timeout - Timeouts occurred with the request
type Timeout struct {
	Message              *string        `json:"message,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

var _ error = &Timeout{}

func (e *Timeout) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}