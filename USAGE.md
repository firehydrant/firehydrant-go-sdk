<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrant.New(
		firehydrant.WithSecurity(components.Security{
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