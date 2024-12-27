package domain

import (
	"context"

	"github.com/alex-guoba/gin-clean-template/internal/dao"
	"github.com/alex-guoba/gin-clean-template/internal/entity"

	"gorm.io/gorm"
)

type ArticleDomain struct {
	ctx          context.Context
	articleDao   dao.ArticleDao
	tagDao       dao.TagDao
	artileTagDao dao.ArticleTagDao
}

func NewArticleDomain(ctx context.Context, db *gorm.DB) *ArticleDomain {
	d := &ArticleDomain{ctx: ctx}
	d.tagDao = dao.NewTagDao(db)
	d.articleDao = dao.NewArticleDaoDB(db)
	d.artileTagDao = dao.NewArticleTagDao(db)
	return d
}

func (d *ArticleDomain) GetArticle(id uint32, state uint8) (*entity.ArticleEntity, error) {
	// Query article info
	article, err := d.articleDao.GetArticle(id, state)
	if err != nil {
		return nil, err
	}

	// Query state id
	articleTags, err := d.artileTagDao.GetArticleTagsByAID(article.ID)
	if err != nil {
		return nil, err
	}

	tids := make([]uint32, len(articleTags))
	for i, v := range articleTags {
		tids[i] = v.TagID
	}

	// Query tag info
	tags, err := d.tagDao.GetTags(tids, dao.StateOpen)
	if err != nil {
		return nil, err
	}

	// TODO: convert to entry object
	en := &entity.ArticleEntity{
		ID:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageURL: article.CoverImageURL,
		State:         article.State,
		Tags:          []entity.TagEntity{},
	}

	for _, tag := range tags {
		en.Tags = append(en.Tags, entity.TagEntity{
			ID:        tag.ID,
			Name:      tag.Name,
			State:     tag.State,
			CreatedBy: tag.CreatedBy,
		})
	}
	return en, nil
}

func (d *ArticleDomain) countArticleByTagID(id uint32) (int, error) {
	cnt, err := d.articleDao.CountArticleListByTagID(id, 1)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (d *ArticleDomain) GetArticleList(id uint32, state uint8, page, pageSize int) ([]*entity.ArticleEntity, int, error) {
	// Query article count
	cnt, err := d.articleDao.CountArticleListByTagID(id, state)
	if err != nil {
		return nil, 0, err
	}

	// Query article list. Make sure it ordered by article id
	artileTags, err := d.articleDao.GetArticleListByTagID(id, state)
	if err != nil {
		return nil, 0, err
	}

	var articleList []*entity.ArticleEntity
	var last *entity.ArticleEntity = nil
	for _, row := range artileTags {
		if last != nil && last.ID == row.Article.ID {
			last.Tags = append(last.Tags, entity.TagEntity{
				ID:   row.Tag.ID,
				Name: row.Tag.Name,
			})
			continue
		}

		last = &entity.ArticleEntity{
			ID:            row.Article.ID,
			Title:         row.Article.Title,
			Desc:          row.Article.Desc,
			Content:       row.Article.Content,
			CoverImageURL: row.Article.CoverImageURL,
			State:         row.Article.State,
			Tags: []entity.TagEntity{{
				ID:   row.Tag.ID,
				Name: row.Tag.Name,
			}},
		}
		articleList = append(articleList, last)
	}

	// deal with pagination
	if page > 0 && pageSize > 0 {
		start := (page - 1) * pageSize
		if start > len(articleList) {
			return nil, 0, nil
		}

		end := start + pageSize
		if end > len(articleList) {
			end = len(articleList)
		}
		articleList = articleList[start:end]
	}

	return articleList, int(cnt), nil
}

func (d *ArticleDomain) CreateArticle(title string, desc string, content string, image string,
	state uint8, createdBy string, tagIDs []uint32) error {
	// Insert article
	article, err := d.articleDao.CreateArticle(title, desc, content, image, state, createdBy)
	if err != nil {
		return err
	}

	// Insert article tag relation
	for _, tagID := range tagIDs {
		if err = d.artileTagDao.CreateArticleTag(article.ID, tagID, createdBy); err != nil {
			return err
		}
	}
	return nil
}
