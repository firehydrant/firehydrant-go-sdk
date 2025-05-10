# UpdateScimGroup

SCIM endpoint to update a Team (Colloquial for Group in the SCIM protocol). Any members defined in the payload will be assigned to the team with no defined role, any missing members will be removed from the team.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `DisplayName`                                                                          | *string*                                                                               | :heavy_check_mark:                                                                     | The name of the team being updated                                                     |
| `Members`                                                                              | [][components.UpdateScimGroupMember](../../models/components/updatescimgroupmember.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |