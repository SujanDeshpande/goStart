package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	connectUri := "mongodb://localhost:6565"
	fmt.Println("Hello Mongo")
	fmt.Println(connectUri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectUri))
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("Connected")
	}
	coll := client.Database("photoStore").Collection("files")
	var result bson.D
	coll.FindOne(context.TODO(), bson.D{}).Decode(&result)
	if err != nil {
		fmt.Println("error", err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	fmt.Println("End Mongo")
}
