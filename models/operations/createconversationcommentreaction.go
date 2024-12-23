// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateConversationCommentReactionRequest struct {
	ConversationID                                              string                                                                 `pathParam:"style=simple,explode=false,name=conversation_id"`
	CommentID                                                   string                                                                 `pathParam:"style=simple,explode=false,name=comment_id"`
	PostV1ConversationsConversationIDCommentsCommentIDReactions components.PostV1ConversationsConversationIDCommentsCommentIDReactions `request:"mediaType=application/json"`
}

func (o *CreateConversationCommentReactionRequest) GetConversationID() string {
	if o == nil {
		return ""
	}
	return o.ConversationID
}

func (o *CreateConversationCommentReactionRequest) GetCommentID() string {
	if o == nil {
		return ""
	}
	return o.CommentID
}

func (o *CreateConversationCommentReactionRequest) GetPostV1ConversationsConversationIDCommentsCommentIDReactions() components.PostV1ConversationsConversationIDCommentsCommentIDReactions {
	if o == nil {
		return components.PostV1ConversationsConversationIDCommentsCommentIDReactions{}
	}
	return o.PostV1ConversationsConversationIDCommentsCommentIDReactions
}

type CreateConversationCommentReactionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *CreateConversationCommentReactionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
