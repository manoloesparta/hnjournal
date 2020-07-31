package database

import (
	"context"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Output is the values returned by the API
type Output struct {
	Title string
	URL   string
}

// GetRandom gets the output of the API
func (c *Connection) GetRandom() Output {
	all := c.getAll()

	rand.Seed(time.Now().UTC().UnixNano())

	index := rand.Intn(len(all) - 1)

	output := Output{
		Title: all[index]["title"].(string),
		URL:   all[index]["url"].(string),
	}

	return output
}

func (c *Connection) getAll() []bson.M {
	collection := c.client.Database(c.database).Collection(c.collection)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	must(err)

	var all []bson.M
	err = cursor.All(context.TODO(), &all)
	must(err)

	cursor.Close(context.TODO())

	return all
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}
