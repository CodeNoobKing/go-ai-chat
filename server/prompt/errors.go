package prompt

type ErrorCode string

func (e ErrorCode) Code() string {
	return string(e)
}

const (
	ErrCodeTypeNotFound = ErrorCode("TypeNotFound")
)
