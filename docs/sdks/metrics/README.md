# Metrics
(*Metrics*)

## Overview

### Available Operations

* [GetMttx](#getmttx) - Fetch infrastructure metrics based on custom query
* [GetInfrastructure](#getinfrastructure) - Get metrics for a specific catalog entry

## GetMttx

Fetch infrastructure metrics based on custom query

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/types"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Metrics.GetMttx(ctx, operations.GetV1MetricsMttxRequest{
        StartDate: types.MustDateFromString("2024-12-27"),
        EndDate: types.MustDateFromString("2024-02-03"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsMttxDataEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetV1MetricsMttxRequest](../../models/operations/getv1metricsmttxrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetV1MetricsMttxResponse](../../models/operations/getv1metricsmttxresponse.md), error**

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

## GetInfrastructure

Return metrics for a specific component

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

    res, err := s.Metrics.GetInfrastructure(ctx, operations.PathParamInfraTypeFunctionalities, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsInfrastructureMetricsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `infraType`                                                                    | [operations.PathParamInfraType](../../models/operations/pathparaminfratype.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `infraID`                                                                      | *string*                                                                       | :heavy_check_mark:                                                             | Component UUID                                                                 |
| `startDate`                                                                    | [*types.Date](../../types/date.md)                                             | :heavy_minus_sign:                                                             | The start date to return metrics from; defaults to 30 days ago                 |
| `endDate`                                                                      | [*types.Date](../../types/date.md)                                             | :heavy_minus_sign:                                                             | The end date to return metrics from, defaults to today                         |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetInfrastructureMetricsResponse](../../models/operations/getinfrastructuremetricsresponse.md), error**

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