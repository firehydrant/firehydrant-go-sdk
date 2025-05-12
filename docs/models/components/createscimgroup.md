# CreateScimGroup

SCIM endpoint to create a new Team (Colloquial for Group in the SCIM protocol). Any members defined in the payload will be assigned to the team with no defined role.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `DisplayName`                                                                          | *string*                                                                               | :heavy_check_mark:                                                                     | The name of the team being created                                                     |
| `Members`                                                                              | [][components.CreateScimGroupMember](../../models/components/createscimgroupmember.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |