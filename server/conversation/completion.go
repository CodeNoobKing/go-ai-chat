package conversation

type Completion interface {
	// GetChunkedValue
	// Because the completion is usually provided in a stream style, therefore I made an abstraction over streeamed message.
	// The first return value is the chunk of the completion.
	// The second return value is a flag that indicates whether the stream is ended.
	// The third return value is potential error.
	GetChunkedValue(ctx *Context) ([]byte, bool, error)

	// Terminate
	// Stop the generation of completion if it is not desired.
	Terminate(ctx *Context) error
}
