# Communication
(*Communication*)

## Overview

Operations related to Communication

### Available Operations

* [ListTemplates](#listtemplates) - List status update templates
* [GetTemplate](#gettemplate) - Get a status update template
* [DeleteStatusUpdateTemplate](#deletestatusupdatetemplate) - Delete a status update template
* [UpdateTemplate](#updatetemplate) - Update a status update template

## ListTemplates

List all status update templates for your organization

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

    res, err := s.Communication.ListTemplates(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.StatusUpdateTemplateEntity != nil {
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

**[*operations.ListStatusUpdateTemplatesResponse](../../models/operations/liststatusupdatetemplatesresponse.md), error**

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

## GetTemplate

Get a single status update template by ID

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

    res, err := s.Communication.GetTemplate(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.StatusUpdateTemplateEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `statusUpdateTemplateID`                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetStatusUpdateTemplateResponse](../../models/operations/getstatusupdatetemplateresponse.md), error**

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

## DeleteStatusUpdateTemplate

Delete a single status update template

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

    res, err := s.Communication.DeleteStatusUpdateTemplate(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.StatusUpdateTemplateEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `statusUpdateTemplateID`                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteStatusUpdateTemplateResponse](../../models/operations/deletestatusupdatetemplateresponse.md), error**

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

## UpdateTemplate

Update a single status update template

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

    res, err := s.Communication.UpdateTemplate(ctx, "<id>", components.PatchV1StatusUpdateTemplatesStatusUpdateTemplateID{})
    if err != nil {
        log.Fatal(err)
    }
    if res.StatusUpdateTemplateEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `statusUpdateTemplateID`                                                                                                                       | *string*                                                                                                                                       | :heavy_check_mark:                                                                                                                             | N/A                                                                                                                                            |
| `patchV1StatusUpdateTemplatesStatusUpdateTemplateID`                                                                                           | [components.PatchV1StatusUpdateTemplatesStatusUpdateTemplateID](../../models/components/patchv1statusupdatetemplatesstatusupdatetemplateid.md) | :heavy_check_mark:                                                                                                                             | N/A                                                                                                                                            |
| `opts`                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.UpdateStatusUpdateTemplateResponse](../../models/operations/updatestatusupdatetemplateresponse.md), error**

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