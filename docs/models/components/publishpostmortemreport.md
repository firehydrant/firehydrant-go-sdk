# PublishPostMortemReport

Marks an incident retrospective as published and emails all of the participants in the report the summary


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Publish`                                                 | [*components.Publish](../../models/components/publish.md) | :heavy_minus_sign:                                        | N/A                                                       |
| `UserIds`                                                 | []*string*                                                | :heavy_minus_sign:                                        | An array of user IDs with whom to share the report        |
| `TeamIds`                                                 | []*string*                                                | :heavy_minus_sign:                                        | An array of team IDs with whom to share the report        |