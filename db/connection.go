package db

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func getConnetction() (client  *mongo.Client, ctx context.Context) {
	
	var cred options.Credential

	cred.Username = "adm"
	cred.Password = "123"
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(cred))
	if err != nil {
		log.Fatal(err)
	}
	return
}