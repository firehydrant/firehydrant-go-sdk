# ExportSignalsShiftAnalyticsRequest


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `UserIds`                                    | []*string*                                   | :heavy_minus_sign:                           | Array of user IDs to fetch oncall hours for  |
| `TeamIds`                                    | []*string*                                   | :heavy_minus_sign:                           | Array of team IDs to fetch oncall hours for  |
| `PeriodStart`                                | [time.Time](https://pkg.go.dev/time#Time)    | :heavy_check_mark:                           | Start of the period to fetch hours for (UTC) |
| `PeriodEnd`                                  | [time.Time](https://pkg.go.dev/time#Time)    | :heavy_check_mark:                           | End of the period to fetch hours for (UTC)   |