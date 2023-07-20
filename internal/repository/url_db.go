package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlInterface interface {
	Create(ctx context.Context) error
	Get(ctx context.Context, email string) error
}

type UrlRepo struct {
	db *mongo.Collection
}

func NewUrlRepo(database *mongo.Database, collection string) *UrlRepo {
	database.CreateCollection(context.Background(), collection)
	return &UrlRepo{
		db: database.Collection(collection),
	}
}

func (u UrlRepo) Create(ctx context.Context) error {
	return nil
}

func (u UrlRepo) Get(ctx context.Context, email string) error {
	return nil
}
