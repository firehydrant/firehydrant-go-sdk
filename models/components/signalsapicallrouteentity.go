// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type SignalsAPICallRouteEntityRoutingMode string

const (
	SignalsAPICallRouteEntityRoutingModeRoutingModeTakeMessage   SignalsAPICallRouteEntityRoutingMode = "ROUTING_MODE_TAKE_MESSAGE"
	SignalsAPICallRouteEntityRoutingModeRoutingModeDirectConnect SignalsAPICallRouteEntityRoutingMode = "ROUTING_MODE_DIRECT_CONNECT"
)

func (e SignalsAPICallRouteEntityRoutingMode) ToPointer() *SignalsAPICallRouteEntityRoutingMode {
	return &e
}
func (e *SignalsAPICallRouteEntityRoutingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ROUTING_MODE_TAKE_MESSAGE":
		fallthrough
	case "ROUTING_MODE_DIRECT_CONNECT":
		*e = SignalsAPICallRouteEntityRoutingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SignalsAPICallRouteEntityRoutingMode: %v", v)
	}
}

type SignalsAPICallRouteEntityConnectMode string

const (
	SignalsAPICallRouteEntityConnectModeConnectModeConference SignalsAPICallRouteEntityConnectMode = "CONNECT_MODE_CONFERENCE"
	SignalsAPICallRouteEntityConnectModeConnectModeDirectDial SignalsAPICallRouteEntityConnectMode = "CONNECT_MODE_DIRECT_DIAL"
)

func (e SignalsAPICallRouteEntityConnectMode) ToPointer() *SignalsAPICallRouteEntityConnectMode {
	return &e
}
func (e *SignalsAPICallRouteEntityConnectMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CONNECT_MODE_CONFERENCE":
		fallthrough
	case "CONNECT_MODE_DIRECT_DIAL":
		*e = SignalsAPICallRouteEntityConnectMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SignalsAPICallRouteEntityConnectMode: %v", v)
	}
}

// SignalsAPICallRouteEntity - Signals_API_CallRouteEntity model
type SignalsAPICallRouteEntity struct {
	ID              *string                                `json:"id,omitempty"`
	Name            *string                                `json:"name,omitempty"`
	Description     *string                                `json:"description,omitempty"`
	PhoneNumber     *string                                `json:"phone_number,omitempty"`
	GreetingMessage *string                                `json:"greeting_message,omitempty"`
	RoutingMode     *SignalsAPICallRouteEntityRoutingMode  `json:"routing_mode,omitempty"`
	ConnectMode     *SignalsAPICallRouteEntityConnectMode  `json:"connect_mode,omitempty"`
	Steps           *NullableSignalsAPICallRouteStepEntity `json:"steps,omitempty"`
	Target          *NullableSignalsAPITargetEntity        `json:"target,omitempty"`
}

func (o *SignalsAPICallRouteEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SignalsAPICallRouteEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SignalsAPICallRouteEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SignalsAPICallRouteEntity) GetPhoneNumber() *string {
	if o == nil {
		return nil
	}
	return o.PhoneNumber
}

func (o *SignalsAPICallRouteEntity) GetGreetingMessage() *string {
	if o == nil {
		return nil
	}
	return o.GreetingMessage
}

func (o *SignalsAPICallRouteEntity) GetRoutingMode() *SignalsAPICallRouteEntityRoutingMode {
	if o == nil {
		return nil
	}
	return o.RoutingMode
}

func (o *SignalsAPICallRouteEntity) GetConnectMode() *SignalsAPICallRouteEntityConnectMode {
	if o == nil {
		return nil
	}
	return o.ConnectMode
}

func (o *SignalsAPICallRouteEntity) GetSteps() *NullableSignalsAPICallRouteStepEntity {
	if o == nil {
		return nil
	}
	return o.Steps
}

func (o *SignalsAPICallRouteEntity) GetTarget() *NullableSignalsAPITargetEntity {
	if o == nil {
		return nil
	}
	return o.Target
}
