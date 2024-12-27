package entity

type TagEntity struct {
	ID         uint32 `json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	IsDel      uint8  `json:"is_del"`
	Name       string `json:"name"`
	State      uint8  `json:"state"`
}

type ArticleEntity struct {
	ID            uint32      `json:"id"`
	Title         string      `json:"title"`
	Desc          string      `json:"desc"`
	Content       string      `json:"content"`
	CoverImageURL string      `json:"cover_image_url"`
	State         uint8       `json:"state"`
	Tags          []TagEntity `json:"tags"`
}
