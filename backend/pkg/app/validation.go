package app

import (
	"github.com/alex-guoba/gin-clean-template/pkg/errcode"
	"github.com/alex-guoba/gin-clean-template/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Validation(c *gin.Context, param any, response *Response) error {
	if err := c.ShouldBind(param); err != nil {
		logger.WithTrace(c).Errorf("params errs: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return err
	}
	return nil
}
