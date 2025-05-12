# Scim
(*Scim*)

## Overview

Operations related to SCIM

### Available Operations

* [GetScimGroup](#getscimgroup) - Get a SCIM group
* [UpdateScimGroup](#updatescimgroup) - Update a SCIM group and assign members
* [DeleteScimGroup](#deletescimgroup) - Delete a SCIM group
* [ListScimGroups](#listscimgroups) - List SCIM groups
* [CreateScimGroup](#createscimgroup) - Create a SCIM group and assign members
* [GetScimUser](#getscimuser) - Get a SCIM user
* [UpdateScimUser](#updatescimuser) - Update a User from SCIM data
* [DeleteScimUser](#deletescimuser) - Delete a User matching SCIM data
* [PatchScimUser](#patchscimuser) - Update a User from SCIM data
* [ListScimUsers](#listscimusers) - List SCIM users
* [CreateScimUser](#createscimuser) - Create a User from SCIM data

## GetScimGroup

SCIM endpoint that lists a Team (Colloquial for Group in the SCIM protocol)

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

    err := s.Scim.GetScimGroup(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateScimGroup

SCIM endpoint to update a Team (Colloquial for Group in the SCIM protocol). Any members defined in the payload will be assigned to the team with no defined role, any missing members will be removed from the team.

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

    err := s.Scim.UpdateScimGroup(ctx, "<id>", components.UpdateScimGroup{
        DisplayName: "Monte_Schowalter",
        Members: []components.UpdateScimGroupMember{
            components.UpdateScimGroupMember{
                Value: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `id`                                                                     | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `updateScimGroup`                                                        | [components.UpdateScimGroup](../../models/components/updatescimgroup.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteScimGroup

SCIM endpoint to delete a Team (Colloquial for Group in the SCIM protocol).

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

    err := s.Scim.DeleteScimGroup(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListScimGroups

SCIM endpoint that lists all Teams (Colloquial for Group in the SCIM protocol)

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

    err := s.Scim.ListScimGroups(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                          | Type                                                                                                                                                                                               | Required                                                                                                                                                                                           | Description                                                                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                |
| `startIndex`                                                                                                                                                                                       | **int*                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                 | N/A                                                                                                                                                                                                |
| `count`                                                                                                                                                                                            | **int*                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                 | N/A                                                                                                                                                                                                |
| `filter`                                                                                                                                                                                           | **string*                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                 | This is a string used to query groups by displayName.<br/>        Proper example syntax for this would be `?filter=displayName eq "My Team Name"`.<br/>        Currently we only support the `eq` operator |
| `opts`                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                 | The options for this request.                                                                                                                                                                      |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateScimGroup

SCIM endpoint to create a new Team (Colloquial for Group in the SCIM protocol). Any members defined in the payload will be assigned to the team with no defined role.

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

    err := s.Scim.CreateScimGroup(ctx, components.CreateScimGroup{
        DisplayName: "Allie.Schowalter",
        Members: []components.CreateScimGroupMember{
            components.CreateScimGroupMember{
                Value: "<value>",
            },
            components.CreateScimGroupMember{
                Value: "<value>",
            },
            components.CreateScimGroupMember{
                Value: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.CreateScimGroup](../../models/components/createscimgroup.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetScimUser

SCIM endpoint that lists a User

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

    err := s.Scim.GetScimUser(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateScimUser

PUT SCIM endpoint to update a User. This endpoint is used to replace a resource's attributes.

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

    err := s.Scim.UpdateScimUser(ctx, "<id>", components.UpdateScimUser{})
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |
| `updateScimUser`                                                       | [components.UpdateScimUser](../../models/components/updatescimuser.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteScimUser

SCIM endpoint to delete a User. This endpoint will deactivate a confirmed User record in our system.

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

    err := s.Scim.DeleteScimUser(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchScimUser

PATCH SCIM endpoint to update a User. This endpoint is used to update a resource's attributes.

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

    err := s.Scim.PatchScimUser(ctx, "<id>", components.PatchScimUser{
        Operations: []components.Operation{
            components.Operation{
                Op: "<value>",
                Path: "/System",
            },
            components.Operation{
                Op: "<value>",
                Path: "/lib",
            },
            components.Operation{
                Op: "<value>",
                Path: "/etc/namedb",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `id`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `patchScimUser`                                                      | [components.PatchScimUser](../../models/components/patchscimuser.md) | :heavy_check_mark:                                                   | N/A                                                                  |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListScimUsers

SCIM endpoint that lists users. This endpoint will display a list of Users currently in the system.

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

    err := s.Scim.ListScimUsers(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                | The context to use for the request.                                                                                                                                                                                                               |
| `filter`                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                | This is a string used to query users by either userName or email.<br/>        Proper example syntax for this would be `?filter=userName eq john` or `?filter=userName eq "john@firehydrant.com"`.<br/>        Currently we only support the `eq` operator |
| `startIndex`                                                                                                                                                                                                                                      | **int*                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                | This is an integer which represents a pagination offset                                                                                                                                                                                           |
| `count`                                                                                                                                                                                                                                           | **int*                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                | This is an integer which represents the number of items per page in the response                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                | The options for this request.                                                                                                                                                                                                                     |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateScimUser

SCIM endpoint to create and provision a new User. This endpoint will provision the User, which allows them to accept their account throught their IDP or via the Forgot Password flow.

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

    err := s.Scim.CreateScimUser(ctx, components.CreateScimUser{
        UserName: "Angeline72",
        Name: components.CreateScimUserName{
            FamilyName: "<value>",
            GivenName: "<value>",
        },
        Emails: []components.CreateScimUserEmail{
            components.CreateScimUserEmail{
                Value: "<value>",
                Primary: false,
            },
            components.CreateScimUserEmail{
                Value: "<value>",
                Primary: false,
            },
            components.CreateScimUserEmail{
                Value: "<value>",
                Primary: true,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.CreateScimUser](../../models/components/createscimuser.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |