package usecase

import (
	"github.com/gin-gonic/gin"
)

// GetHealthCheckStatus -
func (u *Usecase) GetHealthCheckStatus(session *gin.Context) (status HealthcheckStatusOutput, err error) {
	config := u.config
	status.Version = config.AppVersion()

	return status, err
}
