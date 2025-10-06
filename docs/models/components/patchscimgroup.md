# PatchScimGroup

SCIM endpoint to partially update a Team (Colloquial for Group in the SCIM protocol). Supports adding, removing, or replacing members using SCIM PATCH operations.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Trail`                                                                                    | **string*                                                                                  | :heavy_minus_sign:                                                                         | An optional trail to log the request                                                       |
| `Operations`                                                                               | [][components.PatchScimGroupOperation](../../models/components/patchscimgroupoperation.md) | :heavy_check_mark:                                                                         | An array of operations to perform on the group                                             |