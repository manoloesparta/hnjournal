package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection interacts with the database
type Connection struct {
	host       string
	database   string
	collection string
	client     *mongo.Client
}

// NewConnection creates an instance of connection
func NewConnection(host string, database string, collection string) *Connection {
	opts := options.Credential{
		AuthSource: "admin", Username: "root", Password: "toor",
	}

	clientOptions := options.Client().ApplyURI(host).SetAuth(opts)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return &Connection{
		host,
		database,
		collection,
		client,
	}
}

// Close terminates the connection
func (c *Connection) Close() {
	err := c.client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
}
