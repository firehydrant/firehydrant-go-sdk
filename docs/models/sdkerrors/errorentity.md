# ErrorEntity

ErrorEntity model


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Detail`                                          | `*string`                                         | :heavy_minus_sign:                                | N/A                                               |
| `Messages`                                        | []`string`                                        | :heavy_minus_sign:                                | N/A                                               |
| `Meta`                                            | [*sdkerrors.Meta](../../models/sdkerrors/meta.md) | :heavy_minus_sign:                                | An object with additional error metadata          |
| `Code`                                            | `*string`                                         | :heavy_minus_sign:                                | A stable code on which to match errors            |