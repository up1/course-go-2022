package test

import (
	"context"
	"demo/internal/mongodb"
	"demo/internal/web"
	"demo/user"
	"demo/user/repository"
	"flag"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	httpAddr   = flag.String("http", ":8000", "Http address")
	mongoDbUrl = flag.String("mongodb", "mongodb://mongoadmin:secret@localhost:27017", "MongoDb address")
)

func NewTestRouter() *echo.Echo {
	flag.Parse()
	router := web.NewRouter()

	client := mongodb.MongoDB(*mongoDbUrl)
	repository := repository.NewMongoDb(client)
	accountService := user.NewAccountService(repository)

	router.POST("users", user.CreateAccountHandler(accountService))

	return router
}

type MockAccountRepository struct {}
func (mas MockAccountRepository) AddAccount(ctx context.Context, account *repository.Account) (*repository.Account, error){
	newAccount := repository.Account{
		ID: primitive.NewObjectID(),
		Firstname: "Mock firstname",
		Lastname: "Mock lastname",
		Password: "Mock password",
		Email: "Mock email",
	}
	return &newAccount, nil
}

func NewTestRouterWithMock() *echo.Echo {
	flag.Parse()
	router := web.NewRouter()
	mockRepository := MockAccountRepository{}
	accountService := user.NewAccountService(mockRepository)

	router.POST("users", user.CreateAccountHandler(accountService))

	return router
}