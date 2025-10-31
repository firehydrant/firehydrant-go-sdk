# CreateSeverity

Create a new severity


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Slug`                                                                              | *string*                                                                            | :heavy_check_mark:                                                                  | N/A                                                                                 |
| `Description`                                                                       | **string*                                                                           | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `Position`                                                                          | **int*                                                                              | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `Color`                                                                             | [*components.CreateSeverityColor](../../models/components/createseveritycolor.md)   | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `AllowedRoleIds`                                                                    | []*string*                                                                          | :heavy_minus_sign:                                                                  | IDs of roles allowed to use this severity. Empty array means all roles are allowed. |