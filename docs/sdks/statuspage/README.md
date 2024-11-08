# Statuspage
(*Statuspage*)

## Overview

### Available Operations

* [UpdateConnection](#updateconnection) - Update a Statuspage connection

## UpdateConnection

Update the given Statuspage integration connection.

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Statuspage.UpdateConnection(ctx, "<id>", components.PatchV1IntegrationsStatuspageConnectionsConnectionID{})
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationsStatuspageConnectionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `connectionID`                                                                                                                                     | *string*                                                                                                                                           | :heavy_check_mark:                                                                                                                                 | Connection UUID                                                                                                                                    |
| `patchV1IntegrationsStatuspageConnectionsConnectionID`                                                                                             | [components.PatchV1IntegrationsStatuspageConnectionsConnectionID](../../models/components/patchv1integrationsstatuspageconnectionsconnectionid.md) | :heavy_check_mark:                                                                                                                                 | N/A                                                                                                                                                |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.UpdateStatuspageConnectionResponse](../../models/operations/updatestatuspageconnectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |