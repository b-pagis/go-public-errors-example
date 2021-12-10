package legacy

type internalError struct {
	code string
	msg  string
}

func (i internalError) Error() string {
	return i.msg
}

func (i internalError) Code() string {
	return i.code
}

func (i internalError) NotFound() bool {
	return i.code == "notFound"
}
