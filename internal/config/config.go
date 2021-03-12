package config

// Config holds all the runtime config information.
type Config struct {
	// Stock market index
	Index string `mapstructure:"index"`
	// Source platform for stock data
	Platform string `mapstructure:"platform"`
	// Verbose toggles the verbosity
	Debug bool
	// LogLevel is the level with with to log for this config
	LogLevel string `mapstructure:"log_level"`
	// LogFormat is the format that is used for logging
	LogFormat string `mapstructure:"log_format"`
}

const (
	// DefaultIndex is the default stock market index.
	DefaultIndex = "sina"
	// DefaultPlatform is the default source platform.
	DefaultPlatform = "sina"
	// DefaultLogLevel is the default logging level.
	DefaultLogLevel = "warn"
	// DefaultLogFormat is the default format of the logger
	DefaultLogFormat = "text"
	// DefaultDebug is the default debug status.
	DefaultDebug = false
)

// New returns a new Config
func New() *Config {
	return &Config{
		Index:     DefaultIndex,
		Platform:  DefaultPlatform,
		Debug:     DefaultDebug,
		LogLevel:  DefaultLogLevel,
		LogFormat: DefaultLogFormat,
	}
}
