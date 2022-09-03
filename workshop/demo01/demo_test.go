package demo_test

import (
	"demo"
	"testing"
)

func TestXxx(t *testing.T) {
	acc := demo.NewAccountService()
	acc.CreateAccount("", "")
}