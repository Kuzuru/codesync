package db

import (
	"context"

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
	//TODO implement me
	panic("implement me")
}

func (r repository) UpdateSnippet(ctx context.Context, snippet snippet.Snippet) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) DeleteByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
