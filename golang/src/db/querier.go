package db

type Querier interface {
	GetAccount(id string) Account
}
