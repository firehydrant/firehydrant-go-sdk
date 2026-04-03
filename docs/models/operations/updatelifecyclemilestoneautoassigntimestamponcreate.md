# UpdateLifecycleMilestoneAutoAssignTimestampOnCreate

The setting for auto-assigning the milestone's timestamp during incident declaration

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/operations"
)

value := operations.UpdateLifecycleMilestoneAutoAssignTimestampOnCreateAlwaysSetOnCreate
```


## Values

| Name                                                                       | Value                                                                      |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `UpdateLifecycleMilestoneAutoAssignTimestampOnCreateAlwaysSetOnCreate`     | always_set_on_create                                                       |
| `UpdateLifecycleMilestoneAutoAssignTimestampOnCreateOnlySetOnManualCreate` | only_set_on_manual_create                                                  |
| `UpdateLifecycleMilestoneAutoAssignTimestampOnCreateNeverSetOnCreate`      | never_set_on_create                                                        |