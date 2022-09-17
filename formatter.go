package log

import (
	"bytes"
	"fmt"
	"time"
)

type Formatter func(level Level, msg string, option *Option) []byte

func textFormatter(level Level, msg string, option *Option) []byte {
	if option == nil {
		option = defOption
	}

	var sourceString string
	if option.SourceString != "" {
		sourceString = option.SourceString
	} else {
		sourceString = source(option.AddSourceSkip + defaultSourceSkip)
	}

	levelString := level.ShortString()
	if !option.NoColor {
		levelString = level.ColorShortString()
	}

	gid := GetGoId()
	s := getState(gid)
	now := time.Now()

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("%s %s(%s,%d) %s <%s> %s %s\n",
		levelString, module, pid, gid,
		now.Format("06-01-02T15:04:05.0000"),
		now.Format("0215-")+s.traceId,
		sourceString, msg,
	))

	return b.Bytes()
}
