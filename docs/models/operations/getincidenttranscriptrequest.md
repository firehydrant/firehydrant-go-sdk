# GetIncidentTranscriptRequest


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `After`                                                                 | **string*                                                               | :heavy_minus_sign:                                                      | The ID of the transcript entry to start after.                          |
| `Before`                                                                | **string*                                                               | :heavy_minus_sign:                                                      | The ID of the transcript entry to start before.                         |
| `Sort`                                                                  | [*operations.QueryParamSort](../../models/operations/queryparamsort.md) | :heavy_minus_sign:                                                      | The order to sort the transcript entries.                               |
| `IncidentID`                                                            | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |