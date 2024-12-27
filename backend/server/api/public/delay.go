package public

import (
	"time"

	"github.com/alex-guoba/gin-clean-template/pkg/app"
	"github.com/alex-guoba/gin-clean-template/pkg/convert"

	"github.com/gin-gonic/gin"
)

type Delay struct{}

type DelayRequest struct {
	Seconds int `form:"seconds" binding:"required,gte=1,lte=60"`
}

type DelayResponse map[string]int

// @Summary	Delay for the specified period
// @Produce	json
// @Param		seconds	path		int				true	"seconds to wait for"
// @Success	200		{object}	string			"成功"
// @Failure	400		{object}	errcode.Error	"请求错误"
// @Failure	500		{object}	errcode.Error	"内部错误"
// @Router		/public/delay/{seconds} [get]
func (*Delay) DelayHandler(c *gin.Context) {
	param := DelayRequest{Seconds: convert.StrTo(c.Param("seconds")).MustInt()}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response) != nil {
		return
	}

	time.Sleep(time.Duration(param.Seconds) * time.Second)

	response.ToResponse(DelayResponse{"delay": param.Seconds})
}
