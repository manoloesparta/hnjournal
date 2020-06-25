package urls

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

// Insert data into mongodb
func (c *Connection) Insert(link URL) {
	bsonlink := bson.D{
		{Key: "title", Value: link.title},
		{Key: "url", Value: link.link},
	}

	collection := c.client.Database(c.database).Collection(c.collection)
	res := collection.FindOne(context.TODO(), bsonlink)

	if res.Err() == nil {
		return
	}

	_, err := collection.InsertOne(context.TODO(), bsonlink)
	fmt.Print("s")

	if err != nil {
		log.Fatal(err)
	}
}

// Close terminates the connection
func (c *Connection) Close() {
	err := c.client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
}
