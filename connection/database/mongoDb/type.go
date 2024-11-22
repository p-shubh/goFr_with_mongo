package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type Database interface {
	GetMongoConnecttion() *mongo.Client
}
