# PostV1TeamsTeamIDSignalRules

Create a Signals rule for a team.


## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `Name`                                                                                              | *string*                                                                                            | :heavy_check_mark:                                                                                  | The rule's name.                                                                                    |
| `Expression`                                                                                        | *string*                                                                                            | :heavy_check_mark:                                                                                  | The CEL expression that defines the rule.                                                           |
| `TargetType`                                                                                        | [components.TargetType](../../models/components/targettype.md)                                      | :heavy_check_mark:                                                                                  | The type of target that the rule will notify when matched.                                          |
| `TargetID`                                                                                          | *string*                                                                                            | :heavy_check_mark:                                                                                  | The ID of the target that the rule will notify when matched.                                        |
| `IncidentTypeID`                                                                                    | **string*                                                                                           | :heavy_minus_sign:                                                                                  | The ID of an incident type that should be used when an alert is promoted to an incident             |
| `NotificationPriorityOverride`                                                                      | [*components.NotificationPriorityOverride](../../models/components/notificationpriorityoverride.md) | :heavy_minus_sign:                                                                                  | A notification priority that will be set on the resulting alert (default: HIGH)                     |