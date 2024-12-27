package service

import (
	"context"

	"github.com/alex-guoba/gin-clean-template/internal/domain"
	"github.com/alex-guoba/gin-clean-template/internal/entity"
	"github.com/alex-guoba/gin-clean-template/pkg/app"

	"gorm.io/gorm"
)

// for validateor.
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// service loyer design principle:
// 1. domain field should not separate with each other
// 2. service layer should not know the implementation of DAO layer
// 3. service layer should be replaceable to other service layer

// service.
type TagService struct {
	ctx    context.Context
	db     *gorm.DB
	domain *domain.TagDomain
}

func NewTagService(ctx context.Context, db *gorm.DB) *TagService {
	svc := &TagService{ctx: ctx, db: db}
	svc.domain = domain.NewTagDomain(ctx, db)
	return svc
}

func (svc *TagService) CountTag(param *CountTagRequest) (int, error) {
	return svc.domain.CountTag(param.Name, param.State)
}

func (svc *TagService) GetTagListWithCnt(param *TagListRequest, pager *app.Pager) ([]*entity.TagEntity, int, error) {
	return svc.domain.GetTagListWithCnt(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *TagService) CreateTag(param *CreateTagRequest) error {
	return svc.domain.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *TagService) UpdateTag(param *UpdateTagRequest) error {
	return svc.domain.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *TagService) DeleteTag(param *DeleteTagRequest) error {
	return svc.domain.DeleteTag(param.ID, false)
}
