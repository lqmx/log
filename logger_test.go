package log

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {

	Set(WithLevel(DEBUG), Module("log"), EnableStd(), EnableColor())

	Trace(func() {
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
	})
}

func TestPanic(t *testing.T) {
	Set(Module("panic"), EnableStd(), EnableColor())
	Trace(func() {
		Panic("test panic")
		Info("never log")
	})
}

func TestFatal(t *testing.T) {
	Set(Module("fatal"), EnableStd(), EnableColor())
	Trace(func() {
		Fatal("test fatal")
		Info("never log")
	})
}
