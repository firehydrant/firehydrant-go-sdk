# CreateSignalsEmailTargetType

The type of target that the inbound email will notify when matched.

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

value := components.CreateSignalsEmailTargetTypeTeam
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `CreateSignalsEmailTargetTypeTeam`                  | Team                                                |
| `CreateSignalsEmailTargetTypeEntireTeam`            | EntireTeam                                          |
| `CreateSignalsEmailTargetTypeEscalationPolicy`      | EscalationPolicy                                    |
| `CreateSignalsEmailTargetTypeOnCallSchedule`        | OnCallSchedule                                      |
| `CreateSignalsEmailTargetTypeUser`                  | User                                                |
| `CreateSignalsEmailTargetTypeSlackChannel`          | SlackChannel                                        |
| `CreateSignalsEmailTargetTypeMicrosoftTeamsChannel` | MicrosoftTeamsChannel                               |
| `CreateSignalsEmailTargetTypeWebhook`               | Webhook                                             |