package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"github.com/italosm/go-erp/model"
	"context"
	"log"
	"github.com/italosm/go-erp/database"
)


//InsertProduct inserts a product
func InsertProduct(product model.Product){
	client, _:= database.ConnectDatabase()
	collection := client.Database("my_database").Collection("products")
	_, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal(err.Error())
	}
}
//GetProductBySaleCost finds by sale cost
func GetProductBySaleCost(saleCost int) []model.Product{
	client, _ := database.ConnectDatabase()
	collection := client.Database("my_database").Collection("products")
	cursor, error := collection.Find(context.TODO(), bson.M{"salecost":saleCost})
	if(error != nil){
		log.Fatal(error.Error())
	}
	var products []model.Product
	cursor.All(context.TODO(), &products)
	return products
}

//GetAllProducts gets all products 
func GetAllProducts() []model.Product{
	var products []model.Product
	client, _ := database.ConnectDatabase()
	collection := client.Database("my_database").Collection("products")
	cursor, error := collection.Find(context.TODO(), bson.M{})
	if(error != nil){
		log.Fatal(error.Error())
	}
	cursor.All(context.TODO(), &products)
	return products
}