# UpdateTeamOnCallScheduleHandoffDay

The day of the week on which on-call shifts should hand off, as its long-form name (e.g. "monday", "tuesday", etc). This value is only used if the strategy type is "weekly".

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

value := components.UpdateTeamOnCallScheduleHandoffDayMonday
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `UpdateTeamOnCallScheduleHandoffDayMonday`    | monday                                        |
| `UpdateTeamOnCallScheduleHandoffDayTuesday`   | tuesday                                       |
| `UpdateTeamOnCallScheduleHandoffDayWednesday` | wednesday                                     |
| `UpdateTeamOnCallScheduleHandoffDayThursday`  | thursday                                      |
| `UpdateTeamOnCallScheduleHandoffDayFriday`    | friday                                        |
| `UpdateTeamOnCallScheduleHandoffDaySaturday`  | saturday                                      |
| `UpdateTeamOnCallScheduleHandoffDaySunday`    | sunday                                        |