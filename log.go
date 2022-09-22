package log

import "github.com/petermattis/goid"

var (
	log *logger
)

func init() {
	log = newLogger(
		EnableStd(),
		EnableColor(),
		Module("DEFAULT"),
	)
}

func Set(ss ...Setter) {
	log.setConfig(setConfig(ss...))
}

// SetTraceId
// SetTraceId(gid int64, tranceId string)
// SetTraceId(tranceId string)
// SetTraceId(gid int64)
// SetTraceId()
func SetTraceId(values ...interface{}) {
	valLen := len(values)
	if valLen == 0 {
		states.Store(goid.Get(), state{
			traceId: genTraceId(),
		})
	} else if valLen == 1 {
		if gid, ok := values[0].(int64); ok {
			states.Store(gid, state{
				traceId: genTraceId(),
			})
		} else if traceId, ok := values[0].(string); ok {
			states.Store(goid.Get(), state{
				traceId: traceId,
			})
		}
	} else {
		if gid, ok := values[0].(int64); ok {
			if traceId, ok := values[1].(string); ok {
				states.Store(gid, state{
					traceId: traceId,
				})
			}
		}
	}
}

// DelTraceId
// DelTraceId(gid int64)
// DelTraceId()
func DelTraceId(gid ...int64) {
	valLen := len(gid)
	if valLen == 0 {
		states.Delete(goid.Get())
	} else {
		states.Delete(gid[0])
	}
}

func Go(fn func()) {
	s := getState()
	go func() {
		if fn != nil {
			gid := GetGoId()
			SetTraceId(gid, s.traceId+"."+randString(8))
			defer DelTraceId(gid)
			fn()
		}
	}()
}

func Trace(fn func()) {
	if fn != nil {
		gid := GetGoId()
		SetTraceId(gid)
		defer DelTraceId(gid)
		fn()
	}
}

func Must(err error) {
	if err != nil {
		Logf(PANIC, &Option{AddSourceSkip: 1}, "err:%v", err)
	}
}

func Exit(err error) {
	if err != nil {
		Logf(FATAL, &Option{AddSourceSkip: 1}, "err:%v", err)
	}
}

func GetGoId() int64 { return goid.Get() }

func Log(level level, option *Option, params ...interface{}) { log.Log(level, option, params...) }
func Logf(level level, option *Option, format string, params ...interface{}) {
	log.Logf(level, option, format, params...)
}
func Debug(params ...interface{})                 { log.Debug(params...) }
func Debugf(format string, params ...interface{}) { log.Debugf(format, params...) }

func Info(params ...interface{}) { log.Info(params...) }

func Infof(format string, params ...interface{}) { log.Infof(format, params...) }

func Print(params ...interface{}) { log.Print(params...) }

func Printf(format string, params ...interface{}) { log.Printf(format, params...) }

func Warn(params ...interface{}) { log.Warn(params...) }

func Warnf(format string, params ...interface{}) { log.Warnf(format, params...) }

func Error(params ...interface{}) { log.Error(params...) }

func Errorf(format string, params ...interface{}) { log.Errorf(format, params...) }

func Panic(params ...interface{}) { log.Panic(params...) }

func Panicf(format string, params ...interface{}) { log.Panicf(format, params...) }

func Fatal(params ...interface{}) { log.Fatal(params...) }

func Fatalf(format string, params ...interface{}) { log.Fatalf(format, params...) }
