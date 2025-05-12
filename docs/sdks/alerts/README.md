# Alerts
(*Alerts*)

## Overview

Operations related to Alerts

### Available Operations

* [ListIncidentAlerts](#listincidentalerts) - List alerts for an incident
* [CreateIncidentAlert](#createincidentalert) - Attach an alert to an incident
* [UpdateIncidentAlertPrimary](#updateincidentalertprimary) - Set an alert as primary for an incident
* [DeleteIncidentAlert](#deleteincidentalert) - Remove an alert from an incident
* [ListAlerts](#listalerts) - List alerts
* [GetAlert](#getalert) - Get an alert
* [ListProcessingLogEntries](#listprocessinglogentries) - List alert processing log entries

## ListIncidentAlerts

List alerts that have been attached to an incident

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

    res, err := s.Alerts.ListIncidentAlerts(ctx, "<id>")
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
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.IncidentsAlertEntityPaginated](../../models/components/incidentsalertentitypaginated.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateIncidentAlert

Add an alert to an incident. FireHydrant needs to have ingested the alert from a third party system in order to attach it to the incident.

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

    err := s.Alerts.CreateIncidentAlert(ctx, "<id>", []string{
        "<value>",
        "<value>",
        "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `requestBody`                                            | []*string*                                               | :heavy_check_mark:                                       | Array of alert IDs to be assigned to the incident        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateIncidentAlertPrimary

Setting an alert as primary will overwrite milestone times in the FireHydrant incident with times included in the primary alert. Services attached to the primary alert will also be automatically added to the incident.

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

    res, err := s.Alerts.UpdateIncidentAlertPrimary(ctx, "<id>", "<id>", components.UpdateIncidentAlertPrimary{
        Primary: false,
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
| `incidentAlertID`                                                                              | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `incidentID`                                                                                   | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `updateIncidentAlertPrimary`                                                                   | [components.UpdateIncidentAlertPrimary](../../models/components/updateincidentalertprimary.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.IncidentsAlertEntity](../../models/components/incidentsalertentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteIncidentAlert

Remove an alert from an incident

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

    err := s.Alerts.DeleteIncidentAlert(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentAlertID`                                        | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAlerts

Retrieve all alerts from third parties

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

    res, err := s.Alerts.ListAlerts(ctx, operations.ListAlertsRequest{})
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
| `request`                                                                    | [operations.ListAlertsRequest](../../models/operations/listalertsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*components.AlertsAlertEntityPaginated](../../models/components/alertsalertentitypaginated.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAlert

Retrieve a single alert

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

    res, err := s.Alerts.GetAlert(ctx, "<id>")
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
| `alertID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.AlertsAlertEntity](../../models/components/alertsalertentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListProcessingLogEntries

Processing Log Entries for a specific alert

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

    res, err := s.Alerts.ListProcessingLogEntries(ctx, operations.ListProcessingLogEntriesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListProcessingLogEntriesRequest](../../models/operations/listprocessinglogentriesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*components.AlertsProcessingLogEntryEntityPaginated](../../models/components/alertsprocessinglogentryentitypaginated.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |