# CreateLifecycleMilestoneAutoAssignTimestampOnCreate

The setting for auto-assigning the milestone's timestamp during incident declaration

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/operations"
)

value := operations.CreateLifecycleMilestoneAutoAssignTimestampOnCreateAlwaysSetOnCreate
```


## Values

| Name                                                                       | Value                                                                      |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `CreateLifecycleMilestoneAutoAssignTimestampOnCreateAlwaysSetOnCreate`     | always_set_on_create                                                       |
| `CreateLifecycleMilestoneAutoAssignTimestampOnCreateOnlySetOnManualCreate` | only_set_on_manual_create                                                  |
| `CreateLifecycleMilestoneAutoAssignTimestampOnCreateNeverSetOnCreate`      | never_set_on_create                                                        |