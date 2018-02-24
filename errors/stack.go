package errors

import (
	"fmt"
	"runtime"
)

// Stack represents a stack of program counters
type Stack []uintptr

func (s *Stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				fmt.Fprintf(st, "\n%+v", Frame(pc))
			}
			fmt.Fprintf(st, "\n")
		}
	}
}

func callers() *Stack {
	var pcs [32]uintptr
	n := runtime.Callers(0, pcs[:])
	var st Stack = pcs[:n]
	return &st
}
