package log

var (
	defOption         = &Option{}
	defaultSourceSkip = 5
)

type Option struct {
	AddSourceSkip int
	SourceString  string
	NoColor       bool
	isPrint       bool
}
