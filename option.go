package log

var (
	defaultSourceSkip = 6
)

type Option struct {
	AddSourceSkip int
	SourceString  string

	c config
}

func newDefOption() *Option {
	return &Option{}
}
