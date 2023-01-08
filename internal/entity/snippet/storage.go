package snippet

import "context"

type Storage interface {
	Create(ctx context.Context, snippet *Snippet) error
	GetByLink(ctx context.Context, link string) (*Snippet, error)
	UpdateSnippet(ctx context.Context, snippet Snippet) error
	DeleteByID(ctx context.Context, id string) error
}
