package http

import "net/http"

type publicErrors map[string]int

func (p publicErrors) isPublic(err error) bool {
	type coder interface {
		Code() string
	}

	if e, ok := err.(coder); ok {
		if _, ok := p[e.Code()]; ok {
			return true
		}
	}

	return false
}

func (p publicErrors) Code(err error) string {
	type coder interface {
		Code() string
	}

	if e, ok := err.(coder); ok {
		return e.Code()
	}

	return ""

}

func (p publicErrors) HTTPStatusCode(err error) int {
	type coder interface {
		Code() string
	}

	if e, ok := err.(coder); ok {
		if _, ok := p[e.Code()]; ok {
			return p[e.Code()]
		}
	}

	return http.StatusInternalServerError

}

type sessionError struct {
	code string
	msg  string
}

func (s sessionError) Error() string {
	return s.msg
}

func (s sessionError) Code() string {
	return s.code
}

type handlerError struct {
	allowedList map[string]int
	originalErr error
}

func (h handlerError) Error() string {
	return h.originalErr.Error()
}

func (h handlerError) Code() string {
	type coder interface {
		Code() string
	}

	if e, ok := h.originalErr.(coder); ok {
		return e.Code()
	}

	return ""
}

func (h handlerError) Public() bool {
	type coder interface {
		Code() string
	}

	if e, ok := h.originalErr.(coder); ok {
		if _, ok := h.allowedList[e.Code()]; ok {
			return true
		}
	}

	return false
}

func (h handlerError) HTTPStatusCode() int {
	type coder interface {
		Code() string
	}

	if e, ok := h.originalErr.(coder); ok {
		if _, ok := h.allowedList[e.Code()]; ok {
			return h.allowedList[e.Code()]
		}
	}

	return http.StatusInternalServerError
}
