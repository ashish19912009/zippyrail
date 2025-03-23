package account

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostAccount(ctx context.Context, mobileNo uint64) (*Account, error)
	UpdateAccount(ctx context.Context, mobileNo uint64, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]*Account, error)
}

type Account struct {
	ID       string `json:"id"`
	MobileNo string `json:"mobileNo"`
	Name     string `json:"name"`
}

type accountService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &accountService{r}
}

func (s *accountService) PostAccount(ctx context.Context, mobileNo uint64) (*Account, error) {
	a := &Account{
		MobileNo: mobileNo,
		ID:       ksuid.New().String(),
	}
	if err := s.repository.PutAccount(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *accountService) UpdateAccount(ctx context.Context, mobileNo uint64, name string) (*Account, error) {
	a := &Account{
		Name:     name,
		MobileNo: mobileNo,
	}
	if err := s.repository.UpdateAccount(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	account, err := s.repository.GetAccountByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {

	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	accounts, err := s.repository.ListAccounts(ctx, skip, take)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
