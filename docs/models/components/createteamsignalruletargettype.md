# CreateTeamSignalRuleTargetType

The type of target that the rule will notify when matched.

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

value := components.CreateTeamSignalRuleTargetTypeEscalationPolicy
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `CreateTeamSignalRuleTargetTypeEscalationPolicy`      | EscalationPolicy                                      |
| `CreateTeamSignalRuleTargetTypeOnCallSchedule`        | OnCallSchedule                                        |
| `CreateTeamSignalRuleTargetTypeUser`                  | User                                                  |
| `CreateTeamSignalRuleTargetTypeWebhook`               | Webhook                                               |
| `CreateTeamSignalRuleTargetTypeSlackChannel`          | SlackChannel                                          |
| `CreateTeamSignalRuleTargetTypeMicrosoftTeamsChannel` | MicrosoftTeamsChannel                                 |