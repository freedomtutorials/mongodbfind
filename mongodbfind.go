package main

import (
    "context"
    "fmt"
    "log"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Device struct 
{
	ID int
	Name string
}


func main() {
	// set the mongodb uri 
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	
	//connect to mongodb
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

	freedomtutorialsDB := client.Database("freedomtutorials")
	mydataCollection := freedomtutorialsDB.Collection("mydata")

	
	cursor, err := mydataCollection.Find( context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	
	var devices []Device
	
	/*
	if err = cursor.All( context.TODO(), &devices); err != nil {
		log.Fatal(err)
	}
	fmt.Println(devices)
	*/
	
	for cursor.Next( context.TODO()) {
		var device Device
		if err = cursor.Decode( &device); err != nil {
			log.Fatal(err)
		}
		devices = append( devices, device)
	}
	
	fmt.Println(devices)
}


