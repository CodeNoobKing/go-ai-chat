package conversation

// Factory is the interface that helps to generate a conversation.
type Factory interface {
	// NewConversation Generate a new conversation.
	NewConversation(ctx Context) Conversation
}
