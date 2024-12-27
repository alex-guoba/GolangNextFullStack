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

type Article struct {
	db  *gorm.DB
	cfg *setting.Configuration
}

func NewArticle(db *gorm.DB, cfg *setting.Configuration) Article {
	return Article{
		db:  db,
		cfg: cfg,
	}
}

// @Summary	创建文章
// @Produce	json
// @Param		tag_ids			body		[]string			true	"标签ID"
// @Param		title			body		string			true	"文章标题"
// @Param		desc			body		string			false	"文章简述"
// @Param		cover_image_url	body		string			true	"封面图片地址"
// @Param		content			body		string			true	"文章内容"
// @Param		created_by		body		int				true	"创建者"
// @Param		state			body		int				false	"状态"
// @Success	200				{object}	string			"成功"
// @Failure	400				{object}	errcode.Error	"请求错误"
// @Failure	500				{object}	errcode.Error	"内部错误"
// @Router		/api/v1/articles [post].
func (art *Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response) != nil {
		return
	}

	svc := service.NewArticleService(c.Request.Context(), art.db)
	err := svc.CreateArticle(&param)
	if err != nil {
		logger.WithTrace(c).Errorf("svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary	获取单个文章
// @Produce	json
// @Param		id	path		int						true	"文章ID"
// @Success	200	{object}	entity.ArticleEntity	"成功"
// @Failure	400	{object}	errcode.Error			"请求错误"
// @Failure	500	{object}	errcode.Error			"内部错误"
// @Router		/api/v1/articles/{id} [get].
func (art *Article) Get(c *gin.Context) {
	param := service.ArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response) != nil {
		return
	}

	svc := service.NewArticleService(c.Request.Context(), art.db)
	article, err := svc.GetArticle(&param)
	if err != nil {
		logger.WithTrace(c).Errorf("svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(article)
}

// @Summary	获取多个文章
// @Produce	json
// @Param		tag_id		query		int					false	"标签ID"
// @Param		state		query		int					false	"状态"
// @Param		page		query		int					false	"页码"
// @Param		page_size	query		int					false	"每页数量"
// @Success	200			{object}	app.ListResponse	"成功"
// @Failure	400			{object}	errcode.Error		"请求错误"
// @Failure	500			{object}	errcode.Error		"内部错误"
// @Router		/api/v1/articles [get].
func (art *Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response) != nil {
		return
	}

	rsp := &app.ListResponse{
		Pager: app.Pager{
			Page:     app.GetPage(c),
			PageSize: app.GetPageSize(c, art.cfg.App.DefaultPageSize, art.cfg.App.MaxPageSize),
		},
	}

	svc := service.NewArticleService(c.Request.Context(), art.db)
	articles, totalRows, err := svc.GetArticleList(&param, &rsp.Pager)
	if err != nil {
		logger.WithTrace(c).Errorf("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}

	rsp.Pager.TotalRows = totalRows
	rsp.List = articles
	response.ToResponse(rsp)
}
