package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// collect every "repository" interface
type Repository struct {
	U UrlInterface
}

// link interface and implementation
func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		U: NewUrlRepo(db, "url-collection"),
	}
}
