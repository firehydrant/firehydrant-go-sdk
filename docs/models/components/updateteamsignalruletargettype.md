# UpdateTeamSignalRuleTargetType

The type of target that the rule will notify when matched.

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

value := components.UpdateTeamSignalRuleTargetTypeEscalationPolicy
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `UpdateTeamSignalRuleTargetTypeEscalationPolicy`      | EscalationPolicy                                      |
| `UpdateTeamSignalRuleTargetTypeOnCallSchedule`        | OnCallSchedule                                        |
| `UpdateTeamSignalRuleTargetTypeUser`                  | User                                                  |
| `UpdateTeamSignalRuleTargetTypeWebhook`               | Webhook                                               |
| `UpdateTeamSignalRuleTargetTypeSlackChannel`          | SlackChannel                                          |
| `UpdateTeamSignalRuleTargetTypeMicrosoftTeamsChannel` | MicrosoftTeamsChannel                                 |