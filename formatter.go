package log

import (
	"bytes"
	"fmt"
	"time"
)

type Formatter func(level level, msg string, option *Option) []byte

func stdFormatter(level level, msg string, option *Option) []byte {
	var ss string
	if option.SourceString != "" {
		ss = option.SourceString
	} else {
		ss = source(option.AddSourceSkip + defaultSourceSkip)
	}

	var ls = level.ShortString()
	if option.c.color {
		ls = level.ColorShortString()
	}

	gid := GetGoId()
	s := getState(gid)
	now := time.Now()

	traceId := s.traceId
	if traceId != "" {
		traceId = now.Format("0215")+traceId
	}

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("%s %s(%s,%d) %s <%s> %s %s\n",
		ls, option.c.module, pid, gid,
		now.Format("06-01-02T15:04:05.0000"),
		traceId,
		ss, msg,
	))

	return b.Bytes()
}
