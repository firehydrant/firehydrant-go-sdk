# FireHydrantStatuspage
(*Integrations.Statuspage*)

## Overview

### Available Operations

* [ListConnections](#listconnections) - List Statuspage connections
* [ListPages](#listpages) - List StatusPage pages for a connection

## ListConnections

Lists the available and configured Statuspage integrations connections for the authenticated organization.

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Integrations.Statuspage.ListConnections(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationsStatuspageConnectionEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `page`                                                   | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `perPage`                                                | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListStatuspageConnectionsResponse](../../models/operations/liststatuspageconnectionsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequest              | 400, 413, 414, 415, 422, 431, 510 | application/json                  |
| sdkerrors.Unauthorized            | 401, 403, 407, 511                | application/json                  |
| sdkerrors.NotFound                | 404, 501, 505                     | application/json                  |
| sdkerrors.Timeout                 | 408, 504                          | application/json                  |
| sdkerrors.RateLimited             | 429                               | application/json                  |
| sdkerrors.InternalServerError     | 500, 502, 503, 506, 507, 508      | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## ListPages

List StatusPage pages for a connection

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Integrations.Statuspage.ListPages(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `connectionID`                                           | *string*                                                 | :heavy_check_mark:                                       | Connection UUID                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListStatuspagePagesResponse](../../models/operations/liststatuspagepagesresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequest              | 400, 413, 414, 415, 422, 431, 510 | application/json                  |
| sdkerrors.Unauthorized            | 401, 403, 407, 511                | application/json                  |
| sdkerrors.NotFound                | 404, 501, 505                     | application/json                  |
| sdkerrors.Timeout                 | 408, 504                          | application/json                  |
| sdkerrors.RateLimited             | 429                               | application/json                  |
| sdkerrors.InternalServerError     | 500, 502, 503, 506, 507, 508      | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |