package log

import (
	"fmt"
	"os"
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

func write(logger logger, l Level, o *Option, format string, params ...interface{}) {
	if l < logger.c.DefLevel {
		return
	}

	var m string
	if format == "" {
		m = fmt.Sprint(params...)
	} else {
		m = fmt.Sprintf(format, params...)
	}

	if o.isPrint {
		fmt.Print(string(textFormatter(l, m, o)))
		return
	}

	if logger.c.OutStd {
		fmt.Print(string(textFormatter(l, m, o)))
	}

	if logger.w != nil {
		_, err := logger.w.Write(logger.f(l, m, o))
		if err != nil {
			write(logger, ERROR, &Option{
				isPrint: true,
			}, "err:%v", err)
		}
	}
}

func afterPanic(logger logger, o *Option) {
	printStack(logger, PANIC, o, 4)
	panic("")
}

func afterFatal(logger logger, o *Option) {
	printStack(logger, FATAL, o, 4)
	os.Exit(1)
}

func printStack(logger logger, l Level, o *Option, skip int) {
	for ; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		name := runtime.FuncForPC(pc)
		if name.Name() == "runtime.goexit" {
			break
		}
		write(logger, l, o, "#STACK: %s %s:%d", name.Name(), file, line)
	}
}
