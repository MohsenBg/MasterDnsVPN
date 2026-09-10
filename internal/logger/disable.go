package logger

import "io"

func (l *Logger) Disable() {
	l.consoleWriter = io.Discard
}
