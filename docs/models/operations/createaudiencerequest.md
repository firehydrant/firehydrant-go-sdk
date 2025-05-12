# CreateAudienceRequest


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Name`                                            | *string*                                          | :heavy_check_mark:                                | Name of the audience (max 255 characters)         |
| `Description`                                     | *string*                                          | :heavy_check_mark:                                | Description of the audience (max 4000 characters) |
| `Default`                                         | **bool*                                           | :heavy_minus_sign:                                | Whether this is the default audience              |
| `DetailsQuestion`                                 | []*string*                                        | :heavy_check_mark:                                | The incident detail question (max 255 characters) |
| `DetailsPrompt`                                   | []*string*                                        | :heavy_check_mark:                                | The prompt to display when collecting this detail |
| `DetailsSlug`                                     | []*string*                                        | :heavy_minus_sign:                                | Optional unique identifier for this detail        |