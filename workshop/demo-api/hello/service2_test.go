package hello_test

import (
	"api/hello"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Day int

type SuccessRepo struct{}
func (hr SuccessRepo) Get(id int) (string,error) {
    res := "mock"
	return res, nil
}

func TestGetSuccess(t *testing.T) {
	r := SuccessRepo{}
	s := hello.HelloService{Repo: r}
	result, err := s.GetDataById(1)
	assert.Nil(t, err)
	assert.Equal(t, "mock", result)
}