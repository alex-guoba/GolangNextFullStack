package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/alex-guoba/gin-clean-template/internal/service"
	"github.com/alex-guoba/gin-clean-template/pkg/app"
	"github.com/alex-guoba/gin-clean-template/pkg/convert"
	"github.com/alex-guoba/gin-clean-template/pkg/errcode"
	"github.com/alex-guoba/gin-clean-template/pkg/logger"
	"github.com/alex-guoba/gin-clean-template/pkg/setting"
)

type Tag struct {
	db  *gorm.DB
	cfg *setting.Configuration
}

func NewTag(db *gorm.DB, cfg *setting.Configuration) Tag {
	return Tag{
		db:  db,
		cfg: cfg,
	}
}

// @Summary	获取多个标签
// @Produce	json
// @Param		name		query		string				false	"标签名称"	maxlength(100)
// @Param		state		query		int					false	"状态"	Enums(0, 1)	default(1)
// @Param		page		query		int					false	"页码"
// @Param		page_size	query		int					false	"每页数量"
// @Success	200			{object}	app.ListResponse	"成功"
// @Failure	400			{object}	errcode.Error		"请求错误"
// @Failure	500			{object}	errcode.Error		"内部错误"
// @Router		/api/v1/tags [get].
func (tag *Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response) != nil {
		return
	}

	svc := service.NewTagService(c.Request.Context(), tag.db)

	rsp := &app.ListResponse{
		Pager: app.Pager{
			Page:     app.GetPage(c),
			PageSize: app.GetPageSize(c, tag.cfg.App.DefaultPageSize, tag.cfg.App.MaxPageSize),
		},
	}

	tags, cnt, err := svc.GetTagListWithCnt(&param, &rsp.Pager)
	if err != nil {
		logger.WithTrace(c).Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	rsp.Pager.TotalRows = cnt
	rsp.List = tags

	response.ToResponse(rsp)
}

// @Summary	新增标签
// @Produce	json
// @Param		name		body		string			true	"标签名称"	minlength(3)	maxlength(100)
// @Param		state		body		int				false	"状态"	Enums(0, 1)		default(1)
// @Param		created_by	body		string			false	"创建者"	minlength(3)	maxlength(100)
// @Success	200			{object}	app.MapResponse	"成功"
// @Failure	400			{object}	errcode.Error	"请求错误"
// @Failure	500			{object}	errcode.Error	"内部错误"
// @Router		/api/v1/tags [post].
func (tag *Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response) != nil {
		return
	}
	svc := service.NewTagService(c.Request.Context(), tag.db)
	if err := svc.CreateTag(&param); err != nil {
		logger.WithTrace(c).Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary	更新标签
// @Produce	json
// @Param		id			path		int				true	"标签ID"
// @Param		name		body		string			false	"标签名称"	minlength(3)	maxlength(100)
// @Param		state		body		int				false	"状态"	Enums(0, 1)		default(1)
// @Param		modified_by	body		string			true	"修改者"	minlength(3)	maxlength(100)
// @Success	200			{array}		app.MapResponse	"成功"
// @Failure	400			{object}	errcode.Error	"请求错误"
// @Failure	500			{object}	errcode.Error	"内部错误"
// @Router		/api/v1/tags/{id} [put].
func (tag *Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response) != nil {
		return
	}

	svc := service.NewTagService(c.Request.Context(), tag.db)
	if err := svc.UpdateTag(&param); err != nil {
		logger.WithTrace(c).Errorf("svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary	删除标签
// @Produce	json
// @Param		id	path		int				true	"标签ID"
// @Success	200	{string}	string			"成功"
// @Failure	400	{object}	errcode.Error	"请求错误"
// @Failure	500	{object}	errcode.Error	"内部错误"
// @Router		/api/v1/tags/{id} [delete].
func (tag *Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response) != nil {
		return
	}

	svc := service.NewTagService(c.Request.Context(), tag.db)
	if err := svc.DeleteTag(&param); err != nil {
		logger.WithTrace(c).Errorf("svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
}
