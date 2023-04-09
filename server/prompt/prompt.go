package prompt

import (
	"io"
)

// ContentType the type of prompt.
type ContentType string

// Content the content of prompt.
// The content may be text, image, audio, video, etc.
type Content interface {
	// GetType get the prompt's type.
	GetType() ContentType

	io.ReadCloser
}

// Prompt
// A prompt is a message that is sent to ai.
// The GetType method can allow the enhancer make better prompt hint for ai.
// In practice, the prompt may contain various content: text, image, audio, video, etc.
// Therefore, the prompt is a collection of content, to support various content type.
// Usually, the prompt contains only one content.
type Prompt struct {
	Contents []Content
}
