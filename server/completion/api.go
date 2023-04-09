package completion

import (
	"context"

	"github.com/CodeNoobKing/go-ai-chat/server/prompt"
)

// API
// The API makes an abstraction of AI completion.
// Various implementations of AI completion can be plugged in.
type API interface {
	// Complete get the completion of the prompt.
	// The prompt is the user input.
	// The completion is the AI's response.
	Complete(ctx *context.Context, prompt prompt.Prompt) (completion Completion, err error)
}
