package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shreyank031/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"Watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	//This result gives you how many value is being updated
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)

}

// delete 1 record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId) //It converts string to object id which is acceptable/understood by mongoDB
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie deleted with delete count: ", deleteCount)

}

// Delete all values
func deleteAllMovie() int64 {
	deleteRes, err := collection.DeleteMany(context.Background(), bson.D{{}})
	//or
	//deleteRes, err := collection.DeleteMany(context.Background(), bson.D{{}},nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted: ", deleteRes.DeletedCount)
	//or
	return deleteRes.DeletedCount
}

// Get all movies from database
func getAllMovies() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M //or []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// Actual controller function
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ := json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params)
}
