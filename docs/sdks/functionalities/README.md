# Functionalities
(*Functionalities*)

## Overview

Operations related to Functionalities

### Available Operations

* [List](#list) - List functionalities
* [Create](#create) - Create a functionality
* [Get](#get) - Get a functionality
* [Delete](#delete) - Archive a functionality
* [Update](#update) - Update a functionality
* [ListServices](#listservices) - List services for a functionality

## List

List all of the functionalities that have been added to the organiation

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Functionalities.List(ctx, operations.ListFunctionalitiesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionalityEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListFunctionalitiesRequest](../../models/operations/listfunctionalitiesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListFunctionalitiesResponse](../../models/operations/listfunctionalitiesresponse.md), error**

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

## Create

Creates a functionality for the organization

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Functionalities.Create(ctx, components.PostV1Functionalities{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionalityEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.PostV1Functionalities](../../models/components/postv1functionalities.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateFunctionalityResponse](../../models/operations/createfunctionalityresponse.md), error**

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

## Get

Retrieves a single functionality by ID

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

    res, err := s.Functionalities.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionalityEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `functionalityID`                                        | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetFunctionalityResponse](../../models/operations/getfunctionalityresponse.md), error**

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

## Delete

Archive a functionality

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

    res, err := s.Functionalities.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionalityEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `functionalityID`                                        | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteFunctionalityResponse](../../models/operations/deletefunctionalityresponse.md), error**

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

## Update

Update a functionalities attributes

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Functionalities.Update(ctx, "<id>", components.PatchV1FunctionalitiesFunctionalityID{})
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionalityEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `functionalityID`                                                                                                    | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `patchV1FunctionalitiesFunctionalityID`                                                                              | [components.PatchV1FunctionalitiesFunctionalityID](../../models/components/patchv1functionalitiesfunctionalityid.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.UpdateFunctionalityResponse](../../models/operations/updatefunctionalityresponse.md), error**

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

## ListServices

List services for a functionality

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

    res, err := s.Functionalities.ListServices(ctx, "<id>")
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
| `functionalityID`                                        | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetFunctionalityServicesResponse](../../models/operations/getfunctionalityservicesresponse.md), error**

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