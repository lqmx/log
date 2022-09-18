package log

type config struct {
	level  level
	std    bool
	module string
	fmt    Formatter
	color  bool
}

type Setter func(c *config)

func Level(l level) Setter {
	return func(c *config) {
		c.level = l
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