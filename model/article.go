package model

type Article struct {
	meta
	Title string `json:"title" yaml:"title"`
}
