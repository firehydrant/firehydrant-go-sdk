<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrantgosdk.New(
		firehydrantgosdk.WithSecurity(components.Security{
			APIKey: "<YOUR_API_KEY_HERE>",
		}),
	)

	res, err := s.AccountSettings.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->