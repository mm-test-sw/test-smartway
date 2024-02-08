package entity

import "context"

type IProviderService interface {
	AddProvider(ctx context.Context, provider *Provider) (*Provider, error)
	DeleteProvider(ctx context.Context, id string) error
	GetAirlines(ctx context.Context, id string) ([]Airline, error)
}

type IProviderRepository interface {
	InsertProvider(ctx context.Context, provider *Provider) (*Provider, error)
	DeleteProvider(ctx context.Context, id string) error
	SelectAirlinesByProvider(ctx context.Context, id string) ([]Airline, error)
}

type Provider struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
