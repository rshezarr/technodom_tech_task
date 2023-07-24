package pkg

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context) (*mongo.Database, error) {
	//set connection URI
	cliOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//connection
	client, err := mongo.Connect(ctx, cliOptions)
	if err != nil {
		return nil, err
	}

	//ping to check the connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	//return and init database by setting name
	return client.Database("url-db"), nil
}

func CreateCollections(ctx context.Context, database *mongo.Database) error {
	//slice with all collection names
	collections := []string{"url-collection"}

	//creating every collection in mongo
	for _, collection := range collections {
		if err := database.CreateCollection(context.Background(), collection); err != nil {
			return err
		}
	}

	return nil
}
