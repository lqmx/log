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

	SetTraceId()
	defer DelTraceId()

	log.Debugf(DEBUG.String())
	log.Infof(INFO.String())
	log.Printf(INFO.String())
	log.Warnf(WARN.String())
	log.Errorf(ERROR.String())

	Go(func() {
		log.Debugf(DEBUG.String())
		log.Infof(INFO.String())
		log.Printf(INFO.String())
		log.Warnf(WARN.String())
		log.Errorf(ERROR.String())

		Go(func() {
			log.Debugf(DEBUG.String())
			log.Infof(INFO.String())
			log.Printf(INFO.String())
			log.Warnf(WARN.String())
			log.Errorf(ERROR.String())
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
	log.Panic("test panic")
	log.Info("never log")
}

func TestFatal(t *testing.T) {
	SetLogger(New(Config{
		OutStd:   true,
		DefLevel: DEBUG,
		Module:   "fatal",
	}))
	log.Fatal("test fatal")
	log.Info("never log")
}
