package db

import "context"


func Insert(collection string, data any) error {
	client, ctx := getConnetction()
	defer client.Disconnect(ctx)

	c := client.Database("webform").Collection(collection)
	_, err := c.InsertOne(context.Background(), data)

	return err
}