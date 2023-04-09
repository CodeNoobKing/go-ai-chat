package conversation

import "github.com/CodeNoobKing/go-ai-chat/server/prompt"

const (
	// EnhancementTypeConversationHistory
	// Enhance the prompt with previous qa history
	EnhancementTypeConversationHistory prompt.EnhancementType = "ConversationHistory"

	// EnhancerTypeConversation
	// Enhance the prompt with conversation
	EnhancerTypeConversation prompt.EnhancerType = "Conversation"
)
