# CreateOnCallShift

Create a Signals on-call shift in a schedule.


## Fields

| Field                                                                                           | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `StartTime`                                                                                     | *string*                                                                                        | :heavy_check_mark:                                                                              | The start time of the shift in ISO8601 format.                                                  |
| `EndTime`                                                                                       | *string*                                                                                        | :heavy_check_mark:                                                                              | The end time of the shift in ISO8601 format.                                                    |
| `UserID`                                                                                        | **string*                                                                                       | :heavy_minus_sign:                                                                              | The ID of the user who is on-call for the shift. If not provided, the shift will be unassigned. |