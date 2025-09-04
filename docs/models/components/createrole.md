# CreateRole

Create a new role


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `Name`                                                                                  | *string*                                                                                | :heavy_check_mark:                                                                      | The name of the role.                                                                   |
| `Slug`                                                                                  | **string*                                                                               | :heavy_minus_sign:                                                                      | A unique identifier for the role. If not provided, one will be generated from the name. |
| `Description`                                                                           | **string*                                                                               | :heavy_minus_sign:                                                                      | A long-form description of the role's purpose.                                          |
| `Permissions`                                                                           | [][components.CreateRolePermission](../../models/components/createrolepermission.md)    | :heavy_minus_sign:                                                                      | An array of permission slugs to assign to the role.                                     |