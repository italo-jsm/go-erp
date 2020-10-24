package database

import (
	"context"
	"log"
	"time"
 
	_"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_"go.mongodb.org/mongo-driver/mongo/readpref"
 )

func ConnectDatabase() (*mongo.Client, context.Context){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
    if err != nil {
        log.Fatal(err.Error())
    }
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err.Error())
    }
	return client, ctx
}