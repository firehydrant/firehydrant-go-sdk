# CreateRunbookExecution

Attaches a runbook to an incident and executes it


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ExecuteFor`                                                                    | *string*                                                                        | :heavy_check_mark:                                                              | The incident to attach the runbook to. Format must be: `incident/${incidentId}` |
| `RunbookID`                                                                     | *string*                                                                        | :heavy_check_mark:                                                              | ID of runbook to attach                                                         |