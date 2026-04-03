# UpdateSignalsEmailTargetType

The type of target that the inbound email will notify when matched.

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

value := components.UpdateSignalsEmailTargetTypeTeam
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `UpdateSignalsEmailTargetTypeTeam`                  | Team                                                |
| `UpdateSignalsEmailTargetTypeEntireTeam`            | EntireTeam                                          |
| `UpdateSignalsEmailTargetTypeEscalationPolicy`      | EscalationPolicy                                    |
| `UpdateSignalsEmailTargetTypeOnCallSchedule`        | OnCallSchedule                                      |
| `UpdateSignalsEmailTargetTypeUser`                  | User                                                |
| `UpdateSignalsEmailTargetTypeSlackChannel`          | SlackChannel                                        |
| `UpdateSignalsEmailTargetTypeMicrosoftTeamsChannel` | MicrosoftTeamsChannel                               |
| `UpdateSignalsEmailTargetTypeWebhook`               | Webhook                                             |