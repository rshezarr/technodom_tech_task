package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	U UrlInterface
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		U: NewUrlRepo(db, "url_collection"),
	}
}
