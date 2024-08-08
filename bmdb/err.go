package bmdb

import "errors"

var (
	errDbNotFound          = errors.New("db not found")
	errDbNotFoundInContext = errors.New("db not found in context")
	errInputType           = errors.New("input type error")
)

func ErrIsDbNotFound(err error) bool {
	return errors.Is(err, errDbNotFound)
}

func ErrIsDbNotFoundInContext(err error) bool {
	return errors.Is(err, errDbNotFoundInContext)
}

func ErrIsInputType(err error) bool {
	return errors.Is(err, errInputType)
}
