# ChangeEntity

ChangeEntity model


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ID`                                                                            | `*string`                                                                       | :heavy_minus_sign:                                                              | UUID of the Change                                                              |
| `Summary`                                                                       | `*string`                                                                       | :heavy_minus_sign:                                                              | Description of the Change                                                       |
| `CreatedAt`                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                      | :heavy_minus_sign:                                                              | The time the change entry was created                                           |
| `UpdatedAt`                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                      | :heavy_minus_sign:                                                              | N/A                                                                             |
| `Labels`                                                                        | [*components.ChangeEntityLabels](../../models/components/changeentitylabels.md) | :heavy_minus_sign:                                                              | Arbitrary key/value pairs of labels.                                            |
| `Description`                                                                   | `*string`                                                                       | :heavy_minus_sign:                                                              | Description of the Change                                                       |