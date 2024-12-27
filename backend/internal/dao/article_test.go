package dao

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Demonstrate how to use sqlmock for DAO layer to do unit tests

func newTestArticleDao() (*gorm.DB, sqlmock.Sqlmock) {
	mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, mock
}

func TestArticleDao_CreateArticle(t *testing.T) {
	type fields struct {
		engine *gorm.DB
		mock   sqlmock.Sqlmock
	}
	type args struct {
		title     string
		desc      string
		content   string
		image     string
		state     uint8
		createdBy string
	}
	db, mock := newTestArticleDao()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ArticleModel
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: fields{
				engine: db,
				mock:   mock,
			},
			args: args{
				title:     "title",
				desc:      "desc",
				content:   "content",
				image:     "image",
				state:     0,
				createdBy: "author",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDaoDB{
				engine: tt.fields.engine,
			}

			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO `blog_article` (`created_by`,`modified_by`,`is_del`,`title`,`desc`,`content`,`cover_image_url`,`state`) VALUES (?,?,?,?,?,?,?,?)").
				WithArgs(tt.args.createdBy, "", 0, tt.args.title, tt.args.desc, tt.args.content, tt.args.image, tt.args.state).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			got, err := d.CreateArticle(tt.args.title, tt.args.desc, tt.args.content, tt.args.image, tt.args.state, tt.args.createdBy)

			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleDaoDB.CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil || ((got.Title != tt.args.title) || (got.Desc != tt.args.desc) ||
				(got.Content != tt.args.content) || (got.CoverImageURL != tt.args.image) || (got.State != tt.args.state)) {
				t.Errorf("ArticleDaoDB.CreateArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleDao_GetArticle(t *testing.T) {
	type fields struct {
		engine *gorm.DB
		mock   sqlmock.Sqlmock
	}
	type args struct {
		id    uint32
		state uint8
	}
	db, mock := newTestArticleDao()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ArticleModel
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: fields{
				engine: db,
				mock:   mock,
			},
			args: args{
				id:    1024,
				state: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDaoDB{
				engine: tt.fields.engine,
			}
			article := &ArticleModel{
				Title:         "title",
				Desc:          "desc",
				Content:       "content",
				CoverImageURL: "image",
				State:         tt.args.state,
				Model: &Model{
					ID: tt.args.id,
				},
			}
			mock.ExpectQuery("SELECT * FROM `blog_article` WHERE id = ? AND state = ? AND is_del = 0 ORDER BY `blog_article`.`id` LIMIT 1").
				WithArgs(tt.args.id, tt.args.state).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "title", "desc", "cover_image_url", "content", "created_by", "modified_by", "is_del", "state"}).
						AddRow("1024", article.Title, article.Desc, article.CoverImageURL, article.Content, article.CreatedBy, article.ModifiedBy, "0", article.State),
				)

			got, err := d.GetArticle(tt.args.id, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleDaoDB.GetArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got.Title != article.Title) || (got.Desc != article.Desc) || (got.CoverImageURL != article.CoverImageURL) ||
				(got.Content != article.Content) || (got.State != article.State) || (got.ID != article.ID) {
				t.Errorf("ArticleDaoDB.GetArticle() = %v, want %v", got, article)
			}
		})
	}
}

func TestArticleDao_DeleteArticle(t *testing.T) {
	type fields struct {
		engine *gorm.DB
	}
	type args struct {
		id uint32
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
			d := &ArticleDaoDB{
				engine: tt.fields.engine,
			}
			if err := d.DeleteArticle(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ArticleDaoDB.DeleteArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleDao_CountArticleListByTagID(t *testing.T) {
	type fields struct {
		engine *gorm.DB
	}
	type args struct {
		id    uint32
		state uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDaoDB{
				engine: tt.fields.engine,
			}
			got, err := d.CountArticleListByTagID(tt.args.id, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleDaoDB.CountArticleListByTagID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ArticleDaoDB.CountArticleListByTagID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newArticleTagRow(t *testing.T) {
	tests := []struct {
		name string
		want *ArticleTagRow
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newArticleTagRow(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newArticleTagRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleDao_GetArticleListByTagID(t *testing.T) {
	type fields struct {
		engine *gorm.DB
	}
	type args struct {
		id    uint32
		state uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ArticleTagRow
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDaoDB{
				engine: tt.fields.engine,
			}
			got, err := d.GetArticleListByTagID(tt.args.id, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleDaoDB.GetArticleListByTagID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleDaoDB.GetArticleListByTagID() = %v, want %v", got, tt.want)
			}
		})
	}
}
