package dao

const (
	StateOpen  = 1
	StateClose = 0
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	IsDel      uint8  `json:"is_del"`
}

// Model to tag.
type TagModel struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (TagModel) TableName() string {
	return "blog_tag"
}

// Model to Artile.
type ArticleModel struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageURL string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (ArticleModel) TableName() string {
	return "blog_article"
}

// Model to Artile-Tag.
type ArticleTagModel struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (ArticleTagModel) TableName() string {
	return "blog_article_tag"
}
