package conversation

// Conversation is an interface to manipulate an existing session with ai.
type Conversation interface {
	// Ask Send a prompt to ai and get the completion.
	Ask(prompt Prompt) (Completion, error)

	// GetLastIndex Get the last index of current conversation.
	GetLastIndex() int

	// GetHistory Get the conversation's history.
	GetHistory(offset int) ([]Completion, error)

	// GetCompletion Get the completion by index.
	GetCompletion(index int) (Completion, error)

	// Clear the conversation's state.
	Clear() error
}
