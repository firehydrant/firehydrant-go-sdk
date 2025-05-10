# ListProcessingLogEntriesRequest


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Page`                                                          | **int*                                                          | :heavy_minus_sign:                                              | N/A                                                             |
| `PerPage`                                                       | **int*                                                          | :heavy_minus_sign:                                              | N/A                                                             |
| `IntegrationSlug`                                               | **string*                                                       | :heavy_minus_sign:                                              | Scopes returned log entries to a specific integration ID        |
| `ConnectionID`                                                  | **string*                                                       | :heavy_minus_sign:                                              | Scopes returned log entries to a specific connection ID         |
| `OfLevel`                                                       | [*operations.OfLevel](../../models/operations/oflevel.md)       | :heavy_minus_sign:                                              | Returns logs of all levels equal to or above the provided level |
| `ExactLevel`                                                    | [*operations.ExactLevel](../../models/operations/exactlevel.md) | :heavy_minus_sign:                                              | Returns log entries of all levels equal to the provided level   |