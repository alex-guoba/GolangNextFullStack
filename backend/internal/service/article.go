package service

import (
	"context"

	"github.com/alex-guoba/gin-clean-template/internal/domain"
	"github.com/alex-guoba/gin-clean-template/internal/entity"
	"github.com/alex-guoba/gin-clean-template/pkg/app"

	"gorm.io/gorm"
)

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=0"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagIDs        []uint32 `form:"tag_id" binding:"required"`
	Title         string   `form:"title" binding:"required,min=2,max=100"`
	Desc          string   `form:"desc" binding:"required,min=2,max=255"`
	Content       string   `form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageURL string   `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string   `form:"created_by" binding:"required,min=2,max=100"`
	State         uint8    `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type ArticleService struct {
	ctx    context.Context
	db     *gorm.DB
	domain *domain.ArticleDomain
}

func NewArticleService(ctx context.Context, db *gorm.DB) *ArticleService {
	return &ArticleService{
		ctx:    ctx,
		db:     db,
		domain: domain.NewArticleDomain(ctx, db)}
}

func (svc *ArticleService) GetArticle(param *ArticleRequest) (*entity.ArticleEntity, error) {
	article, err := svc.domain.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (svc *ArticleService) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*entity.ArticleEntity, int, error) {
	articleList, cnt, err := svc.domain.GetArticleList(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	return articleList, cnt, nil
}

func (svc *ArticleService) CreateArticle(param *CreateArticleRequest) error {
	return svc.domain.CreateArticle(param.Title, param.Desc,
		param.Content, param.CoverImageURL, param.State, param.CreatedBy, param.TagIDs)
}
