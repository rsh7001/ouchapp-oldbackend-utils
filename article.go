package ouchapp

type Article struct {
	ID          string `json:"id"`
	EmbeddedID  string `json:"EmbeddedId"`
	ArticleType string
	Title       string
	Summary     string
	HTMLContent string `json:"HtmlContent"`
}
