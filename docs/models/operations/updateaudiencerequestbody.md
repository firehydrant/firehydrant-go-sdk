# UpdateAudienceRequestBody


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Name`                                                  | **string*                                               | :heavy_minus_sign:                                      | Name of the audience (max 255 characters)               |
| `Description`                                           | **string*                                               | :heavy_minus_sign:                                      | Description of the audience (max 4000 characters)       |
| `Default`                                               | **bool*                                                 | :heavy_minus_sign:                                      | Whether this is the default audience                    |
| `Active`                                                | **bool*                                                 | :heavy_minus_sign:                                      | Whether the audience is active or discarded             |
| `DetailsQuestion`                                       | []*string*                                              | :heavy_minus_sign:                                      | The incident detail question (max 255 characters)       |
| `DetailsPrompt`                                         | []*string*                                              | :heavy_minus_sign:                                      | The prompt to display when collecting this detail       |
| `DetailsSlug`                                           | []*string*                                              | :heavy_minus_sign:                                      | Optional unique identifier for this detail              |
| `DetailsPosition`                                       | []*int*                                                 | :heavy_minus_sign:                                      | Position of the question in the list (1-based indexing) |