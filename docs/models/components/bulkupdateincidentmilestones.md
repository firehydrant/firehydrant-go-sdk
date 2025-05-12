# BulkUpdateIncidentMilestones

Update milestone times in bulk for a given incident. All milestone
times for an incident must occur in chronological order
corresponding to the configured order of milestones. If the result
of this request would cause any milestone(s) to appear out of place,
a 422 response will instead be returned. This includes milestones
not explicitly submitted or updated in this request.



## Fields

| Field                                                                                                                  | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `Bulk`                                                                                                                 | [*components.Bulk](../../models/components/bulk.md)                                                                    | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `Milestones`                                                                                                           | [][components.BulkUpdateIncidentMilestonesMilestone](../../models/components/bulkupdateincidentmilestonesmilestone.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |