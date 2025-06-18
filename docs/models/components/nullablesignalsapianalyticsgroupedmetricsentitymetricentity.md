# NullableSignalsAPIAnalyticsGroupedMetricsEntityMetricEntity


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `GroupedID`                                  | **int*                                       | :heavy_minus_sign:                           | The ID of the group                          |
| `GroupedName`                                | **string*                                    | :heavy_minus_sign:                           | The name of the group                        |
| `TotalOpenedAlerts`                          | **int*                                       | :heavy_minus_sign:                           | The total number of opened alerts            |
| `TotalAckedAlerts`                           | **int*                                       | :heavy_minus_sign:                           | The total number of acknowledged alerts      |
| `TotalIncidents`                             | **int*                                       | :heavy_minus_sign:                           | The total number of incidents                |
| `AckedPercentage`                            | **float32*                                   | :heavy_minus_sign:                           | The percentage of acknowledged alerts        |
| `IncidentsPercentage`                        | **float32*                                   | :heavy_minus_sign:                           | The percentage of alerts that have incidents |