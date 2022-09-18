package log

import (
	"fmt"
)

type level int8

const (
	DEBUG level = iota - 1
	INFO
	WARN
	ERROR
	PANIC
	FATAL
)

func (l level) String() string {
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

func (l level) ShortString() string {
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

func (l level) ColorShortString() string {
	switch l {
	case DEBUG:
		return "\u001B[106m \u001B[0m"
	case INFO:
		return "\u001B[102m \u001B[0m"
	case WARN:
		return "\u001B[103m \u001B[0m"
	case ERROR:
		return "\u001B[101m \u001B[0m"
	case PANIC:
		return "\u001B[105m \u001B[0m"
	case FATAL:
		return "\u001B[107m \u001B[0m"
	default:
		return "\u001B[104m \u001B[0m"
	}
}
