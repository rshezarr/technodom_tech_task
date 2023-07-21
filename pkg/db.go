package pkg

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context) (*mongo.Database, error) {
	cliOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, cliOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database("url-db"), nil
}

func CreateCollections(ctx context.Context, database *mongo.Database) error {
	collections := []string{"url-collection"}

	for _, collection := range collections {
		if err := database.CreateCollection(context.Background(), collection); err != nil {
			return err
		}
	}

	return nil
}
