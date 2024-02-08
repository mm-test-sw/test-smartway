package entity

import "context"

type IAccountService interface {
	AddAccount(ctx context.Context, account *Account) (*Account, error)
	DeleteAccount(ctx context.Context, id string) error
	UpdateAccount(ctx context.Context, account *Account) (*Account, error)
	GetAirlines(ctx context.Context, id string) ([]Airline, error)
}

type IAccountRepository interface {
	InsertAccount(ctx context.Context, account *Account) (*Account, error)
	DeleteAccount(ctx context.Context, id string) error
	UpdateAccount(ctx context.Context, account *Account) (*Account, error)
	SelectAirlinesByAccount(ctx context.Context, id string) ([]Airline, error)
}

type Account struct {
	Id       int `json:"id"`
	SchemaId int `json:"schemaId"`
}
