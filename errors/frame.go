package errors

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"strings"
)

// Frame represents a program counter inside a stack frame
type Frame uintptr

func (f Frame) pc() uintptr { return uintptr(f) - 1 }

func (f Frame) line() int {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return 0
	}

	_, ln := fn.FileLine(f.pc())
	return ln
}

func (f Frame) file() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}

	file, _ := fn.FileLine(f.pc())
	return file
}

// Format formats the frame according to the fmt.Formatter interface
func (f Frame) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		f.Format(st, 's')
		io.WriteString(st, ":")
		f.Format(st, 'd')
	case 's':
		switch {
		case st.Flag('+'):
			pc := f.pc()
			fn := runtime.FuncForPC(pc)
			if fn == nil {
				io.WriteString(st, "unknown")
			} else {
				file, _ := fn.FileLine(f.pc())
				fmt.Fprintf(st, "%s\n\t%s", fn.Name(), file)
			}
		default:
			io.WriteString(st, path.Base(f.file()))
		}
	case 'd':
		fmt.Fprintf(st, "%d", f.line())
	case 'n':
		name := runtime.FuncForPC(f.pc()).Name()
		io.WriteString(st, funcName(name))
	}
}

func funcName(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}
