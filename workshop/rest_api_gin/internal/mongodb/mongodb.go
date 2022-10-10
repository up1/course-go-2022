package mongodb

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDB(connectionString string) *mongo.Client {
	if strings.TrimSpace(connectionString) == "" {
		logrus.Error("missing connection string")
		os.Exit(1)
		return nil
	}
	clientOptions := options.Client().ApplyURI(connectionString)
	clientOptions.SetSocketTimeout(20 * time.Second)
	clientOptions.SetServerSelectionTimeout(time.Second)
	clientOptions.SetConnectTimeout(20 * time.Second)
	clientOptions.SetMaxConnIdleTime(20 * time.Second)
	clientOptions.SetMinPoolSize(100)
	clientOptions.SetMaxPoolSize(100)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logrus.Error("failed to connect database: ", err)
		os.Exit(1)
		return nil
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logrus.Error("failed ping to database: ", err)
		os.Exit(1)
		return nil
	}
	return client
}
