package completion

import (
	"github.com/CodeNoobKing/go-ai-chat/server/context"
)

// CompletionID
// Each completion has an unique id.
type CompletionID string

// CompletionType
// The type of completion.
type CompletionType string

// ValueReader
// is the reader of the value of a given completion.
type ValueReader interface {
	// Read the value of a chunk of completion.
	// Arg ctx is the context of the prompt request.
	// Arg chunk contains a portion of the completion.
	// Arg completed indicates whether the completion is completed.
	// Arg err is the potential error of receiving completion value.
	Read(ctx *context.Context, chunk []byte, completed bool, err error) error
}

// Completion is the completion of a given prompt.
type Completion interface {
	// GetId returns the id of the completion.
	GetId(ctx *context.Context) CompletionID

	// GetType returns the type of the completion.
	GetType(ctx *context.Context) CompletionType

	// OnReceive register a value reader to read the value of the completion.
	OnReceive() <-chan []byte

	// Terminate terminates the completion.
	Terminate(ctx *context.Context) error
}

// ReadAllValue reads all the value of the completion.
func ReadAllValue(completion Completion) []byte {
	var buf []byte

	for chunk := range completion.OnReceive() {
		buf = append(buf, chunk...)
	}
	return buf
}
