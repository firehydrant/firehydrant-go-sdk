# Audiences
(*Audiences*)

## Overview

Operations related to Audiences

### Available Operations

* [ListAudiences](#listaudiences) - List audiences
* [CreateAudience](#createaudience) - Create audience
* [GetAudience](#getaudience) - Get audience
* [ArchiveAudience](#archiveaudience) - Archive audience
* [UpdateAudience](#updateaudience) - Update audience
* [RestoreAudience](#restoreaudience) - Restore audience
* [GetMemberDefaultAudience](#getmemberdefaultaudience) - Get default audience
* [SetMemberDefaultAudience](#setmemberdefaultaudience) - Set default audience
* [GetAudienceSummary](#getaudiencesummary) - Get latest summary
* [GenerateAudienceSummary](#generateaudiencesummary) - Generate summary
* [ListAudienceSummaries](#listaudiencesummaries) - List audience summaries

## ListAudiences

List all audiences

### Example Usage

<!-- UsageSnippet language="go" operationID="list_audiences" method="get" path="/v1/audiences" -->
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

    res, err := s.Audiences.ListAudiences(ctx, firehydrantgosdk.Bool(false))
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
| `includeArchived`                                        | **bool*                                                  | :heavy_minus_sign:                                       | Include archived (discarded) audiences                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.AudiencesEntitiesAudienceEntity](../../models/components/audiencesentitiesaudienceentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAudience

Create a new audience

### Example Usage

<!-- UsageSnippet language="go" operationID="create_audience" method="post" path="/v1/audiences" -->
```go
package main

import(
	"context"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	"github.com/firehydrant/firehydrant-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Audiences.CreateAudience(ctx, operations.CreateAudienceRequest{
        Name: "<value>",
        Description: "simple contractor hmph along amongst thump provision crowded fragrant and",
        DetailsQuestion: []string{},
        DetailsPrompt: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
        },
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateAudienceRequest](../../models/operations/createaudiencerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*components.AudiencesEntitiesAudienceEntity](../../models/components/audiencesentitiesaudienceentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAudience

Get audience details

### Example Usage

<!-- UsageSnippet language="go" operationID="get_audience" method="get" path="/v1/audiences/{audience_id}" -->
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

    res, err := s.Audiences.GetAudience(ctx, "<id>")
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
| `audienceID`                                             | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the audience                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.AudiencesEntitiesAudienceEntity](../../models/components/audiencesentitiesaudienceentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ArchiveAudience

Archive an audience

### Example Usage

<!-- UsageSnippet language="go" operationID="archive_audience" method="delete" path="/v1/audiences/{audience_id}" -->
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

    res, err := s.Audiences.ArchiveAudience(ctx, "<id>")
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
| `audienceID`                                             | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the audience                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.AudiencesEntitiesAudienceEntity](../../models/components/audiencesentitiesaudienceentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAudience

Update an existing audience

### Example Usage

<!-- UsageSnippet language="go" operationID="update_audience" method="patch" path="/v1/audiences/{audience_id}" -->
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

    res, err := s.Audiences.UpdateAudience(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |
| `audienceID`                                                                                  | *string*                                                                                      | :heavy_check_mark:                                                                            | Unique identifier of the audience                                                             |
| `requestBody`                                                                                 | [*operations.UpdateAudienceRequestBody](../../models/operations/updateaudiencerequestbody.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |

### Response

**[*components.AudiencesEntitiesAudienceEntity](../../models/components/audiencesentitiesaudienceentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RestoreAudience

Restore a previously archived audience

### Example Usage

<!-- UsageSnippet language="go" operationID="restore_audience" method="patch" path="/v1/audiences/{audience_id}/restore" -->
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

    res, err := s.Audiences.RestoreAudience(ctx, "<id>")
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
| `audienceID`                                             | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the audience                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.AudiencesEntitiesAudienceEntity](../../models/components/audiencesentitiesaudienceentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMemberDefaultAudience

Get member's default audience

### Example Usage

<!-- UsageSnippet language="go" operationID="get_member_default_audience" method="get" path="/v1/audiences/member/{member_id}/default" -->
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

    res, err := s.Audiences.GetMemberDefaultAudience(ctx, 9319)
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
| `memberID`                                               | *int*                                                    | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.AudiencesEntitiesAudienceEntity](../../models/components/audiencesentitiesaudienceentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetMemberDefaultAudience

Set member's default audience

### Example Usage

<!-- UsageSnippet language="go" operationID="set_member_default_audience" method="put" path="/v1/audiences/member/{member_id}/default" -->
```go
package main

import(
	"context"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	"github.com/firehydrant/firehydrant-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Audiences.SetMemberDefaultAudience(ctx, 411452, operations.SetMemberDefaultAudienceRequestBody{
        AudienceID: "<id>",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `memberID`                                                                                                       | *int*                                                                                                            | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `requestBody`                                                                                                    | [operations.SetMemberDefaultAudienceRequestBody](../../models/operations/setmemberdefaultaudiencerequestbody.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*components.AudiencesEntitiesAudienceEntity](../../models/components/audiencesentitiesaudienceentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAudienceSummary

Get the latest audience-specific summary for an incident

### Example Usage

<!-- UsageSnippet language="go" operationID="get_audience_summary" method="get" path="/v1/audiences/{audience_id}/summaries/{incident_id}" -->
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

    res, err := s.Audiences.GetAudienceSummary(ctx, "<id>", "<id>")
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
| `audienceID`                                             | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the audience                        |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the incident to summarize           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.AIEntitiesIncidentSummaryEntity](../../models/components/aientitiesincidentsummaryentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GenerateAudienceSummary

Generate a new audience-specific summary for an incident

### Example Usage

<!-- UsageSnippet language="go" operationID="generate_audience_summary" method="post" path="/v1/audiences/{audience_id}/summaries/{incident_id}" -->
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

    res, err := s.Audiences.GenerateAudienceSummary(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                       | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                           | :heavy_check_mark:                                                                                              | The context to use for the request.                                                                             |
| `audienceID`                                                                                                    | *string*                                                                                                        | :heavy_check_mark:                                                                                              | Unique identifier of the audience                                                                               |
| `incidentID`                                                                                                    | *string*                                                                                                        | :heavy_check_mark:                                                                                              | Unique identifier of the incident to summarize                                                                  |
| `requestBody`                                                                                                   | [*operations.GenerateAudienceSummaryRequestBody](../../models/operations/generateaudiencesummaryrequestbody.md) | :heavy_minus_sign:                                                                                              | N/A                                                                                                             |
| `opts`                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                        | :heavy_minus_sign:                                                                                              | The options for this request.                                                                                   |

### Response

**[*components.AIEntitiesIncidentSummaryEntity](../../models/components/aientitiesincidentsummaryentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAudienceSummaries

List all audience summaries for an incident

### Example Usage

<!-- UsageSnippet language="go" operationID="list_audience_summaries" method="get" path="/v1/audiences/summaries/{incident_id}" -->
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

    res, err := s.Audiences.ListAudienceSummaries(ctx, "<id>")
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
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the incident to summarize           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.AudiencesEntitiesAudienceSummariesEntity](../../models/components/audiencesentitiesaudiencesummariesentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |