package errors

import (
	"fmt"
	"io"
	"strings"
)

type Error struct {
	msg   string
	stack *Stack

	cause error
}

func (e *Error) Error() string { return e.msg }

func (e *Error) Cause() error { return e.cause }

func (e *Error) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		if st.Flag('+') {
			io.WriteString(st, e.msg)
			e.stack.Format(st, verb)

			if e.Cause() != nil {
				fmt.Fprintf(st, "caused by:\n%+v\n", e.Cause())
			}

			return
		}

		fallthrough
	case 's':
		io.WriteString(st, e.msg)
	case 'q':
		fmt.Fprintf(st, "%q", e.msg)
	}
}

// New returns an error with the supplied message, and it also records the stack trace at the point it is called
func New(format string, a ...interface{}) error {
	return newError(format, a...)
}

func newError(format string, a ...interface{}) *Error {
	msg := strings.TrimSpace(fmt.Sprintf(format, a...))
	if msg == "" {
		return nil
	}

	return &Error{
		msg:   msg,
		stack: callers(),
	}
}

func Wrap(e error, format string, a ...interface{}) error {
	if e == nil {
		return nil
	}

	err := newError(format, a...)
	if err == nil {
		return nil
	}

	err.cause = e

	return err
}
