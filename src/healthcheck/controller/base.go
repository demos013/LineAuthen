package controller

import (
	"healthcheck/usecase"
)

// Controller -
type Controller struct {
	usecase usecase.Interface
}

// New -
func New(usecase usecase.Interface) *Controller {
	return &Controller{usecase: usecase}
}
