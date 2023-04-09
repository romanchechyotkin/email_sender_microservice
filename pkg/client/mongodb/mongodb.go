package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, username, password, database string) (db *mongo.Database, err error) {
	var mongoDBURL string

	if username == "" && password == "" {
		mongoDBURL = fmt.Sprintf("mongodb+srv://%s:%s@cluster0.xygrgb0.mongodb.net/db", username, password)
	} else {
		mongoDBURL = fmt.Sprintf("mongodb+srv://%s:%s@cluster0.xygrgb0.mongodb.net/db", username, password)
	}

	clientOpts := options.Client().ApplyURI(mongoDBURL)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb due to error: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongodb due to error: %v", err)
	}

	return client.Database(database), nil

}
