# CreateTeamEscalationPolicyDistributionType

The round robin configuration for the step. One of 'unspecified', 'round_robin_by_alert', or 'round_robin_by_escalation_policy'.

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

value := components.CreateTeamEscalationPolicyDistributionTypeUnspecified
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `CreateTeamEscalationPolicyDistributionTypeUnspecified`                  | unspecified                                                              |
| `CreateTeamEscalationPolicyDistributionTypeRoundRobinByAlert`            | round_robin_by_alert                                                     |
| `CreateTeamEscalationPolicyDistributionTypeRoundRobinByEscalationPolicy` | round_robin_by_escalation_policy                                         |