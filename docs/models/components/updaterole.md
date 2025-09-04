# UpdateRole

Update a role


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Name`                                                                               | *string*                                                                             | :heavy_check_mark:                                                                   | The name of the role.                                                                |
| `Description`                                                                        | **string*                                                                            | :heavy_minus_sign:                                                                   | A long-form description of the role's purpose.                                       |
| `Permissions`                                                                        | [][components.UpdateRolePermission](../../models/components/updaterolepermission.md) | :heavy_minus_sign:                                                                   | An array of permission slugs to assign to the role.                                  |