package hello

import "github.com/go-redis/redis"

type HelloService struct {
	Repo IRepository
}

func (hs HelloService) GetDataById(id int) (string, error) {
    return hs.Repo.Get(id)
}

type IRepository interface {
	Get(id int) (string,error)
}

type HelloRepository struct{
	R *redis.Client
}

func (hr HelloRepository) Get(id int) (string,error) {
    res := "result"
	return res, nil
}