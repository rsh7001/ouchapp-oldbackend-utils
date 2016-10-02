package ouchapp

type Page struct {
	ID         string `json:"id"`
	EmbeddedID string `json:"EmbeddedId"`
	ArticleID  string `json:"ArticleId"`
	LinkTitle  string
	LinkIDs    string `json:"LinkIds"`
}
