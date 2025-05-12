# CreateIncidentChatMessage

Create a new generic chat message on an incident timeline. These are independent of any specific chat provider.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Body`                                                                | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `OccurredAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                            | :heavy_minus_sign:                                                    | ISO8601 timestamp for when the chat message occurred                  |
| `VoteDirection`                                                       | [*components.VoteDirection](../../models/components/votedirection.md) | :heavy_minus_sign:                                                    | N/A                                                                   |