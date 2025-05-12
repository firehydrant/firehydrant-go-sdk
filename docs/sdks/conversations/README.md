# Conversations
(*Conversations*)

## Overview

Operations related to Conversations

### Available Operations

* [GetVoteStatus](#getvotestatus) - Get votes
* [UpdateVote](#updatevote) - Update votes
* [DeleteCommentReaction](#deletecommentreaction) - Delete a reaction from a conversation comment
* [ListCommentReactions](#listcommentreactions) - List reactions for a conversation comment
* [CreateCommentReaction](#createcommentreaction) - Create a reaction for a conversation comment
* [GetComment](#getcomment) - Get a conversation comment
* [DeleteComment](#deletecomment) - Archive a conversation comment
* [UpdateComment](#updatecomment) - Update a conversation comment
* [ListComments](#listcomments) - List comments for a conversation
* [CreateComment](#createcomment) - Create a conversation comment

## GetVoteStatus

Get an object's current vote counts

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

    res, err := s.Conversations.GetVoteStatus(ctx, "<id>", "<id>")
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
| `eventID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.VotesEntity](../../models/components/votesentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateVote

Upvote or downvote an object

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

    res, err := s.Conversations.UpdateVote(ctx, "<id>", "<id>", components.UpdateVote{
        Direction: components.DirectionDig,
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

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `incidentID`                                                   | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `eventID`                                                      | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `updateVote`                                                   | [components.UpdateVote](../../models/components/updatevote.md) | :heavy_check_mark:                                             | N/A                                                            |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*components.VotesEntity](../../models/components/votesentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteCommentReaction

ALPHA - Archive a reaction

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

    err := s.Conversations.DeleteCommentReaction(ctx, "<id>", "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `reactionID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `commentID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommentReactions

ALPHA - List all of the reactions that have been added to a comment

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

    err := s.Conversations.ListCommentReactions(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `commentID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCommentReaction

ALPHA - Create a reaction on a comment

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

    err := s.Conversations.CreateCommentReaction(ctx, "<id>", "<id>", components.CreateCommentReaction{
        Reaction: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `conversationID`                                                                     | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `commentID`                                                                          | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `createCommentReaction`                                                              | [components.CreateCommentReaction](../../models/components/createcommentreaction.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetComment

ALPHA - Retrieves a single comment by ID

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

    err := s.Conversations.GetComment(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `commentID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteComment

ALPHA - Archive a comment

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

    err := s.Conversations.DeleteComment(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `commentID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateComment

ALPHA - Update a comment's attributes

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

    err := s.Conversations.UpdateComment(ctx, "<id>", "<id>", components.UpdateComment{})
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `commentID`                                                          | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `conversationID`                                                     | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `updateComment`                                                      | [components.UpdateComment](../../models/components/updatecomment.md) | :heavy_check_mark:                                                   | N/A                                                                  |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListComments

ALPHA - List all of the comments that have been added to the organization

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

    err := s.Conversations.ListComments(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `conversationID`                                                                         | *string*                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `before`                                                                                 | [*time.Time](https://pkg.go.dev/time#Time)                                               | :heavy_minus_sign:                                                                       | An ISO8601 timestamp that allows filtering for comments posted before the provided time. |
| `after`                                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                               | :heavy_minus_sign:                                                                       | An ISO8601 timestamp that allows filtering for comments posted after the provided time.  |
| `sort`                                                                                   | [*operations.ListCommentsSort](../../models/operations/listcommentssort.md)              | :heavy_minus_sign:                                                                       | Allows sorting comments by the time they were posted, ascending or descending.           |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateComment

ALPHA - Creates a comment for a project

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

    err := s.Conversations.CreateComment(ctx, "<id>", components.CreateComment{
        Body: "<value>",
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
| `conversationID`                                                     | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `createComment`                                                      | [components.CreateComment](../../models/components/createcomment.md) | :heavy_check_mark:                                                   | N/A                                                                  |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |