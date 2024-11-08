# ListSavedSearchesRequest


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ResourceType`                                                     | [operations.ResourceType](../../models/operations/resourcetype.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `UserID`                                                           | **string*                                                          | :heavy_minus_sign:                                                 | The user ID used to filter saved searches.                         |
| `Query`                                                            | **string*                                                          | :heavy_minus_sign:                                                 | Filter saved searches with a query on their name                   |
| `Page`                                                             | **int*                                                             | :heavy_minus_sign:                                                 | N/A                                                                |
| `PerPage`                                                          | **int*                                                             | :heavy_minus_sign:                                                 | N/A                                                                |