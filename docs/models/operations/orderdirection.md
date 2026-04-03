# OrderDirection

Allows assigning a direction to how the specified `order_by` parameter is sorted. This parameter must be paired with `order_by` and does nothing on its own.

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/operations"
)

value := operations.OrderDirectionAsc
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `OrderDirectionAsc`  | asc                  |
| `OrderDirectionDesc` | desc                 |