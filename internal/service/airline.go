package service

import (
	"context"
	"test-smartway/internal/entity"
)

type airlineService struct {
	airlineRepo entity.IAirlineRepository
}

func NewAirlineService(airlineRepo entity.IAirlineRepository) entity.IAirlineService {
	return &airlineService{airlineRepo: airlineRepo}
}

func (s *airlineService) AddAirline(ctx context.Context, airline *entity.Airline) (*entity.Airline, error) {
	return s.airlineRepo.InsertAirline(ctx, airline)
}

func (s *airlineService) DeleteAirline(ctx context.Context, code string) error {
	return s.airlineRepo.DeleteAirline(ctx, code)
}

func (s *airlineService) PutAirlineProviders(ctx context.Context, airlineProviders *entity.AirlineProviders) (*entity.AirlineProviders, error) {

	ok, err := s.airlineRepo.CheckAirline(ctx, airlineProviders.AirlineCode)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, entity.NewLogicError(nil, "airline not exist", 400)
	}

	return s.airlineRepo.ReplaceAirlineProviders(ctx, airlineProviders)
}
