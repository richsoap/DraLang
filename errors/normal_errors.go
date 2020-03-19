package errors

import "errors"

var (
	ERR_TOO_MANY_PARAMS = errors.New("error: too many params")
	ERR_EOF = errors.New("error: EOF")
)
