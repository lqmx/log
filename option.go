package log

var (
	defOption         = &Option{}
	defaultSourceSkip = 4
)

type Option struct {
	AddSourceSkip int
	SourceString  string
	NoColor       bool
	isPrint       bool
}
