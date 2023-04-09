package prompt

// EnhancementType the type of enhancement.
type EnhancementType string

// Enhancement the enhancement of prompt.
// All filed are readonly.
type Enhancement struct {
	// Type the type of enhancement.
	// nonnull
	Type EnhancementType

	// EnhancerType the type of enhancer.
	// nonnull
	EnhancerType EnhancerType

	// Description the description of enhancement.
	// nullable
	Description string

	// Content the content of enhancement.
	// nonnull
	Content interface{}
}

// EnhancerType the type of enhancer.
type EnhancerType string

// Enhancer is the interface to enhance prompt.
type Enhancer interface {
	// Enhance the prompt content, and return the enhanced value.
	Enhance(ctx *Content, prompt Prompt) error
}

// EnhancerRegistry the registry of enhancer.
type EnhancerRegistry interface {
	// GetEnhancer get the enhancer by type.
	GetEnhancer(enhancerType EnhancerType) Enhancer
}
