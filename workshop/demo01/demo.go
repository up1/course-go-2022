package demo

import "fmt"

type Account struct {
	name string
	pass string
}

type Account2 struct {
	Account
	age int
}

func xx(s IService) {
	
}

func yy() {
	a := AccountService{}
	xx(a)
}

// === Service
type IService interface {
	CreateAccount(username string, password string) (*Account, error)
}

type AccountService struct {
}

func NewAccountService() AccountService {
	return AccountService{}
}

func (a AccountService) CreateAccount(username string, password string) (*Account, error) {
	if true {
		return nil, fmt.Errorf("Error")
	}
	return &Account{name: username, pass: password}, nil
}
