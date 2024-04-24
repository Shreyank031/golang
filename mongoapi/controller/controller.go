package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/Shreyank031/mongoapi/model"
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

// MongoDB helper - file. For insertion of data into databse. A simple helper method
//insert 1 record. It takes some data and adds the data into the mongoDB

func insertOneMovie(movie model.Netflix) { //As it's a helper method, we won't be exporting it to anywhere. so small case
	inserted, err := collection.InsertOne(context.Background(), movie) //inserts one value
	if err != nil {
		log.Fatal(err)
	}
	//whenever you add a value into the database, that value gets/recieves the unique id form db.
	//And that unique id as a success  recieves back in the below operation.
	fmt.Println("Inserted 1 movie in database with id: ", inserted.InsertedID)
}
