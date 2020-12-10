package entity

//Article represents article data from json
type Article struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	PublishedAt string `json:"published_at"`
}
