package log

func SetLevel() {}

func Debug() {}

func Debugf() {}

func Info() {}

func Infof() {}

func Warn() {}

func Warnf() {}

func Fatal() {}

func Fatalf() {}

func DebugFlag(msg string, flag bool) {
	if flag {
		println(msg)
	}
}
