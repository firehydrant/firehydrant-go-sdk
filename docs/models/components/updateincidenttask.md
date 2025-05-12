# UpdateIncidentTask

Update a task's attributes


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Title`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | The title of the task.                                               |
| `Description`                                                        | **string*                                                            | :heavy_minus_sign:                                                   | A description of what the task is for.                               |
| `State`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | The state of the task. One of: open, in_progress, cancelled, done    |
| `AssigneeID`                                                         | **string*                                                            | :heavy_minus_sign:                                                   | The ID of the user assigned to the task.                             |
| `DueAt`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | The due date of the task. Relative values are supported such as '5m' |