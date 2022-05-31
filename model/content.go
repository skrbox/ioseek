package model

// 图文标签
type Tag struct {
	Name string `json:"name" gorm:"unique"`
	meta
}

func (t Tag) TableName() string {
	return "tag"
}

// 图文内容
type Post struct {
	Title  string `json:"title"`
	Tags   []Tag  `json:"tags" gorm:"many2many:post_tags;"`
	Author string `json:"author"`
	Origin string `json:"origin"` // 原始地址

	meta
}

func (p Post) TableName() string {
	return "post"
}
