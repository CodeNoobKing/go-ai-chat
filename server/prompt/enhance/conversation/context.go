package conversation

import "github.com/CodeNoobKing/go-ai-chat/server/context"

const (
	ContextKeyConversationId = "conversationID"
)

// GetConversationId returns the conversation id from context.
func GetConversationId(ctx *context.Context) ID {
	return ctx.Value(ContextKeyConversationId).(ID)
}
