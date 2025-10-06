# PatchScimUser

PATCH SCIM endpoint to update a User. This endpoint is used to update a resource's attributes.


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `Trail`                                                                                  | **string*                                                                                | :heavy_minus_sign:                                                                       | An optional trail to log the request                                                     |
| `Operations`                                                                             | [][components.PatchScimUserOperation](../../models/components/patchscimuseroperation.md) | :heavy_check_mark:                                                                       | An array of operations to perform on the user                                            |