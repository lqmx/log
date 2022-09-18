package log

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func getPackageName(f string) (string, string) {
	slashIndex := strings.LastIndex(f, "/")
	if slashIndex > 0 {
		idx := strings.Index(f[slashIndex:], ".") + slashIndex
		return f[:idx], f[idx+1:]
	}
	return f, ""
}

func source(skip int) string {
	var callerName, callerFile string
	var callerLine int
	var ok bool
	var pc uintptr
	pc, callerFile, callerLine, ok = runtime.Caller(skip)
	callerName = ""
	if ok {
		callerName = runtime.FuncForPC(pc).Name()
	}
	filePath, _ := getPackageName(callerName)
	return fmt.Sprintf("%s:%d:%s", path.Join(filePath, path.Base(callerFile)), callerLine, callerName)
}

func write(logger logger, l level, o *Option, format string, params ...interface{}) {
	if l < logger.c.level {
		return
	}

	var m string
	if format == "" {
		m = fmt.Sprint(params...)
	} else {
		m = fmt.Sprintf(format, params...)
	}

	if logger.c.std {
		fmt.Print(string(stdFormatter(l, m, o)))
	}

	if logger.w != nil {
		_, err := logger.w.Write(logger.c.fmt(l, m, o))
		if err != nil {
			write(logger, ERROR, o, "err:%v", err)
		}
	}
}

func printStack(logger logger, l level, o *Option, skip int) {
	for ; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		name := runtime.FuncForPC(pc)
		if name.Name() == "runtime.goexit" {
			break
		}
		write(logger, l, &Option{
			c: o.c,
			AddSourceSkip: 1,
		}, "#STACK: %s %s:%d", name.Name(), file, line)
	}
}
