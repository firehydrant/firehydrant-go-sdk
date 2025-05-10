# ListRetrospectivesRequest


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Page`                                                       | **int*                                                       | :heavy_minus_sign:                                           | N/A                                                          |
| `PerPage`                                                    | **int*                                                       | :heavy_minus_sign:                                           | N/A                                                          |
| `IncidentID`                                                 | **string*                                                    | :heavy_minus_sign:                                           | Filter the reports by an incident ID                         |
| `UpdatedSince`                                               | [*time.Time](https://pkg.go.dev/time#Time)                   | :heavy_minus_sign:                                           | Filter for reports updated after the given ISO8601 timestamp |