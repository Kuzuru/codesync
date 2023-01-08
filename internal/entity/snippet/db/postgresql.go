package db

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/kuzuru/codesync/internal/entity/snippet"
	"github.com/kuzuru/codesync/pkg/db/postgresql"
)

type repository struct {
	client postgresql.Client
}

func NewRepository(client postgresql.Client) snippet.Storage {
	return &repository{
		client: client,
	}
}

func (r repository) Create(ctx context.Context, snippet *snippet.Snippet) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetByLink(ctx context.Context, link string) (*snippet.Snippet, error) {
	query := `SELECT
    	id, link, title, lang, code, author_id, is_anonymous, created_at, updated_at
		FROM public.snippets
		WHERE link = $1;`

	var s snippet.Snippet

	err := r.client.QueryRow(ctx, query, link).Scan(&s.ID, &s.Link, &s.Title, &s.Language, &s.Code, &s.AuthorID, &s.IsAnonymous, &s.CreatedAt, &s.UpdatedAt)

	// Check if no rows was found with ErrNoRows
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("nothing was found :(")
		}

		return nil, err
	}

	return &s, nil
}

func (r repository) UpdateSnippet(ctx context.Context, snippet snippet.Snippet) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) DeleteByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
