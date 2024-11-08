# GetSignalsAnalyticsTimeseriesRequest


## Fields

| Field                                                                                                                                               | Type                                                                                                                                                | Required                                                                                                                                            | Description                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Bucket`                                                                                                                                            | [*operations.Bucket](../../models/operations/bucket.md)                                                                                             | :heavy_minus_sign:                                                                                                                                  | String that determines how records are grouped                                                                                                      |
| `SignalRules`                                                                                                                                       | **string*                                                                                                                                           | :heavy_minus_sign:                                                                                                                                  | A comma separated list of signal rule IDs                                                                                                           |
| `Teams`                                                                                                                                             | **string*                                                                                                                                           | :heavy_minus_sign:                                                                                                                                  | A comma separated list of team IDs                                                                                                                  |
| `Environments`                                                                                                                                      | **string*                                                                                                                                           | :heavy_minus_sign:                                                                                                                                  | A comma separated list of environment IDs                                                                                                           |
| `Services`                                                                                                                                          | **string*                                                                                                                                           | :heavy_minus_sign:                                                                                                                                  | A comma separated list of service IDs                                                                                                               |
| `Tags`                                                                                                                                              | **string*                                                                                                                                           | :heavy_minus_sign:                                                                                                                                  | A comma separated list of tags                                                                                                                      |
| `Users`                                                                                                                                             | **string*                                                                                                                                           | :heavy_minus_sign:                                                                                                                                  | A comma separated list of user IDs                                                                                                                  |
| `GroupBy`                                                                                                                                           | [*operations.GetSignalsAnalyticsTimeseriesQueryParamGroupBy](../../models/operations/getsignalsanalyticstimeseriesqueryparamgroupby.md)             | :heavy_minus_sign:                                                                                                                                  | String that determines how records are grouped                                                                                                      |
| `SortBy`                                                                                                                                            | [*operations.GetSignalsAnalyticsTimeseriesQueryParamSortBy](../../models/operations/getsignalsanalyticstimeseriesqueryparamsortby.md)               | :heavy_minus_sign:                                                                                                                                  | String that determines how records are sorted                                                                                                       |
| `SortDirection`                                                                                                                                     | [*operations.GetSignalsAnalyticsTimeseriesQueryParamSortDirection](../../models/operations/getsignalsanalyticstimeseriesqueryparamsortdirection.md) | :heavy_minus_sign:                                                                                                                                  | String that determines how records are sorted                                                                                                       |
| `StartDate`                                                                                                                                         | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                          | :heavy_minus_sign:                                                                                                                                  | The start date to return metrics from                                                                                                               |
| `EndDate`                                                                                                                                           | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                          | :heavy_minus_sign:                                                                                                                                  | The end date to return metrics from                                                                                                                 |