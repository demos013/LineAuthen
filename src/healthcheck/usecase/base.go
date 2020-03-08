package usecase

import (
	"healthcheck/config"

	"github.com/gin-gonic/gin"
)

// HealthcheckStatusOutput -
type HealthcheckStatusOutput struct {
	Version string `json:"version"`
}

// Interface -
type Interface interface {
	// health_usecase.go
	GetHealthCheckStatus(*gin.Context) (HealthcheckStatusOutput, error)
	PostLineAuthen(*gin.Context) error
	GetLineAuthen(*gin.Context, string, string) error
}

// Usecase -
type Usecase struct {
	config config.Interface
}

var _ Interface = &Usecase{}

// New -
func New(config config.Interface) *Usecase {
	return &Usecase{config: config}
}
