# Communication
(*Communication*)

## Overview

Operations related to Communication

### Available Operations

* [ListStatusUpdateTemplates](#liststatusupdatetemplates) - List status update templates
* [CreateStatusUpdateTemplate](#createstatusupdatetemplate) - Create a status update template
* [GetStatusUpdateTemplate](#getstatusupdatetemplate) - Get a status update template
* [DeleteStatusUpdateTemplate](#deletestatusupdatetemplate) - Delete a status update template
* [UpdateStatusUpdateTemplate](#updatestatusupdatetemplate) - Update a status update template

## ListStatusUpdateTemplates

List all status update templates for your organization

### Example Usage

<!-- UsageSnippet language="go" operationID="list_status_update_templates" method="get" path="/v1/status_update_templates" -->
```go
package main

import(
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

    res, err := s.Communication.ListStatusUpdateTemplates(ctx, nil, nil)
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
| `page`                                                   | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `perPage`                                                | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.StatusUpdateTemplateEntity](../../models/components/statusupdatetemplateentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateStatusUpdateTemplate

Create a status update template for your organization

### Example Usage

<!-- UsageSnippet language="go" operationID="create_status_update_template" method="post" path="/v1/status_update_templates" -->
```go
package main

import(
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

    res, err := s.Communication.CreateStatusUpdateTemplate(ctx, components.CreateStatusUpdateTemplate{
        Name: "<value>",
        Body: "<value>",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.CreateStatusUpdateTemplate](../../models/components/createstatusupdatetemplate.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.StatusUpdateTemplateEntity](../../models/components/statusupdatetemplateentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetStatusUpdateTemplate

Get a single status update template by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_status_update_template" method="get" path="/v1/status_update_templates/{status_update_template_id}" -->
```go
package main

import(
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

    res, err := s.Communication.GetStatusUpdateTemplate(ctx, "<id>")
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
| `statusUpdateTemplateID`                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.StatusUpdateTemplateEntity](../../models/components/statusupdatetemplateentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteStatusUpdateTemplate

Delete a single status update template

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_status_update_template" method="delete" path="/v1/status_update_templates/{status_update_template_id}" -->
```go
package main

import(
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

    res, err := s.Communication.DeleteStatusUpdateTemplate(ctx, "<id>")
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
| `statusUpdateTemplateID`                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.StatusUpdateTemplateEntity](../../models/components/statusupdatetemplateentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateStatusUpdateTemplate

Update a single status update template

### Example Usage

<!-- UsageSnippet language="go" operationID="update_status_update_template" method="patch" path="/v1/status_update_templates/{status_update_template_id}" -->
```go
package main

import(
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

    res, err := s.Communication.UpdateStatusUpdateTemplate(ctx, "<id>", components.UpdateStatusUpdateTemplate{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `statusUpdateTemplateID`                                                                       | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `updateStatusUpdateTemplate`                                                                   | [components.UpdateStatusUpdateTemplate](../../models/components/updatestatusupdatetemplate.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.StatusUpdateTemplateEntity](../../models/components/statusupdatetemplateentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |