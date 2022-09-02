package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CreateAccount struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func CreateAccountHandler(service *AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *CreateAccount
		if err := c.ShouldBindJSON(&req); err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		account, err := service.AddAccount(c.Request.Context(), req)
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, account)
	}
}
