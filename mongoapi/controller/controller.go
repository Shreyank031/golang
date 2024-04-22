package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// global variables
const connectionStoring = "" //connection string for MongoDB
const dbName = "netflix"     //Name of the mongoDB database
const colName = "watchList"  //Name of the collection within the database

// IMP
// Take the reference of Mongodb collection = *mongo.Collection-> original reference as we are marking it as pointer
var collection *mongo.Collection

//connect with mongoDB
//Create a init method: It is a specialized method is golang which runs only at very first time entire application is going to execute
//In Go, the init function is a special function that is called automatically, without being explicitly invoked, when a package is initialized.
//The init function is commonly used to initialize package-level variables or perform setup tasks that need to be executed before the program starts running.

// it run first time. only one time.
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionStoring)

	//connect to mongoDb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	//Till here we just made data base connection, now it's time to add data into to the database
	collection = client.Database(dbName).Collection(colName)
	//collection instasnce
	fmt.Println("Collection instasnce is ready")
}
