package conversation

import (
	"sync"

	"github.com/CodeNoobKing/go-ai-chat/server/context"
	"github.com/CodeNoobKing/go-ai-chat/server/prompt"
)

var _ prompt.Enhancer = (*Enhancer)(nil)

type Enhancer struct {
	lock *sync.RWMutex

	// conversation map
	idToConversation map[ID]*Conversation
}

func (e *Enhancer) Enhance(ctx *context.Context, p prompt.Prompt) error {
	conversationId := GetConversationId(ctx)

	e.lock.Lock()
	defer e.lock.Unlock()

	conversation, ok := e.idToConversation[conversationId]
	if !ok {
		conversation = NewConversation(ctx)
		e.idToConversation[conversationId] = conversation
	}

	return conversation.Enhance(ctx, p)
}
