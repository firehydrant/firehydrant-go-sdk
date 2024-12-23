// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetScheduledMaintenanceRequest struct {
	ScheduledMaintenanceID string `pathParam:"style=simple,explode=false,name=scheduled_maintenance_id"`
}

func (o *GetScheduledMaintenanceRequest) GetScheduledMaintenanceID() string {
	if o == nil {
		return ""
	}
	return o.ScheduledMaintenanceID
}

type GetScheduledMaintenanceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Fetch the details of a scheduled maintenance event.
	ScheduledMaintenanceEntity *components.ScheduledMaintenanceEntity
}

func (o *GetScheduledMaintenanceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetScheduledMaintenanceResponse) GetScheduledMaintenanceEntity() *components.ScheduledMaintenanceEntity {
	if o == nil {
		return nil
	}
	return o.ScheduledMaintenanceEntity
}
