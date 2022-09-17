package log

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

	Log(level Level, option *Option, params ...interface{})
	Logf(level Level, option *Option, format string, params ...interface{})
}

type Config struct {
	OutStd    bool
	DefLevel  Level
	Module    string
	Formatter Formatter
}

type logger struct {
	c *Config
	f Formatter
	w Writer
}

func New(c Config) Logger {
	var l logger
	SetTraceId()

	if c.Module == "" {
		c.Module = "UNKNOWN"
	}
	if c.Formatter == nil {
		c.Formatter = textFormatter
	}

	module = c.Module
	l.c = &c
	return &l
}

func (l logger) Debug(params ...interface{}) {
	write(l, DEBUG, defOption, "", params...)
}

func (l logger) Debugf(format string, params ...interface{}) {
	write(l, DEBUG, defOption, format, params...)
}

func (l logger) Info(params ...interface{}) {
	write(l, INFO, defOption,  "", params...)
}

func (l logger) Infof(format string, params ...interface{}) {
	write(l, INFO, defOption,  format, params...)
}

func (l logger) Print(params ...interface{}) {
	write(l, INFO, defOption,  "", params...)
}

func (l logger) Printf(format string, params ...interface{}) {
	write(l, INFO, defOption,  format, params...)
}

func (l logger) Warn(params ...interface{}) {
	write(l, WARN, defOption,  "", params...)
}

func (l logger) Warnf(format string, params ...interface{}) {
	write(l, WARN, defOption,  format, params...)
}

func (l logger) Error(params ...interface{}) {
	write(l, ERROR, defOption,  "", params...)
}

func (l logger) Errorf(format string, params ...interface{}) {
	write(l, ERROR, defOption,  format, params...)
}

func (l logger) Panic(params ...interface{}) {
	write(l, PANIC, defOption,  "", params...)
	afterPanic(l, defOption)
}

func (l logger) Panicf(format string, params ...interface{}) {
	write(l, PANIC, defOption,  format, params...)
	afterPanic(l, defOption)
}

func (l logger) Fatal(params ...interface{}) {
	write(l, FATAL, defOption,  "", params...)
	afterFatal(l, defOption)
}

func (l logger) Fatalf(format string, params ...interface{}) {
	write(l, FATAL, defOption,  format, params...)
	afterFatal(l, defOption)
}

func (l logger) Log(level Level, option *Option, params ...interface{}) {
	write(l, level, option,  "", params...)
}

func (l logger) Logf(level Level, option *Option, format string, params ...interface{}) {
	write(l, level, option,  format, params...)
}
