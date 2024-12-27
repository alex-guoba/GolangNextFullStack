package domain

import (
	"context"

	"github.com/alex-guoba/gin-clean-template/internal/dao"
	"github.com/alex-guoba/gin-clean-template/internal/entity"
	"github.com/alex-guoba/gin-clean-template/pkg/errcode"

	"gorm.io/gorm"
)

type TagDomain struct {
	ctx context.Context
	db  *gorm.DB
	dao dao.TagDao
}

func NewTagDomain(ctx context.Context, db *gorm.DB) *TagDomain {
	d := &TagDomain{ctx: ctx, db: db}
	d.dao = dao.NewTagDao(db)
	return d
}

func (d *TagDomain) CountTag(name string, state uint8) (int, error) {
	cnt, err := d.dao.CountTag(name, state)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (d *TagDomain) GetTagListWithCnt(name string, state uint8, page, pageSize int) ([]*entity.TagEntity, int, error) {
	cnt, err := d.dao.CountTag(name, state)
	if err != nil {
		return nil, 0, err
	}

	tags, err := d.dao.GetTagList(name, state, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	var tagList []*entity.TagEntity
	for _, tag := range tags {
		tagList = append(tagList, &entity.TagEntity{
			ID:         tag.ID,
			Name:       tag.Name,
			State:      tag.State,
			CreatedBy:  tag.CreatedBy,
			ModifiedBy: tag.ModifiedBy,
		})
	}
	return tagList, int(cnt), nil
}

func (d *TagDomain) CreateTag(name string, state uint8, createdBy string) error {
	return d.dao.CreateTag(name, state, createdBy)
}

func (d *TagDomain) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	return d.dao.UpdateTag(id, name, state, modifiedBy)
}

func (d *TagDomain) DeleteTag(tagID uint32, force bool) error {
	// From a domain calling other domain methods to fulfill the requirements of the business logic
	if !force {
		ad := NewArticleDomain(d.ctx, d.db)
		if cnt, err := ad.countArticleByTagID(tagID); err != nil {
			return err
		} else if cnt > 0 {
			return errcode.ErrTagIDForbidden
		}
	}
	return d.dao.DeleteTag(tagID)
}
