package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/italosm/go-erp/model"
	"context"
	"log"
	"github.com/italosm/go-erp/database"
)

func InsertProduct(product model.Product){
	client, _:= database.ConnectDatabase()
	collection := client.Database("my_database").Collection("products")
	_, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetProductBySaleCost(saleCost int64) []model.Product{
	client, _ := database.ConnectDatabase()
	collection := client.Database("my_database").Collection("products")
	cursor, error := collection.Find(context.TODO(), bson.M{"salecost":saleCost})
	if(error != nil){
		log.Fatal(error.Error())
	}
	var products []model.Product
	strProds := bson.M{}
	cursor.All(context.TODO(), &strProds)
	cursor.Decode(&products)
	fmt.Println(strProds)
	return products
}