package users

// Error custom error type for the users domains
type Error struct {
	errCode    string
	errMessage string
}

func newError(code string, message string) error {
	return Error{errCode: code, errMessage: message}
}

func (e Error) Error() string {
	return e.errMessage
}

func (e Error) Code() string {
	return e.errCode
}
