# CreateIncidentTask

Create a task


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Title`                                                              | *string*                                                             | :heavy_check_mark:                                                   | The title of the task.                                               |
| `State`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | The state of the task. One of: open, in_progress, cancelled, done    |
| `Description`                                                        | **string*                                                            | :heavy_minus_sign:                                                   | A description of what the task is for.                               |
| `AssigneeID`                                                         | **string*                                                            | :heavy_minus_sign:                                                   | The ID of the user assigned to the task.                             |
| `DueAt`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | The due date of the task. Relative values are supported such as '5m' |