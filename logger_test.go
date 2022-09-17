package log

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	SetLogger(New(Config{
		OutStd:   true,
		DefLevel: DEBUG,
		Module:   "log",
	}))

	Debugf(DEBUG.String())
	Infof(INFO.String())
	Printf(INFO.String())
	Warnf(WARN.String())
	Errorf(ERROR.String())

	Go(func() {
		Debugf(DEBUG.String())
		Infof(INFO.String())
		Printf(INFO.String())
		Warnf(WARN.String())
		Errorf(ERROR.String())

		Go(func() {
			Debugf(DEBUG.String())
			Infof(INFO.String())
			Printf(INFO.String())
			Warnf(WARN.String())
			Errorf(ERROR.String())
		})
	})
	time.Sleep(time.Second)
}

func TestPanic(t *testing.T) {
	SetLogger(New(Config{
		OutStd:   true,
		DefLevel: DEBUG,
		Module:   "panic",
	}))
	Panic("test panic")
	Info("never log")
}

func TestFatal(t *testing.T) {
	SetLogger(New(Config{
		OutStd:   true,
		DefLevel: DEBUG,
		Module:   "fatal",
	}))
	Fatal("test fatal")
	Info("never log")
}
