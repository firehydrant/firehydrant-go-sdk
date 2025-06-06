// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type UpdateRunbookExecutionStepScriptRequest struct {
	ExecutionID string `pathParam:"style=simple,explode=false,name=execution_id"`
	StepID      string `pathParam:"style=simple,explode=false,name=step_id"`
	State       string `pathParam:"style=simple,explode=false,name=state"`
}

func (o *UpdateRunbookExecutionStepScriptRequest) GetExecutionID() string {
	if o == nil {
		return ""
	}
	return o.ExecutionID
}

func (o *UpdateRunbookExecutionStepScriptRequest) GetStepID() string {
	if o == nil {
		return ""
	}
	return o.StepID
}

func (o *UpdateRunbookExecutionStepScriptRequest) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}
