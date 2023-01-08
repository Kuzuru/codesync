package snippet

type Snippet struct {
	ID          int    `json:"id"`
	Link        string `json:"link"`
	Title       string `json:"title"`
	Language    string `json:"language"`
	Code        string `json:"code"`
	AuthorID    int    `json:"author_id"`
	IsAnonymous bool   `json:"is_anonymous"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
