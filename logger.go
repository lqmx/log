package log

import "os"

type Logger interface {
	Debug(params ...interface{})
	Debugf(format string, params ...interface{})

	Info(params ...interface{})
	Infof(format string, params ...interface{})

	Print(params ...interface{})
	Printf(format string, params ...interface{})

	Warn(params ...interface{})
	Warnf(format string, params ...interface{})

	Error(params ...interface{})
	Errorf(format string, params ...interface{})

	Panic(params ...interface{})
	Panicf(format string, params ...interface{})

	Fatal(params ...interface{})
	Fatalf(format string, params ...interface{})

	Log(level level, option *Option, params ...interface{})
	Logf(level level, option *Option, format string, params ...interface{})
}

type logger struct {
	c *config
	w Writer
}

func newLogger(ss ...Setter) *logger {
	var l logger

	for _, s := range ss {
		s(&c)
	}
	l.c = &c

	return &l
}

func (l logger) setConfig(c config) {
	l.c = &c
}

func (l logger) Debug(params ...interface{}) {
	l.log(DEBUG, nil, "", params...)
}

func (l logger) Debugf(format string, params ...interface{}) {
	l.log(DEBUG, nil, format, params...)
}

func (l logger) Info(params ...interface{}) {
	l.log(INFO, nil, "", params...)
}

func (l logger) Infof(format string, params ...interface{}) {
	l.log(INFO, nil, format, params...)
}

func (l logger) Print(params ...interface{}) {
	l.log(INFO, nil, "", params...)
}

func (l logger) Printf(format string, params ...interface{}) {
	l.log(INFO, nil, format, params...)
}

func (l logger) Warn(params ...interface{}) {
	l.log(WARN, nil, "", params...)
}

func (l logger) Warnf(format string, params ...interface{}) {
	l.log(WARN, nil, format, params...)
}

func (l logger) Error(params ...interface{}) {
	l.log(ERROR, nil, "", params...)
}

func (l logger) Errorf(format string, params ...interface{}) {
	l.log(ERROR, nil, format, params...)
}

func (l logger) Panic(params ...interface{}) {
	l.log(PANIC, nil, "", params...)
}

func (l logger) Panicf(format string, params ...interface{}) {
	l.log(PANIC, nil, format, params...)
}

func (l logger) Fatal(params ...interface{}) {
	l.log(FATAL, nil, "", params...)
}

func (l logger) Fatalf(format string, params ...interface{}) {
	l.log(FATAL, nil, format, params...)
}

func (l logger) Log(level level, option *Option, params ...interface{}) {
	l.log(level, option, "", params...)
}

func (l logger) Logf(level level, option *Option, format string, params ...interface{}) {
	l.log(level, option, format, params...)
}

func (l logger) log(level level, option *Option, format string, params ...interface{}) {
	if option == nil {
		option = newDefOption()
	}
	option.c = *l.c

	write(l, level, option, format, params...)

	if level == FATAL {
		printStack(l, PANIC, option, 4)
		os.Exit(1)
	} else if level == PANIC {
		printStack(l, PANIC, option, 4)
		panic("")
	}
}
