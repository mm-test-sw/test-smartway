package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"test-smartway/internal/entity"
)

type accountRepository struct {
	db *pgxpool.Pool
}

func NewAccountRepository(db *pgxpool.Pool) entity.IAccountRepository {
	return &accountRepository{db: db}
}

func (r accountRepository) InsertAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {

	_, err := r.db.Exec(ctx, "insert into accounts(id, schema_id) values ($1, $2)", account.Id, account.SchemaId)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r accountRepository) DeleteAccount(ctx context.Context, id string) error {

	_, err := r.db.Exec(ctx, `delete from accounts where id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r accountRepository) UpdateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	_, err := r.db.Exec(ctx, `update accounts set schema_id = $1 where id = $2`, account.SchemaId, account.Id)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r accountRepository) SelectAirlinesByAccount(ctx context.Context, id string) ([]entity.Airline, error) {

	rows, err := r.db.Query(ctx, `select airlines.code, airlines.name from accounts
    left join schema_provider as sp on sp.schema_id = accounts.schema_id
    left join airline_provider as ap on ap.provider_id = sp.provider_id
    left join airlines on airlines.code = ap.airline_id
    where accounts.id=$1
    group by airlines.code`, id)
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
