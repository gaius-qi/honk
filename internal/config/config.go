package config

import (
	"fmt"
	"reflect"
)

// Config holds all the runtime config information.
type Config struct {
	// Stock number
	Number string `mapstructure:"number"`
	// Stock market index
	Index IndexType `mapstructure:"index"`
	// Source platform for stock data
	Platform PlatformType `mapstructure:"platform"`
	// Verbose toggles the verbosity
	Debug bool
	// LogLevel is the level with with to log for this config
	LogLevel string `mapstructure:"log_level"`
	// LogFormat is the format that is used for logging
	LogFormat string `mapstructure:"log_format"`
}

type PlatformType string

const (
	SinaPlatformType PlatformType = "sina"
)

func (p *PlatformType) String() string {
	return fmt.Sprint(*p)
}

func (p *PlatformType) Set(value string) error {
	*p = PlatformType(value)
	return nil
}

func (p *PlatformType) Type() string {
	return reflect.TypeOf(p).String()
}

type IndexType string

const (
	ShangHaiIndexType IndexType = "sh"
)

func (i *IndexType) String() string {
	return fmt.Sprint(*i)
}

func (i *IndexType) Set(value string) error {
	*i = IndexType(value)
	return nil
}

func (i *IndexType) Type() string {
	return reflect.TypeOf(i).String()
}

const (
	// DefaultIndex is the default stock market index.
	DefaultIndex = ShangHaiIndexType
	// DefaultPlatform is the default source platform.
	DefaultPlatform = SinaPlatformType
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
