package demo_test

import (
	"demo"
	"testing"
)

func TestFaliCase(t *testing.T) {
	acc := demo.NewAccountService()
	_, err := acc.CreateAccount("", "")
	if err == nil {
		t.Fatal();
	}
}