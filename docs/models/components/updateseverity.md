# UpdateSeverity

Update a specific severity


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Slug`                                                                              | **string*                                                                           | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `Description`                                                                       | **string*                                                                           | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `Position`                                                                          | **int*                                                                              | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `Color`                                                                             | [*components.UpdateSeverityColor](../../models/components/updateseveritycolor.md)   | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `AllowedRoleIds`                                                                    | []*string*                                                                          | :heavy_minus_sign:                                                                  | IDs of roles allowed to use this severity. Empty array means all roles are allowed. |