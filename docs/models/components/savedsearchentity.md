# SavedSearchEntity

SavedSearchEntity model


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ID`                                                         | **string*                                                    | :heavy_minus_sign:                                           | N/A                                                          |
| `Name`                                                       | **string*                                                    | :heavy_minus_sign:                                           | N/A                                                          |
| `ResourceType`                                               | **string*                                                    | :heavy_minus_sign:                                           | N/A                                                          |
| `UserID`                                                     | **string*                                                    | :heavy_minus_sign:                                           | N/A                                                          |
| `IsPrivate`                                                  | **bool*                                                      | :heavy_minus_sign:                                           | Whether or not this saved search is private                  |
| `CreatedAt`                                                  | [*time.Time](https://pkg.go.dev/time#Time)                   | :heavy_minus_sign:                                           | N/A                                                          |
| `UpdatedAt`                                                  | [*time.Time](https://pkg.go.dev/time#Time)                   | :heavy_minus_sign:                                           | N/A                                                          |
| `FilterValues`                                               | map[string]*any*                                             | :heavy_minus_sign:                                           | An unstructured key/value pair of saved values for searching |