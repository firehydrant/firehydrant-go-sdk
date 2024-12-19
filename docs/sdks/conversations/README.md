# Conversations
(*Conversations*)

## Overview

Operations related to Conversations

### Available Operations

* [ListComments](#listcomments) - List comments for a conversation
* [CreateComment](#createcomment) - Create a comment for a conversation
* [GetComment](#getcomment) - Get a conversation comment
* [DeleteComment](#deletecomment) - Delete a conversation comment
* [UpdateComment](#updatecomment) - Update a conversation comment
* [ListCommentReactions](#listcommentreactions) - List reactions for a conversation comment
* [CreateCommentReaction](#createcommentreaction) - Create a reaction for a conversation comment
* [DeleteCommentReaction](#deletecommentreaction) - Delete a reaction from a conversation comment

## ListComments

ALPHA - List all of the comments that have been added to the organization

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

    res, err := s.Conversations.ListComments(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
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
| `sort`                                                                                   | [*operations.Sort](../../models/operations/sort.md)                                      | :heavy_minus_sign:                                                                       | Allows sorting comments by the time they were posted, ascending or descending.           |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListConversationCommentsResponse](../../models/operations/listconversationcommentsresponse.md), error**

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
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Conversations.CreateComment(ctx, "<id>", components.PostV1ConversationsConversationIDComments{
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `conversationID`                                                                                                             | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `postV1ConversationsConversationIDComments`                                                                                  | [components.PostV1ConversationsConversationIDComments](../../models/components/postv1conversationsconversationidcomments.md) | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.CreateConversationCommentResponse](../../models/operations/createconversationcommentresponse.md), error**

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

## GetComment

ALPHA - Retrieves a single comment by ID

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

    res, err := s.Conversations.GetComment(ctx, "<id>", "<id>")
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
| `commentID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetConversationCommentResponse](../../models/operations/getconversationcommentresponse.md), error**

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

## DeleteComment

ALPHA - Archive a comment

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

    res, err := s.Conversations.DeleteComment(ctx, "<id>", "<id>")
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
| `commentID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteConversationCommentResponse](../../models/operations/deleteconversationcommentresponse.md), error**

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
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Conversations.UpdateComment(ctx, "<id>", "<id>", components.PatchV1ConversationsConversationIDCommentsCommentID{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `commentID`                                                                                                                                      | *string*                                                                                                                                         | :heavy_check_mark:                                                                                                                               | N/A                                                                                                                                              |
| `conversationID`                                                                                                                                 | *string*                                                                                                                                         | :heavy_check_mark:                                                                                                                               | N/A                                                                                                                                              |
| `patchV1ConversationsConversationIDCommentsCommentID`                                                                                            | [components.PatchV1ConversationsConversationIDCommentsCommentID](../../models/components/patchv1conversationsconversationidcommentscommentid.md) | :heavy_check_mark:                                                                                                                               | N/A                                                                                                                                              |
| `opts`                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.UpdateConversationCommentResponse](../../models/operations/updateconversationcommentresponse.md), error**

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

## ListCommentReactions

ALPHA - List all of the reactions that have been added to a comment

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

    res, err := s.Conversations.ListCommentReactions(ctx, "<id>", "<id>")
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
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `commentID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListConversationCommentReactionsResponse](../../models/operations/listconversationcommentreactionsresponse.md), error**

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
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Conversations.CreateCommentReaction(ctx, "<id>", "<id>", components.PostV1ConversationsConversationIDCommentsCommentIDReactions{
        Reaction: "<value>",
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

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `conversationID`                                                                                                                                                 | *string*                                                                                                                                                         | :heavy_check_mark:                                                                                                                                               | N/A                                                                                                                                                              |
| `commentID`                                                                                                                                                      | *string*                                                                                                                                                         | :heavy_check_mark:                                                                                                                                               | N/A                                                                                                                                                              |
| `postV1ConversationsConversationIDCommentsCommentIDReactions`                                                                                                    | [components.PostV1ConversationsConversationIDCommentsCommentIDReactions](../../models/components/postv1conversationsconversationidcommentscommentidreactions.md) | :heavy_check_mark:                                                                                                                                               | N/A                                                                                                                                                              |
| `opts`                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.CreateConversationCommentReactionResponse](../../models/operations/createconversationcommentreactionresponse.md), error**

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

## DeleteCommentReaction

ALPHA - Archive a reaction

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

    res, err := s.Conversations.DeleteCommentReaction(ctx, "<id>", "<id>", "<id>")
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
| `reactionID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `commentID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteConversationCommentReactionResponse](../../models/operations/deleteconversationcommentreactionresponse.md), error**

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