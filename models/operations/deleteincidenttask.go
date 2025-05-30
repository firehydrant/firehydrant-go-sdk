// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DeleteIncidentTaskRequest struct {
	TaskID     string `pathParam:"style=simple,explode=false,name=task_id"`
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *DeleteIncidentTaskRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

func (o *DeleteIncidentTaskRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}
