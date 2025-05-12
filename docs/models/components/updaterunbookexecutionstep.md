# UpdateRunbookExecutionStep

Updates a runbook step execution, especially for changing the state of a step execution.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `State`                                    | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `ScheduleFor`                              | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `Data`                                     | map[string]*any*                           | :heavy_minus_sign:                         | Data for execution of this step            |
| `RepeatsAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |