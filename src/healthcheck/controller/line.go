package controller

import (
	"healthcheck/helper/response"

	"github.com/gin-gonic/gin"
)

type LineAuthenInput struct {
	Code  string `form:"code" json:"code"`
	State string ` form:"state"json:"state"`
}

// GetLineAuthen -
func (c *Controller) GetLineAuthen(ctx *gin.Context) {
	usecase := c.usecase

	var lineInput LineAuthenInput
	if err := ctx.Bind(&lineInput); err != nil {
		response.Error(ctx, err)
		return
	}

	output := response.Response{}
	output.Message = "success"
	usecase.GetLineAuthen(ctx, lineInput.Code, lineInput.State)

	response.Success(ctx, output)
}

// PostLineAuthen -
func (c *Controller) PostLineAuthen(ctx *gin.Context) {
	usecase := c.usecase
	err := usecase.PostLineAuthen(ctx)

	if err != nil {
		response.Error(ctx, err)
		return
	}
	output := response.Response{}
	output.Message = "success"
	response.Success(ctx, output)
}
