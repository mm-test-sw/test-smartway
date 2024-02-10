package service

import (
	"context"
	"strconv"
	"test-smartway/internal/app/config"
	"test-smartway/internal/entity"
)

type accountService struct {
	accountRepo entity.IAccountRepository
	schemaRepo  entity.ISchemaRepository
	cfg         *config.Config
}

func NewAccountService(cfg *config.Config, accountRepo entity.IAccountRepository, schemaRepo entity.ISchemaRepository) entity.IAccountService {
	return &accountService{accountRepo: accountRepo, schemaRepo: schemaRepo, cfg: cfg}
}

func (s *accountService) AddAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {

	ok, err := s.schemaRepo.CheckSchema(ctx, strconv.Itoa(account.SchemaId))
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, entity.NewLogicError(nil, "schema not exist", 400)
	}

	return s.accountRepo.InsertAccount(ctx, account)
}

func (s *accountService) DeleteAccount(ctx context.Context, code string) error {
	return s.accountRepo.DeleteAccount(ctx, code)
}

func (s *accountService) PutAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {

	ok, err := s.schemaRepo.CheckSchema(ctx, strconv.Itoa(account.SchemaId))
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, entity.NewLogicError(nil, "schema not exist", 400)
	}

	return s.accountRepo.UpdateAccount(ctx, account)
}

func (s *accountService) GetAirlines(ctx context.Context, id string) ([]entity.Airline, error) {
	if id == "1" {
		return s.cfg.DemoAccountAirlines, nil
	}

	ok, err := s.accountRepo.CheckAccount(ctx, id)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, entity.NewLogicError(nil, "account not exist", 400)
	}

	return s.accountRepo.SelectAirlinesByAccount(ctx, id)
}
