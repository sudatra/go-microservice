package account

import (
	"context"
	"github.com/segmentio/ksuid"
)

type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type Account struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type accountService struct {
	repository Repository
}

func NewAccountService(repository Repository) *accountService {
	return &accountService{repository}
}

func (as *accountService) PostAccount(ctx context.Context, name string) (*Account, error) {
	a := &Account{
		ID: ksuid.New().String(),
		Name: name,
	}

	if err := as.repository.PutAccount(ctx, *a); err != nil {
		return nil, err
	}

	return a, nil;
}

func (as *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	return as.repository.GetAccountByID(ctx, id);
}

func (as *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}

	return as.repository.ListAccounts(ctx, skip, take);
}