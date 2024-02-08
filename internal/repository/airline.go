package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"test-smartway/internal/entity"
)

type airlineRepository struct {
	db *pgxpool.Pool
}

func NewAirlineRepository(db *pgxpool.Pool) entity.IAirlineRepository {
	return &airlineRepository{db: db}
}

func (r airlineRepository) InsertAirline(ctx context.Context, airline *entity.Airline) (*entity.Airline, error) {

	_, err := r.db.Exec(ctx, "insert into airlines(code, name) values ($1, $2)", airline.Code, airline.Name)
	if err != nil {
		return nil, err
	}

	return airline, nil
}

func (r airlineRepository) DeleteAirline(ctx context.Context, code string) error {

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `delete from airlines where code=$1`, code)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `delete from airline_provider where airline_id=$1`, code)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r airlineRepository) ReplaceAirlineProviders(ctx context.Context, airlineProviders *entity.AirlineProviders) (*entity.AirlineProviders, error) {

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `delete from airline_provider where airline_id=$1`, airlineProviders.AirlineCode)
	if err != nil {
		return nil, err
	}

	rows := make([][]interface{}, 0, len(airlineProviders.ProvidersId))

	for _, item := range airlineProviders.ProvidersId {
		rows = append(rows, []interface{}{airlineProviders.AirlineCode, item})
	}

	_, err = tx.CopyFrom(ctx, pgx.Identifier{"airline_provider"}, []string{"airline_id", "provider_id"}, pgx.CopyFromRows(rows))
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return airlineProviders, nil
}
