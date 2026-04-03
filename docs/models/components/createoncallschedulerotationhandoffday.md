# CreateOnCallScheduleRotationHandoffDay

The day of the week on which on-call shifts should hand off, as its long-form name (e.g. "monday", "tuesday", etc). This value is only used if the strategy type is "weekly".

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

value := components.CreateOnCallScheduleRotationHandoffDayMonday
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `CreateOnCallScheduleRotationHandoffDayMonday`    | monday                                            |
| `CreateOnCallScheduleRotationHandoffDayTuesday`   | tuesday                                           |
| `CreateOnCallScheduleRotationHandoffDayWednesday` | wednesday                                         |
| `CreateOnCallScheduleRotationHandoffDayThursday`  | thursday                                          |
| `CreateOnCallScheduleRotationHandoffDayFriday`    | friday                                            |
| `CreateOnCallScheduleRotationHandoffDaySaturday`  | saturday                                          |
| `CreateOnCallScheduleRotationHandoffDaySunday`    | sunday                                            |