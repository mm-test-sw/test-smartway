package entity

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type ISchemaService interface {
	AddSchema(ctx context.Context, schema *Schema) (*Schema, error)
	GetSchema(ctx context.Context, name string) (*Schema, error)
	PatchSchema(ctx context.Context, schema *Schema) (*Schema, error)
	DeleteSchema(ctx context.Context, id string) error
}

type ISchemaRepository interface {
	GetTx(ctx context.Context) (pgx.Tx, error)
	InsertSchema(ctx context.Context, schema *Schema) (*Schema, error)
	TxSelectSchema(ctx context.Context, tx pgx.Tx, id int) (*Schema, error)
	SelectSchemaByName(ctx context.Context, name string) (*Schema, error)
	TxUpdateSchemaName(ctx context.Context, tx pgx.Tx, id int, name string) error
	TxReplaceSchemaProviders(ctx context.Context, tx pgx.Tx, id int, providers []string) error
	DeleteSchema(ctx context.Context, id string) error
	IsSchemeAssignedToAccount(ctx context.Context, id string) (bool, error)
	CheckSchema(ctx context.Context, id string) (bool, error)
}

type Schema struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Providers []string `json:"providers"`
}
