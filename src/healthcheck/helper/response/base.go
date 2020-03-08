package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "wealth_api/errors"
)

// Response -
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Total   int         `json:"total,omitempty"`
}

// New -
func New(statusCode int, input interface{}) *Response {
	msg := http.StatusText(statusCode)

	if err, ok := input.(error); ok {
		res := Response{Error: err.Error(), Status: statusCode, Message: msg}
		return &res
	}

	if res, ok := input.(Response); ok {
		if res.Status == 0 {
			res.Status = statusCode
		}
		if res.Message == "" {
			res.Message = msg
		}
		return &res
	}

	res := Response{Data: input, Status: statusCode, Message: msg}

	if input == nil {
		return &res
	}

	return &res
}

// End -
func End(ctx *gin.Context, statusCode int, output interface{}) {
	res := New(statusCode, output)
	ctx.JSON(statusCode, res)
	ctx.Abort()
}

// Success -
func Success(ctx *gin.Context, output interface{}) {
	End(ctx, http.StatusOK, output)
}

// Error -
func Error(ctx *gin.Context, output interface{}) {
	statusCode := 500

	res := New(statusCode, output)
	ctx.JSON(statusCode, res)
	ctx.Abort()
}
