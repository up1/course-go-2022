package user

import (
	"context"
	"demo/user/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountService struct {
	repo AccountRepository
}

type AccountRepository interface {
	AddAccount(ctx context.Context, account *repository.Account) (*repository.Account, error)
}

func NewAccountService(repo AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (svc *AccountService) AddAccount(ctx context.Context, req *CreateAccount) (*repository.Account, error) {
	newAccount := repository.Account{
		ID: primitive.NewObjectID(),
		Firstname: req.Firstname,
		Lastname: req.Lastname,
		Password: req.Password,
		Email: req.Email,
	}
	account, err := svc.repo.AddAccount(ctx, &newAccount)
	if err != nil {
		return nil, err
	}
	return account, nil
}
