package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
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
	// TODO: Implement author_id and is_anonymous snippets
	// TODO: What should it return after insert?
	query := `INSERT INTO public.snippets (title, lang, code) VALUES ($1, $2, $3) RETURNING link;`

	if err := r.client.QueryRow(ctx, query, snippet.Title, snippet.Language, snippet.Code).Scan(&snippet.Link); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			return errors.New(fmt.Sprintf("An error occurred while creating new user: [%s] %s\n", pgErr.SQLState(), pgErr.Message))
		}

		return err
	}

	return nil
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

func (r repository) GetByID(ctx context.Context, id int) (*snippet.Snippet, error) {
	query := `SELECT
    	id, link, title, lang, code, author_id, is_anonymous, created_at, updated_at
		FROM public.snippets
		WHERE id = $1;`

	var s snippet.Snippet

	err := r.client.QueryRow(ctx, query, id).Scan(&s.ID, &s.Link, &s.Title, &s.Language, &s.Code, &s.AuthorID, &s.IsAnonymous, &s.CreatedAt, &s.UpdatedAt)

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
