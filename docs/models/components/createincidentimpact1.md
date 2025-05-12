# CreateIncidentImpact1


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Type`                                                                              | *string*                                                                            | :heavy_check_mark:                                                                  | The type of impacted infrastructure. One of: environment, functionality, or service |
| `ID`                                                                                | *string*                                                                            | :heavy_check_mark:                                                                  | The ID of the impacted infrastructure                                               |
| `ConditionID`                                                                       | *string*                                                                            | :heavy_check_mark:                                                                  | The ID of the impact condition. Find these at /v1/severity_matrix/conditions        |