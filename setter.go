package log

type config struct {
	level   Level
	psLevel Level
	std     bool
	module  string
	fmt     Formatter
	color   bool
	writer  Writer
}

type Setter func(c *config)

func WithLevel(l Level) Setter {
	return func(c *config) {
		c.level = l
	}
}

func WithPsLevel(l Level) Setter {
	return func(c *config) {
		c.psLevel = l
	}
}

func WithWriter(w Writer) Setter {
	return func(c *config) {
		c.writer = w
	}
}

func EnableStd() Setter {
	return func(c *config) {
		c.std = true
	}
}

func DisableStd() Setter {
	return func(c *config) {
		c.std = false
	}
}

func Module(m string) Setter {
	return func(c *config) {
		c.module = m
	}
}

func Fmt(f Formatter) Setter {
	return func(c *config) {
		c.fmt = f
	}
}

func EnableColor() Setter {
	return func(c *config) {
		c.color = true
	}
}

func DisableColor() Setter {
	return func(c *config) {
		c.color = false
	}
}
