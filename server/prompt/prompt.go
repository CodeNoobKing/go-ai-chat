package prompt

// ContentType the type of prompt.
type ContentType string

// ModelType the type of ai-model.
type ModelType string

// Content the content of prompt.
// The content may be text, image, audio, video, etc.
type Content interface {
	// GetType get the prompt's type.
	GetType() ContentType

	// GetValue Get the message's value.
	GetValue() <-chan []byte
}

// Prompt
// A prompt is a message that is sent to ai.
// The GetType method can allow the enhancer make better prompt hint for ai.
// In practice, the prompt may contain various content: text, image, audio, video, etc.
// Therefore, the prompt is a collection of content, to support various content type.
// Usually, the prompt contains only one content.
type Prompt struct {
	// ModelType the model type would be used.
	ModelType ModelType

	// Contents the content of prompt.
	Contents []Content
}

// ReadAllContent read all content from the prompt.
func ReadAllContent(content Content) []byte {
	var buf []byte
	for chunk := range content.GetValue() {
		buf = append(buf, chunk...)
	}
	return buf
}
