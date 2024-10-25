# TicketingTicketEntity

Ticketing_TicketEntity model


## Fields

| Field                                                                                                                      | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                                       | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `Summary`                                                                                                                  | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `Description`                                                                                                              | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `State`                                                                                                                    | [*components.State](../../models/components/state.md)                                                                      | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `Type`                                                                                                                     | [*components.TicketingTicketEntityType](../../models/components/ticketingticketentitytype.md)                              | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `Assignees`                                                                                                                | [][components.AuthorEntity](../../models/components/authorentity.md)                                                       | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `Priority`                                                                                                                 | [*components.TicketingPriorityEntity](../../models/components/ticketingpriorityentity.md)                                  | :heavy_minus_sign:                                                                                                         | Ticketing_PriorityEntity model                                                                                             |
| `CreatedBy`                                                                                                                | [*components.AuthorEntity](../../models/components/authorentity.md)                                                        | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `Attachments`                                                                                                              | [][components.Attachments](../../models/components/attachments.md)                                                         | :heavy_minus_sign:                                                                                                         | A list of objects attached to this item. Can be one of: LinkEntity, CustomerSupportIssueEntity, or GenericAttachmentEntity |
| `CreatedAt`                                                                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                                                                 | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `UpdatedAt`                                                                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                                                                 | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `TagList`                                                                                                                  | []*string*                                                                                                                 | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `IncidentID`                                                                                                               | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | ID of incident that this ticket is related to                                                                              |
| `IncidentName`                                                                                                             | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Name of incident that this ticket is related to                                                                            |
| `IncidentCurrentMilestone`                                                                                                 | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | Milestone of incident that this ticket is related to                                                                       |
| `TaskID`                                                                                                                   | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | ID of task that this ticket is related to                                                                                  |
| `DueAt`                                                                                                                    | [*time.Time](https://pkg.go.dev/time#Time)                                                                                 | :heavy_minus_sign:                                                                                                         | N/A                                                                                                                        |
| `Link`                                                                                                                     | [*components.AttachmentsLinkEntity](../../models/components/attachmentslinkentity.md)                                      | :heavy_minus_sign:                                                                                                         | Attachments_LinkEntity model                                                                                               |