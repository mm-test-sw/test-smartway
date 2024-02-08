package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"test-smartway/internal/entity"
)

type schemaRepository struct {
	db *pgxpool.Pool
}

func NewSchemaRepository(db *pgxpool.Pool) entity.ISchemaRepository {
	return &schemaRepository{db: db}
}
func (r schemaRepository) GetTx(ctx context.Context) (pgx.Tx, error) {
	return r.db.Begin(ctx)
}

func (r schemaRepository) InsertSchema(ctx context.Context, schema *entity.Schema) (*entity.Schema, error) {

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "insert into schemas(id, name) values ($1, $2)", schema.Id, schema.Name)
	if err != nil {
		return nil, err
	}

	rows := make([][]interface{}, 0, len(schema.Providers))

	for _, item := range schema.Providers {
		rows = append(rows, []interface{}{schema.Id, item})
	}

	_, err = tx.CopyFrom(ctx, pgx.Identifier{"schema_provider"}, []string{"schema_id", "provider_id"}, pgx.CopyFromRows(rows))
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

func (r schemaRepository) SelectSchemaByName(ctx context.Context, name string) (*entity.Schema, error) {

	rows, err := r.db.Query(ctx, `select id, name, sp.provider_id from schemas 
    left join schema_provider as sp on schema_id = id
    where name=$1`, name)
	if err != nil {
		return nil, err
	}

	var schema entity.Schema
	var provider string
	for rows.Next() {
		rows.Scan(
			&schema.Id,
			&schema.Name,
			&provider,
		)
		if len(provider) != 0 {
			schema.Providers = append(schema.Providers, provider)
		}
	}

	return &schema, nil
}

func (r schemaRepository) TxUpdateSchemaName(ctx context.Context, tx pgx.Tx, id int, name string) error {

	_, err := tx.Exec(ctx, `update schemas set name = $1 where id = $2`, name, id)
	if err != nil {
		return err
	}

	return nil
}

func (r schemaRepository) TxReplaceSchemaProviders(ctx context.Context, tx pgx.Tx, id int, providers []string) error {

	_, err := tx.Exec(ctx, `delete from schema_provider where schema_id=$1`, id)
	if err != nil {
		return err
	}

	rows := make([][]interface{}, 0, len(providers))

	for _, item := range providers {
		rows = append(rows, []interface{}{id, item})
	}

	_, err = tx.CopyFrom(ctx, pgx.Identifier{"schema_provider"}, []string{"schema_id", "provider_id"}, pgx.CopyFromRows(rows))
	if err != nil {
		return err
	}

	return nil
}

func (r schemaRepository) DeleteSchema(ctx context.Context, id string) error {

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `delete from schemas where id = $1`, id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `delete from schema_provider where schema_id = $1`, id)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r schemaRepository) IsSchemeAssignedToAccount(ctx context.Context, id string) (bool, error) {

	rows, err := r.db.Query(ctx, `select * from accounts where schema_id = $1`, id)
	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}

func (r schemaRepository) TxSelectSchema(ctx context.Context, tx pgx.Tx, id int) (*entity.Schema, error) {

	rows, err := tx.Query(ctx, `select id, name, sp.provider_id from schemas 
    left join schema_provider as sp on schema_id = id
    where id=$1`, id)
	if err != nil {
		return nil, err
	}

	var schema entity.Schema
	var provider string
	for rows.Next() {
		rows.Scan(
			&schema.Id,
			&schema.Name,
			&provider,
		)
		if len(provider) != 0 {
			schema.Providers = append(schema.Providers, provider)
		}
	}

	return &schema, nil
}
