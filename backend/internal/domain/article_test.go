package domain

import (
	"context"
	"reflect"
	"testing"

	"github.com/alex-guoba/gin-clean-template/internal/dao"
	"github.com/alex-guoba/gin-clean-template/internal/entity"
)

func TestArticleDomain_countArticleByTagID(t *testing.T) {
	type fields struct {
		ctx          context.Context
		articleDao   dao.ArticleDao
		tagDao       dao.TagDao
		artileTagDao dao.ArticleTagDao
	}
	type args struct {
		id uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDomain{
				ctx:          tt.fields.ctx,
				articleDao:   tt.fields.articleDao,
				tagDao:       tt.fields.tagDao,
				artileTagDao: tt.fields.artileTagDao,
			}
			got, err := d.countArticleByTagID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleDomain.countArticleByTagID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ArticleDomain.countArticleByTagID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleDomain_GetArticleList(t *testing.T) {
	type fields struct {
		ctx          context.Context
		articleDao   dao.ArticleDao
		tagDao       dao.TagDao
		artileTagDao dao.ArticleTagDao
	}
	type args struct {
		id       uint32
		state    uint8
		page     int
		pageSize int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entity.ArticleEntity
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDomain{
				ctx:          tt.fields.ctx,
				articleDao:   tt.fields.articleDao,
				tagDao:       tt.fields.tagDao,
				artileTagDao: tt.fields.artileTagDao,
			}
			got, got1, err := d.GetArticleList(tt.args.id, tt.args.state, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleDomain.GetArticleList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleDomain.GetArticleList() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ArticleDomain.GetArticleList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestArticleDomain_CreateArticle(t *testing.T) {
	type fields struct {
		ctx          context.Context
		articleDao   dao.ArticleDao
		tagDao       dao.TagDao
		artileTagDao dao.ArticleTagDao
	}
	type args struct {
		title     string
		desc      string
		content   string
		image     string
		state     uint8
		createdBy string
		tagID     uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDomain{
				ctx:          tt.fields.ctx,
				articleDao:   tt.fields.articleDao,
				tagDao:       tt.fields.tagDao,
				artileTagDao: tt.fields.artileTagDao,
			}
			if err := d.CreateArticle(tt.args.title, tt.args.desc, tt.args.content, tt.args.image, tt.args.state, tt.args.createdBy, []uint32{tt.args.tagID}); (err != nil) != tt.wantErr {
				t.Errorf("ArticleDomain.CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
