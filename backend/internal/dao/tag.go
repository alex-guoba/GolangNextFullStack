package dao

import (
	"gorm.io/gorm"

	"github.com/alex-guoba/gin-clean-template/pkg/app"
)

// type TagSwagger struct {
// 	List  []*Tag
// 	Pager *app.Pager
// }

type TagDaoDB struct {
	engine *gorm.DB
}

type TagDao interface {
	GetTags(ids []uint32, state uint8) ([]TagModel, error)
	GetTagList(name string, state uint8, page, pageSize int) ([]*TagModel, error)
	GetTagListByIDs(ids []uint32, state uint8) ([]*TagModel, error)
	CountTag(name string, state uint8) (int64, error)
	CreateTag(name string, state uint8, createdBy string) error
	UpdateTag(id uint32, name string, state uint8, modifiedBy string) error
	DeleteTag(id uint32) error
}

func NewTagDao(db *gorm.DB) *TagDaoDB {
	return &TagDaoDB{engine: db}
}

func (d *TagDaoDB) GetTags(ids []uint32, state uint8) ([]TagModel, error) {
	var tags []TagModel
	err := d.engine.Where("id in ? AND is_del = ? AND state = ?", ids, 0, state).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tags, err
	}

	return tags, nil
}

func (d *TagDaoDB) GetTagList(name string, state uint8, page, pageSize int) ([]*TagModel, error) {
	var tags []*TagModel
	var err error

	pageOffset := app.GetPageOffset(page, pageSize)
	db := d.engine
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if name != "" {
		db = db.Where("name = ?", name)
	}
	db = db.Where("state = ?", state)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (d *TagDaoDB) GetTagListByIDs(ids []uint32, state uint8) ([]*TagModel, error) {
	var tags []*TagModel
	db := d.engine.Where("state = ? AND is_del = ?", state, 0)
	if err := db.Where("id IN (?)", ids).Find(&tags).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func (d *TagDaoDB) CountTag(name string, state uint8) (int64, error) {
	var count int64
	db := d.engine.Where("state = ? AND is_del = ?", state, 0)
	if name != "" {
		db = db.Where("name = ?", name)
	}
	if err := db.Model(&TagModel{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (d *TagDaoDB) CreateTag(name string, state uint8, createdBy string) error {
	tag := TagModel{
		Name:  name,
		State: state,
		Model: &Model{
			CreatedBy: createdBy,
		},
	}

	// return tag.Create(d.engine)
	return d.engine.Create(&tag).Error
}

func (d *TagDaoDB) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	values := map[string]any{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	// return tag.Update(d.engine, values)
	return d.engine.Model(&TagModel{}).Where("id = ? AND is_del = ?", id, 0).Updates(values).Error
}

func (d *TagDaoDB) DeleteTag(id uint32) error {
	tag := TagModel{Model: &Model{ID: id}}
	return d.engine.Where("id = ? AND is_del = ?", id, 0).Delete(&tag).Error
}
