// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateSignalsAlertGroupingConfigurationRequest struct {
	ID                                      string                                             `pathParam:"style=simple,explode=false,name=id"`
	UpdateSignalsAlertGroupingConfiguration components.UpdateSignalsAlertGroupingConfiguration `request:"mediaType=application/json"`
}

func (o *UpdateSignalsAlertGroupingConfigurationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSignalsAlertGroupingConfigurationRequest) GetUpdateSignalsAlertGroupingConfiguration() components.UpdateSignalsAlertGroupingConfiguration {
	if o == nil {
		return components.UpdateSignalsAlertGroupingConfiguration{}
	}
	return o.UpdateSignalsAlertGroupingConfiguration
}
