# CreateAudience

Create a new audience


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Name`                                                      | *string*                                                    | :heavy_check_mark:                                          | Name of the audience (max 255 characters)                   |
| `Description`                                               | *string*                                                    | :heavy_check_mark:                                          | Description of the audience (max 4000 characters)           |
| `Default`                                                   | **bool*                                                     | :heavy_minus_sign:                                          | Whether this is the default audience                        |
| `Details`                                                   | [][components.Detail](../../models/components/detail.md)    | :heavy_minus_sign:                                          | N/A                                                         |
| `Settings`                                                  | [*components.Settings](../../models/components/settings.md) | :heavy_minus_sign:                                          | audience settings for initial audience creation             |