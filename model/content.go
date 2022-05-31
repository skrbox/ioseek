package model

type Tags struct {
	meta
	Name string `json:"name" yaml:"name"`
}

type Post struct {
	meta
	Title  string `json:"title" yaml:"title"`
	Tags   []Tags `json:"tags" gorm:"many2many:post_tags;"`
	Author string `json:"author" `
}
