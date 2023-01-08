package snippet

import "github.com/jackc/pgtype"

type Snippet struct {
	ID          int                `json:"id"`
	Link        string             `json:"link"`
	Title       string             `json:"title"`
	Language    string             `json:"language"`
	Code        string             `json:"code"`
	AuthorID    int                `json:"author_id"`
	IsAnonymous bool               `json:"is_anonymous"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}
