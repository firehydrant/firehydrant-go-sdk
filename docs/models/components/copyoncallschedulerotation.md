# CopyOnCallScheduleRotation

Copy an on-call rotation into a different schedule, allowing you to merge them together safely.


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `TargetScheduleID`                                | *string*                                          | :heavy_check_mark:                                | The ID of the schedule to clone the rotation into |
| `Name`                                            | **string*                                         | :heavy_minus_sign:                                | The name of the on-call rotation                  |