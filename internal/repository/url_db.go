package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlInterface interface {
}

type UrlRepo struct {
	db *mongo.Collection
}

func NewUrlRepo(database *mongo.Database, collection string) *UrlRepo {
	return &UrlRepo{
		db: database.Collection(collection),
	}
}
