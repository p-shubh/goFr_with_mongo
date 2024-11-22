package mongodb

import (
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main.go/config"
)

type MongoDatabase struct {
	connect *mongo.Client
}

var (
	syncs      sync.Once
	dbInstance *MongoDatabase
)

func ConnectDatabase(config *config.Enviroment) Database {

	syncs.Do(func() {
		uri := fmt.Sprintf("mongodb://%s:%s@%s/%s", config.DATABASE_USERNAME, config.DATABASE_PASSWORD, config.DATABASE_HOST, config.DATABASE_NAME)
		fmt.Println("URI : ", uri)

		db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			logrus.Errorf("Failed to connect :%s", err)
			return
		} else {
			logrus.Infof("Connected to MongoDB")
		}
		dbInstance = &MongoDatabase{connect: db}
	})
	return dbInstance
}

func (m *MongoDatabase) GetMongoConnecttion() *mongo.Client {
	return m.connect
}
