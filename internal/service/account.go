package service

import (
	"context"
	"test-smartway/internal/app/config"
	"test-smartway/internal/entity"
)

type accountService struct {
	accountRepo entity.IAccountRepository
	cfg         *config.Config
}

func NewAccountService(cfg *config.Config, accountRepo entity.IAccountRepository) entity.IAccountService {
	return &accountService{accountRepo: accountRepo, cfg: cfg}
}

func (s *accountService) AddAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	return s.accountRepo.InsertAccount(ctx, account)
}

func (s *accountService) DeleteAccount(ctx context.Context, code string) error {
	return s.accountRepo.DeleteAccount(ctx, code)
}

func (s *accountService) UpdateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	return s.accountRepo.UpdateAccount(ctx, account)
}

func (s *accountService) GetAirlines(ctx context.Context, id string) ([]entity.Airline, error) {
	if id == "1" {
		return s.cfg.DemoAccountAirlines, nil
	}

	return s.accountRepo.SelectAirlinesByAccount(ctx, id)
}
