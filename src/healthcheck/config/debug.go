package config

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

// GinMode -
func (c *Config) GinMode() string {
	mode := getEnvString("GIN_MODE", gin.DebugMode)
	if funk.ContainsString([]string{gin.ReleaseMode, gin.DebugMode}, mode) {
		return mode
	}
	return gin.ReleaseMode
}

// GormLogEnable -
func (c *Config) GormLogEnable() bool {
	debug := getEnvString("GORM_LOG_ENABLE", "1")
	return debug == "1"
}
