# AlertsProcessingLogEntryEntity


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ID`                                                              | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `Context`                                                         | [*components.Context](../../models/components/context.md)         | :heavy_minus_sign:                                                | An unstructured representation of this log entry's context.       |
| `CreatedAt`                                                       | [*time.Time](https://pkg.go.dev/time#Time)                        | :heavy_minus_sign:                                                | N/A                                                               |
| `Level`                                                           | [*components.Level](../../models/components/level.md)             | :heavy_minus_sign:                                                | N/A                                                               |
| `Message`                                                         | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `MessageType`                                                     | [*components.MessageType](../../models/components/messagetype.md) | :heavy_minus_sign:                                                | N/A                                                               |