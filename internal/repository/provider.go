package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"test-smartway/internal/entity"
)

type providerRepository struct {
	db *pgxpool.Pool
}

func NewProviderRepository(db *pgxpool.Pool) entity.IProviderRepository {
	return &providerRepository{db: db}
}

func (r providerRepository) InsertProvider(ctx context.Context, provider *entity.Provider) (*entity.Provider, error) {

	_, err := r.db.Exec(ctx, "insert into providers(id, name) values ($1, $2)", provider.Id, provider.Name)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (r providerRepository) DeleteProvider(ctx context.Context, id string) error {

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `delete from providers where id=$1`, id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `delete from airline_provider where provider_id=$1`, id)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r providerRepository) SelectAirlinesByProvider(ctx context.Context, id string) ([]entity.Airline, error) {

	rows, err := r.db.Query(ctx, `select airlines.code, airlines.name from airline_provider as ap
    left join airlines on airlines.code = ap.airline_id
    where ap.provider_id=$1`, id)
	if err != nil {
		return nil, err
	}

	var airlines []entity.Airline
	var airline entity.Airline
	for rows.Next() {
		rows.Scan(
			&airline.Code,
			&airline.Name,
		)
		airlines = append(airlines, airline)
	}

	return airlines, nil
}
