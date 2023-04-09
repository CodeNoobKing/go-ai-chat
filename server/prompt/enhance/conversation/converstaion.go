package conversation

import (
	"math"
	"sync"

	"github.com/CodeNoobKing/go-ai-chat/server/common"
	"github.com/CodeNoobKing/go-ai-chat/server/completion"
	"github.com/CodeNoobKing/go-ai-chat/server/context"
	"github.com/CodeNoobKing/go-ai-chat/server/prompt"
	"github.com/CodeNoobKing/go-ai-chat/server/utils/errors"
)

// ID the id of conversation.
type ID string

// QA a question and answer pair.
type QA struct {
	// Question
	// The byte value of user prompt content.
	Question []byte `json:"question"`

	// Answer
	// The byte value of AI completion.
	Answer []byte `json:"answer"`
}

// NewConversation creates a new conversation.
func NewConversation(ctx *context.Context) *Conversation {
	return &Conversation{
		Lock:            &sync.RWMutex{},
		ID:              GetConversationId(ctx),
		Persist:         false,
		MaxQALimitation: 10,
		History:         []QA{},
		OnGoingPrompt:   nil,
	}
}

// Conversation is an interface to manipulate an existing session with AI.
// The Conversation would maintain a context with AI which can help user achieve better completion.
type Conversation struct {
	// Lock  used for concurrent access.
	// readonly
	Lock *sync.RWMutex

	// ID the id of conversation.
	// readonly
	ID ID

	// Persist indicates whether the conversation should be persisted.
	// If not, the content would be lost after the server restart.
	Persist bool

	// MaxQALimitation
	// The max limitation of the history of conversation required to be considered.
	MaxQALimitation int

	// History the history of conversation.
	History []QA

	// OnGoingPrompt the current prompt.
	OnGoingPrompt *prompt.Prompt
}

func (conv *Conversation) getEnhancements(_ *context.Context, _ prompt.Prompt) prompt.Enhancement {
	conv.Lock.Lock()
	defer conv.Lock.Unlock()

	limit := int(math.Min(float64(conv.MaxQALimitation), float64(len(conv.History))))

	return prompt.Enhancement{
		Type:         EnhancementTypeConversationHistory,
		EnhancerType: EnhancerTypeConversation,
		Description:  "",
		Content:      conv.History[len(conv.History)-limit:],
	}
}

// Enhance the prompt with conversation history.
func (conv *Conversation) Enhance(ctx *context.Context, p prompt.Prompt) error {
	conv.Lock.Lock()
	defer conv.Lock.Unlock()
	if conv.OnGoingPrompt != nil {
		return errors.New(common.ErrCodePromptIsGenerating, "current prompt is generating")
	}

	conv.OnGoingPrompt = &p

	ctx.Lock.Lock()
	defer ctx.Lock.Unlock()

	ctx.Enhancements = append(ctx.Enhancements, conv.getEnhancements(ctx, p))

	go func() {
		question := prompt.ReadAllContent(p.Contents[0])
		values := completion.ReadAllValue(<-ctx.OnCompletionBegin)

		ctx.Lock.Lock()
		defer ctx.Lock.Unlock()

		conv.History = append(conv.History, QA{
			Question: question,
			Answer:   values,
		})

	}()

	return nil
}
