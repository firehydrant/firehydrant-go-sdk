# IngestCatalogData

Accepts catalog data in the configured format and asyncronously processes the data to incorporate changes into service catalog.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Encoding`                                                 | [components.Encoding](../../models/components/encoding.md) | :heavy_check_mark:                                         | Encoding for submitted data                                |
| `Data`                                                     | *string*                                                   | :heavy_check_mark:                                         | Service data                                               |