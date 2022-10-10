package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CreateAccount struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func CreateAccountHandler(service *AccountService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req *CreateAccount
		if err := c.Bind(&req); err != nil {
			logrus.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		account, err := service.AddAccount(c.Request().Context(), req)
		if err != nil {
			logrus.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, account)
	}
}
