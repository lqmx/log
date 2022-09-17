package log

import (
	"os"
	"strconv"
	"sync"
)

var (
	module string
	pid    string
	states sync.Map
)

func init() {
	pid = strconv.Itoa(os.Getpid())
}

type state struct {
	traceId string
}

func getState(gids ...int64) state {
	var gid int64
	if len(gids) == 0 {
		gid = GetGoId()
	} else {
		gid = gids[0]
	}
	if gid == 0 {
		return state{}
	}
	v, ok := states.Load(gid)
	if !ok {
		return state{}
	}
	return v.(state)
}
