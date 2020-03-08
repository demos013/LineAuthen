package controller

import (
	"healthcheck/helper/response"

	"github.com/gin-gonic/gin"
)

// GetHealthCheckStatus -
func (c *Controller) GetHealthCheckStatus(ctx *gin.Context) {

	usecase := c.usecase
	status, err := usecase.GetHealthCheckStatus(ctx)

	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, status)
}
