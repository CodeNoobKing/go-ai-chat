package conversation

import "context"

type Context struct {
	context.Context

	// The conversation's model type.
	ModelType ModelType `json:"modelType"`

	// Is this a new conversation?
	BrandNew bool `json:"brandNew"`

	// The conversation's extra arguments.
	ExtraArgs map[string]interface{} `json:"extraArgs"`
}
