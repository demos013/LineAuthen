package config

// LineURL -
func (c *Config) LineURL() string {
	return getEnvString("LINE_URL", "https://access.line.me/oauth2/v2.1/authorize")
}

// LineLoginURL -
func (c *Config) LineLoginURL() string {
	return getEnvString("LINE_LOGIN_URL", "https://api.line.me/oauth2/v2.1/token")
}

// LineChannel -
func (c *Config) LineChannel() string {
	return getEnvString("LINE_CHANNEL", "1653926216")
}

// LineSecret -
func (c *Config) LineSecret() string {
	return getEnvString("LINE_SECRET", "e14a99c4eae49a3e17fc5b0f755b4bc2")
}

// LineReportURL -
func (c *Config) LineReportURL() string {
	return getEnvString("LINE_REPORT_URL", "https://backend-challenge.line-apps.com/healthcheck/report")
}
