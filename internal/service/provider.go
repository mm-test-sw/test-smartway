package service

import (
	"context"
	"test-smartway/internal/entity"
)

type providerService struct {
	providerRepo entity.IProviderRepository
}

func NewProviderService(providerRepo entity.IProviderRepository) entity.IProviderService {
	return &providerService{providerRepo: providerRepo}
}

func (s *providerService) AddProvider(ctx context.Context, provider *entity.Provider) (*entity.Provider, error) {
	return s.providerRepo.InsertProvider(ctx, provider)
}

func (s *providerService) DeleteProvider(ctx context.Context, id string) error {
	return s.providerRepo.DeleteProvider(ctx, id)
}

func (s *providerService) GetAirlines(ctx context.Context, id string) ([]entity.Airline, error) {
	return s.providerRepo.SelectAirlinesByProvider(ctx, id)
}
