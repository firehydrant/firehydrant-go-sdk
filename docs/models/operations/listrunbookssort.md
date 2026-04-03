# ListRunbooksSort

Sort runbooks by their updated date. Accepts 'asc', 'desc'. This parameter is deprecated in favor of 'order_by' and 'order_direction'.

## Example Usage

```go
import (
	"github.com/firehydrant/firehydrant-go-sdk/models/operations"
)

value := operations.ListRunbooksSortAsc
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ListRunbooksSortAsc`  | asc                    |
| `ListRunbooksSortDesc` | desc                   |