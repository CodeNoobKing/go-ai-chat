package conversation

// The type of prompt.
type PromptType string

const (
	// NaturalLanguagePromptType the text prompt type.
	NaturalLanguagePromptType PromptType = "NaturalLanguage"
)

// ConversationID the id of conversation.
type ConversationID string

// ModelType the type of ai-model.
type ModelType string

const (
	// GPT3TurboModelType the gpt3 model type.
	GPT3TurboModelType ModelType = "GPT3-Turbo"
)
