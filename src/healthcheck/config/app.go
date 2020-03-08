package config

// AppBaseURL -
func (c *Config) AppBaseURL() string {
	return getEnvString("APP_BASE_URL", "http://localhost:8080/api/line/authorization")
}

// AppURL -
func (c *Config) AppURL() string {
	return getEnvString("APP_URL", "api")
}

// AppPort -
func (c *Config) AppPort() string {
	return getEnvString("APP_PORT", "8080")
}

// AppVersion -
func (c *Config) AppVersion() string {
	return getEnvString("APP_VERSION", "1.0.0")
}
