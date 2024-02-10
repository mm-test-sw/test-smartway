package entity

import "context"

type IAirlineService interface {
	AddAirline(ctx context.Context, airline *Airline) (*Airline, error)
	DeleteAirline(ctx context.Context, code string) error
	PutAirlineProviders(ctx context.Context, airlineProviders *AirlineProviders) (*AirlineProviders, error)
}

type IAirlineRepository interface {
	InsertAirline(ctx context.Context, airline *Airline) (*Airline, error)
	DeleteAirline(ctx context.Context, code string) error
	ReplaceAirlineProviders(ctx context.Context, airlineProviders *AirlineProviders) (*AirlineProviders, error)
	CheckAirline(ctx context.Context, code string) (bool, error)
}

type Airline struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type AirlineProviders struct {
	AirlineCode string   `json:"code"`
	ProvidersId []string `json:"providersId"`
}
