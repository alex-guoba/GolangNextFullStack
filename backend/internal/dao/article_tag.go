package dao

import (
	"gorm.io/gorm"
)

type ArticleTagDaoDB struct {
	engine *gorm.DB
}

type ArticleTagDao interface {
	GetArticleTagsByAID(articleID uint32) ([]ArticleTagModel, error)
	GetArticleTagListByTID(tagID uint32) ([]*ArticleTagModel, error)
	GetArticleTagListByAIDs(articleIDs []uint32) ([]*ArticleTagModel, error)
	CreateArticleTag(articleID, tagID uint32, createdBy string) error
	UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error
	DeleteArticleTag(articleID uint32) error
}

func NewArticleTagDao(db *gorm.DB) *ArticleTagDaoDB {
	return &ArticleTagDaoDB{engine: db}
}

// TODO: combine single query and batch query.
func (d *ArticleTagDaoDB) GetArticleTagsByAID(articleID uint32) ([]ArticleTagModel, error) {
	var articleTag []ArticleTagModel
	if err := d.engine.Where("article_id = ? AND is_del = ?", articleID, 0).Find(&articleTag).Error; err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}

	return articleTag, nil
}

func (d *ArticleTagDaoDB) GetArticleTagListByTID(tagID uint32) ([]*ArticleTagModel, error) {
	var articleTags []*ArticleTagModel
	if err := d.engine.Where("tag_id = ? AND is_del = ?", tagID, 0).Find(&articleTags).Error; err != nil {
		return nil, err
	}

	return articleTags, nil
}

func (d *ArticleTagDaoDB) GetArticleTagListByAIDs(articleIDs []uint32) ([]*ArticleTagModel, error) {
	var ats []*ArticleTagModel
	err := d.engine.Where("article_id IN (?) AND is_del = ?", articleIDs, 0).Find(&ats).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return ats, nil
}

func (d *ArticleTagDaoDB) CreateArticleTag(articleID, tagID uint32, createdBy string) error {
	at := ArticleTagModel{
		Model: &Model{
			CreatedBy: createdBy,
		},
		ArticleID: articleID,
		TagID:     tagID,
	}
	return d.engine.Create(&at).Error
}

func (d *ArticleTagDaoDB) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	values := map[string]any{
		"tag_id":      tagID,
		"modified_by": modifiedBy,
	}
	return d.engine.Model(&ArticleTagModel{}).Where("article_id = ? AND is_del = ?", articleID, 0).Limit(1).Updates(values).Error
}

func (d *ArticleTagDaoDB) DeleteArticleTag(articleID uint32) error {
	articleTag := ArticleTagModel{ArticleID: articleID}
	return d.engine.Where("article_id = ? AND is_del = ?", articleID, 0).Delete(&articleTag).Limit(1).Error
}
