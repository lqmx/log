package log

import (
	"fmt"
)

type Level int8

const (
	DEBUG Level = iota - 1
	INFO
	WARN
	ERROR
	PANIC
	FATAL
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	case PANIC:
		return "panic"
	case FATAL:
		return "fatal"
	default:
		return fmt.Sprintf("level(%d)", l)
	}
}

func (l Level) ShortString() string {
	switch l {
	case DEBUG:
		return "D"
	case INFO:
		return "I"
	case WARN:
		return "W"
	case ERROR:
		return "E"
	case PANIC:
		return "P"
	case FATAL:
		return "F"
	default:
		return fmt.Sprintf("%d", l)
	}
}

func (l Level) ColorShortString() string {
	switch l {
	case DEBUG:
		return "\u001B[106mD\u001B[0m"
	case INFO:
		return "\u001B[102mI\u001B[0m"
	case WARN:
		return "\u001B[103mW\u001B[0m"
	case ERROR:
		return "\u001B[101mE\u001B[0m"
	case PANIC:
		return "\u001B[105mP\u001B[0m"
	case FATAL:
		return "\u001B[107mF\u001B[0m"
	default:
		return fmt.Sprintf("\u001B[104m%d\u001B[0m", l)
	}
}
