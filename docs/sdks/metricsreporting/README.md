# MetricsReporting
(*MetricsReporting*)

## Overview

Operations related to Metrics & Reporting

### Available Operations

* [GetMeanTimeReport](#getmeantimereport) - Get mean time metrics for incidents
* [ListTicketFunnelMetrics](#listticketfunnelmetrics) - List ticket task and follow up creation and completion metrics
* [ListRetrospectiveMetrics](#listretrospectivemetrics) - List retrospective metrics
* [ListMilestoneFunnelMetrics](#listmilestonefunnelmetrics) - List milestone funnel metrics
* [ListUserInvolvementMetrics](#listuserinvolvementmetrics) - List user metrics
* [ListIncidentMetrics](#listincidentmetrics) - List incident metrics and analytics
* [ListMttxMetrics](#listmttxmetrics) - Get infrastructure metrics
* [ListInfrastructureTypeMetrics](#listinfrastructuretypemetrics) - List metrics for a component type
* [ListInfrastructureMetrics](#listinfrastructuremetrics) - Get metrics for a component
* [GetSavedSearch](#getsavedsearch) - Get a saved search
* [DeleteSavedSearch](#deletesavedsearch) - Delete a saved search
* [UpdateSavedSearch](#updatesavedsearch) - Update a saved search
* [ListSavedSearches](#listsavedsearches) - List saved searches
* [CreateSavedSearch](#createsavedsearch) - Create a saved search
* [GetSignalsTimeseriesAnalytics](#getsignalstimeseriesanalytics) - Generate timeseries alert metrics
* [GetSignalsGroupedMetrics](#getsignalsgroupedmetrics) - Generate grouped alert metrics
* [GetSignalsMttxAnalytics](#getsignalsmttxanalytics) - Get MTTX analytics for signals
* [ExportSignalsShiftAnalytics](#exportsignalsshiftanalytics) - Export on-call hours report

## GetMeanTimeReport

Returns a report with time bucketed analytics data

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.GetMeanTimeReport(ctx, operations.GetMeanTimeReportRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetMeanTimeReportRequest](../../models/operations/getmeantimereportrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.ReportEntity](../../models/components/reportentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTicketFunnelMetrics

Returns a report with task and follow up creation and completion data

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.ListTicketFunnelMetrics(ctx, operations.ListTicketFunnelMetricsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListTicketFunnelMetricsRequest](../../models/operations/listticketfunnelmetricsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*components.MetricsTicketFunnelMetricsEntity](../../models/components/metricsticketfunnelmetricsentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListRetrospectiveMetrics

Returns a report with retrospective analytics data

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
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.ListRetrospectiveMetrics(ctx, nil, nil)
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
| `startDate`                                              | [*types.Date](../../types/date.md)                       | :heavy_minus_sign:                                       | The start date to return metrics from                    |
| `endDate`                                                | [*types.Date](../../types/date.md)                       | :heavy_minus_sign:                                       | The end date to return metrics from                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.MetricsRetrospectiveEntity](../../models/components/metricsretrospectiveentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMilestoneFunnelMetrics

Returns a report with time bucketed milestone data

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.ListMilestoneFunnelMetrics(ctx, operations.ListMilestoneFunnelMetricsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListMilestoneFunnelMetricsRequest](../../models/operations/listmilestonefunnelmetricsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*components.MetricsMilestonesFunnelEntity](../../models/components/metricsmilestonesfunnelentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUserInvolvementMetrics

Returns a report with time bucketed analytics data

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.ListUserInvolvementMetrics(ctx, operations.ListUserInvolvementMetricsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListUserInvolvementMetricsRequest](../../models/operations/listuserinvolvementmetricsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*components.MetricsMetricsEntity](../../models/components/metricsmetricsentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListIncidentMetrics

Returns a report with time bucketed analytics data

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.ListIncidentMetrics(ctx, operations.ListIncidentMetricsRequest{})
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
| `request`                                                                                      | [operations.ListIncidentMetricsRequest](../../models/operations/listincidentmetricsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.MetricsMetricsEntity](../../models/components/metricsmetricsentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMttxMetrics

Fetch infrastructure metrics based on custom query

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/types"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.ListMttxMetrics(ctx, operations.ListMttxMetricsRequest{
        StartDate: types.MustDateFromString("2025-05-13"),
        EndDate: types.MustDateFromString("2024-08-28"),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListMttxMetricsRequest](../../models/operations/listmttxmetricsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*components.MetricsMttxDataEntity](../../models/components/metricsmttxdataentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListInfrastructureTypeMetrics

Returns metrics for all components of a given type

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.ListInfrastructureTypeMetrics(ctx, operations.ListInfrastructureTypeMetricsInfraTypeCustomers, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `infraType`                                                                                                            | [operations.ListInfrastructureTypeMetricsInfraType](../../models/operations/listinfrastructuretypemetricsinfratype.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `startDate`                                                                                                            | [*types.Date](../../types/date.md)                                                                                     | :heavy_minus_sign:                                                                                                     | The start date to return metrics from; defaults to 30 days ago                                                         |
| `endDate`                                                                                                              | [*types.Date](../../types/date.md)                                                                                     | :heavy_minus_sign:                                                                                                     | The end date to return metrics from, defaults to today                                                                 |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*components.MetricsInfrastructureListEntity](../../models/components/metricsinfrastructurelistentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListInfrastructureMetrics

Return metrics for a specific component

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.ListInfrastructureMetrics(ctx, operations.ListInfrastructureMetricsInfraTypeEnvironments, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `infraType`                                                                                                    | [operations.ListInfrastructureMetricsInfraType](../../models/operations/listinfrastructuremetricsinfratype.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `infraID`                                                                                                      | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Component UUID                                                                                                 |
| `startDate`                                                                                                    | [*types.Date](../../types/date.md)                                                                             | :heavy_minus_sign:                                                                                             | The start date to return metrics from; defaults to 30 days ago                                                 |
| `endDate`                                                                                                      | [*types.Date](../../types/date.md)                                                                             | :heavy_minus_sign:                                                                                             | The end date to return metrics from, defaults to today                                                         |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*components.MetricsInfrastructureMetricsEntity](../../models/components/metricsinfrastructuremetricsentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSavedSearch

Retrieve a specific save search

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.GetSavedSearch(ctx, operations.GetSavedSearchResourceTypeIncidents, "<id>")
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
| `resourceType`                                                                                 | [operations.GetSavedSearchResourceType](../../models/operations/getsavedsearchresourcetype.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `savedSearchID`                                                                                | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.SavedSearchEntity](../../models/components/savedsearchentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteSavedSearch

Delete a specific saved search

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.DeleteSavedSearch(ctx, operations.DeleteSavedSearchResourceTypeScheduledMaintenances, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `resourceType`                                                                                       | [operations.DeleteSavedSearchResourceType](../../models/operations/deletesavedsearchresourcetype.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `savedSearchID`                                                                                      | *string*                                                                                             | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*components.SavedSearchEntity](../../models/components/savedsearchentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateSavedSearch

Update a specific saved search

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.UpdateSavedSearch(ctx, operations.UpdateSavedSearchResourceTypeIncidentEvents, "<id>", components.UpdateSavedSearch{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `resourceType`                                                                                       | [operations.UpdateSavedSearchResourceType](../../models/operations/updatesavedsearchresourcetype.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `savedSearchID`                                                                                      | *string*                                                                                             | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `updateSavedSearch`                                                                                  | [components.UpdateSavedSearch](../../models/components/updatesavedsearch.md)                         | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*components.SavedSearchEntity](../../models/components/savedsearchentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSavedSearches

Lists saved searches

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.ListSavedSearches(ctx, operations.ListSavedSearchesRequest{
        ResourceType: operations.ListSavedSearchesResourceTypeIncidents,
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListSavedSearchesRequest](../../models/operations/listsavedsearchesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.SavedSearchEntity](../../models/components/savedsearchentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateSavedSearch

Create a new saved search for a particular resource type

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.MetricsReporting.CreateSavedSearch(ctx, operations.CreateSavedSearchResourceTypeServices, components.CreateSavedSearch{
        Name: "<value>",
        FilterValues: map[string]any{
            "key": "<value>",
            "key1": "<value>",
            "key2": "<value>",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `resourceType`                                                                                       | [operations.CreateSavedSearchResourceType](../../models/operations/createsavedsearchresourcetype.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `createSavedSearch`                                                                                  | [components.CreateSavedSearch](../../models/components/createsavedsearch.md)                         | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*components.SavedSearchEntity](../../models/components/savedsearchentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsTimeseriesAnalytics

Generate a timeseries-based report of metrics for Signals alerts

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    err := s.MetricsReporting.GetSignalsTimeseriesAnalytics(ctx, operations.GetSignalsTimeseriesAnalyticsRequest{})
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetSignalsTimeseriesAnalyticsRequest](../../models/operations/getsignalstimeseriesanalyticsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsGroupedMetrics

Generate a report of grouped metrics for Signals alerts

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    err := s.MetricsReporting.GetSignalsGroupedMetrics(ctx, operations.GetSignalsGroupedMetricsRequest{})
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetSignalsGroupedMetricsRequest](../../models/operations/getsignalsgroupedmetricsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsMttxAnalytics

Get mean-time-to-acknowledged (MTTA) and mean-time-to-resolved (MTTR) metrics for Signals alerts

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    err := s.MetricsReporting.GetSignalsMttxAnalytics(ctx, operations.GetSignalsMttxAnalyticsRequest{})
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetSignalsMttxAnalyticsRequest](../../models/operations/getsignalsmttxanalyticsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ExportSignalsShiftAnalytics

Export on-call hours report for users/teams during a time period

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrant.New(
        firehydrant.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    err := s.MetricsReporting.ExportSignalsShiftAnalytics(ctx, types.MustTimeFromString("2025-01-23T11:55:06.817Z"), types.MustTimeFromString("2023-09-24T03:00:50.158Z"), nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                               | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                   | :heavy_check_mark:                                                                                                      | The context to use for the request.                                                                                     |
| `periodStart`                                                                                                           | [time.Time](https://pkg.go.dev/time#Time)                                                                               | :heavy_check_mark:                                                                                                      | Start of the period to fetch hours for (UTC)                                                                            |
| `periodEnd`                                                                                                             | [time.Time](https://pkg.go.dev/time#Time)                                                                               | :heavy_check_mark:                                                                                                      | End of the period to fetch hours for (UTC)                                                                              |
| `requestBody`                                                                                                           | [*operations.ExportSignalsShiftAnalyticsRequestBody](../../models/operations/exportsignalsshiftanalyticsrequestbody.md) | :heavy_minus_sign:                                                                                                      | N/A                                                                                                                     |
| `opts`                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                | :heavy_minus_sign:                                                                                                      | The options for this request.                                                                                           |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |