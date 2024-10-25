# PostV1ChecklistTemplates

Creates a checklist template for the organization


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Name`                                                                         | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |
| `Checks`                                                                       | [][components.Checks](../../models/components/checks.md)                       | :heavy_check_mark:                                                             | An array of checks for the checklist template                                  |
| `Description`                                                                  | **string*                                                                      | :heavy_minus_sign:                                                             | N/A                                                                            |
| `TeamID`                                                                       | **string*                                                                      | :heavy_minus_sign:                                                             | The ID of the Team that owns the checklist template                            |
| `ConnectedServices`                                                            | [][components.ConnectedServices](../../models/components/connectedservices.md) | :heavy_minus_sign:                                                             | Array of service IDs to attach checklist template to                           |