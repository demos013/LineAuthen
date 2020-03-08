package config

import (
	"os"
)

// Interface -
type Interface interface {
	// app.go
	AppBaseURL() string
	AppURL() string
	AppPort() string
	AppVersion() string

	// debug.go
	GinMode() string
	GormLogEnable() bool

	//line.go
	LineURL() string
	LineLoginURL() string
	LineReportURL() string
	LineChannel() string
	LineSecret() string
}

// Config -
type Config struct {
}

var _ Interface = &Config{}

// New -
func New() Interface {
	return &Config{}
}

func getEnvString(name string, defaultValue string) string {
	if val := os.Getenv(name); val != "" {
		return val
	}
	return defaultValue
}
