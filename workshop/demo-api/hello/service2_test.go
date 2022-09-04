package hello_test

import (
	"api/hello"
	"fmt"
	"testing"
)


type SuccessRepo struct{}
func (hr SuccessRepo) Get(id int) (string,error) {
    res := "mock"
	return res, nil
}

func TestGetSuccess(t *testing.T) {
	r := SuccessRepo{}
	s := hello.HelloService{Repo: r}
	result, err := s.GetDataById(1)
	fmt.Println(result, err)
}