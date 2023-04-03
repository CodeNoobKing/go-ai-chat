package conversation

// Prompt
// In the domain of ai chatbot, a prompt is a message that is sent to ai.
// Moreover, sometimes human can achieve better result with a appropriate prompt hint.
// Therefore, I made an abstraction over prompt.
// The GetType method can allow the backend make better prompt hint for ai.
type Prompt interface {
	// GetType get the prompt's type.
	GetType() PromptType

	// GetValue Get the message's value.
	GetValue() []byte
}
