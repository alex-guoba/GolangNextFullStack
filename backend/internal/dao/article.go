package dao

import (
	"gorm.io/gorm"
)

type ArticleDaoDB struct {
	engine *gorm.DB
}

type ArticleDao interface {
	CreateArticle(title string, desc string, content string, image string,
		state uint8, createdBy string) (*ArticleModel, error)
	GetArticle(id uint32, state uint8) (ArticleModel, error)
	DeleteArticle(id uint32) error
	CountArticleListByTagID(id uint32, state uint8) (int64, error)
	GetArticleListByTagID(id uint32, state uint8) ([]*ArticleTagRow, error)
}

func NewArticleDaoDB(db *gorm.DB) *ArticleDaoDB {
	return &ArticleDaoDB{engine: db}
}

func (d *ArticleDaoDB) CreateArticle(title string, desc string, content string, image string,
	state uint8, createdBy string) (*ArticleModel, error) {
	article := ArticleModel{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageURL: image,
		State:         state,
		Model:         &Model{CreatedBy: createdBy},
	}
	if err := d.engine.Create(&article).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

func (d *ArticleDaoDB) GetArticle(id uint32, state uint8) (ArticleModel, error) {
	var article ArticleModel
	db := d.engine.Where("id = ? AND state = ? AND is_del = 0", id, state)
	if err := db.First(&article).Error; err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}

	return article, nil
}

func (d *ArticleDaoDB) DeleteArticle(id uint32) error {
	article := ArticleModel{Model: &Model{ID: id}}
	return d.engine.Where("id = ? AND is_del = ?", id, 0).Delete(&article).Error
}

func (d *ArticleDaoDB) CountArticleListByTagID(id uint32, state uint8) (int64, error) {
	var count int64
	query := d.engine.Table(ArticleTagModel{}.TableName()+" AS at").
		Distinct("at.article_id").
		Joins("LEFT JOIN `"+TagModel{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+ArticleModel{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("ar.state = ? AND ar.is_del = ?", state, 0)
	if id > 0 {
		// append and clause to query
		query = query.Where("at.`tag_id` = ?", id)
	}
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

type ArticleTagRow struct {
	Article ArticleModel
	Tag     TagModel
}

func newArticleTagRow() *ArticleTagRow {
	return &ArticleTagRow{
		Article: ArticleModel{Model: &Model{}},
		Tag:     TagModel{Model: &Model{}},
	}
}

func (d *ArticleDaoDB) GetArticleListByTagID(id uint32, state uint8) ([]*ArticleTagRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title AS article_title", "ar.desc AS article_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	// TODO: check offset on API layer if necessary
	// pageOffset := app.GetPageOffset(page, pageSize)
	// db := d.engine
	// if pageOffset >= 0 && pageSize > 0 {
	// 	db = db.Offset(pageOffset).Limit(pageSize)
	// }
	query := d.engine.Select(fields).Table(ArticleTagModel{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+TagModel{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+ArticleModel{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("ar.state = ? AND ar.is_del = ?", state, 0)
	if id > 0 {
		// append and clause to query
		query.Where("at.`tag_id` = ?", id)
	}
	query.Order("ar.id ASC")

	rows, err := query.Rows()
	if err != nil || rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleTagRow
	for rows.Next() {
		r := newArticleTagRow()
		if err := rows.Scan(&r.Article.ID, &r.Article.Title, &r.Article.Desc, &r.Article.CoverImageURL,
			&r.Article.Content, &r.Tag.ID, &r.Tag.Name); err != nil {
			return nil, err
		}

		articles = append(articles, r)
	}

	return articles, nil
}
