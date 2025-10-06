# Pages
(*Pages*)

## Overview

Operations about Pages

### Available Operations

* [CreateSignalsPage](#createsignalspage) - Page a user, team, on-call schedule, or escalation policy

## CreateSignalsPage

Used for paging an on-call target within FireHydrant's signals product. This can be used for paging users, teams, on-call schedules, and escalation policies.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_signals_page" method="post" path="/v1/page/signals" -->
```go
package main

import(
	"context"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Pages.CreateSignalsPage(ctx, components.CreateSignalsPage{
        Summary: "<value>",
        TargetType: components.CreateSignalsPageTargetTypeTeam,
        TargetID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.CreateSignalsPage](../../models/components/createsignalspage.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*components.AlertsAlertEntity](../../models/components/alertsalertentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |