package general

type Error struct {
	message string
	code    int
}

func NewError(message string, code int) *Error {
	error := &Error{}
	error.message = message
	error.code = code
	return error
}
