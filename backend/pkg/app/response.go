package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alex-guoba/gin-clean-template/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type ResponseSuccess struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

type ListResponse struct {
	List  any   `json:"list"`
	Pager Pager `json:"pager"`
}

type MapResponse map[string]string

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data any) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, &ResponseSuccess{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	r.Ctx.JSON(err.Status, err)
}
