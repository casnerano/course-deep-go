package errors

import (
	"fmt"
	"strings"
)

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	errMessages := make([]string, len(e.errors))
	for index, err := range e.errors {
		errMessages[index] = err.Error()
	}

	return fmt.Sprintf(
		"%d errors occured:\n\t* %s\n",
		len(e.errors),
		strings.Join(errMessages, "\t* "),
	)
}

func Append(err error, errs ...error) *MultiError {
	var multiErr *MultiError

	if err == nil {
		multiErr = &MultiError{}
	} else {
		switch me := err.(type) {
		case *MultiError:
			multiErr = me
		default:
			multiErr = &MultiError{errors: []error{me}}
		}
	}

	for _, e := range errs {
		if e != nil {
			multiErr.errors = append(multiErr.errors, e)
		}
	}

	return multiErr
}
