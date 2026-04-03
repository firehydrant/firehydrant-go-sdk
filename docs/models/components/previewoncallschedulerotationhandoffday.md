# PreviewOnCallScheduleRotationHandoffDay

The day of the week on which on-call shifts should hand off, as its long-form name (e.g. "monday", "tuesday", etc). This value is only used if the strategy type is "weekly".

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

value := components.PreviewOnCallScheduleRotationHandoffDayMonday
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `PreviewOnCallScheduleRotationHandoffDayMonday`    | monday                                             |
| `PreviewOnCallScheduleRotationHandoffDayTuesday`   | tuesday                                            |
| `PreviewOnCallScheduleRotationHandoffDayWednesday` | wednesday                                          |
| `PreviewOnCallScheduleRotationHandoffDayThursday`  | thursday                                           |
| `PreviewOnCallScheduleRotationHandoffDayFriday`    | friday                                             |
| `PreviewOnCallScheduleRotationHandoffDaySaturday`  | saturday                                           |
| `PreviewOnCallScheduleRotationHandoffDaySunday`    | sunday                                             |