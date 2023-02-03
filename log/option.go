package log

type config struct {
	serviceName string
	logFile     string
	level       string
	maxSize     int
	maxAge      int
	maxBackups  int
	compress    bool
}

var _config = config{
	serviceName: "",
	logFile:     "",
	level:       "debug",
	maxSize:     100,
	maxAge:      7,
	maxBackups:  100,
	compress:    false,
}

type Option func(*config)

func WithServiceName(serviceName string) Option {
	return func(c *config) {
		c.serviceName = serviceName
	}
}

func WithLogFile(logFile string) Option {
	return func(c *config) {
		c.logFile = logFile
	}
}

func WithLevel(level string) Option {
	return func(c *config) {
		c.level = level
	}
}

func WithMaxSize(maxSize int) Option {
	return func(c *config) {
		c.maxSize = maxSize
	}
}

func WithMaxAge(maxAge int) Option {
	return func(c *config) {
		c.maxAge = maxAge
	}
}

func WithMaxBackups(maxBackups int) Option {
	return func(c *config) {
		c.maxBackups = maxBackups
	}
}

func WithCompress(compress bool) Option {
	return func(c *config) {
		c.compress = compress
	}
}
