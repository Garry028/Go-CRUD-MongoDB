// package controller

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/Garry028/mongoApi/model"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const connectionString = "mongodb://localhost:27017"
// const dbName = "netflix"
// const colName = "watchlist"

// // MOST IMP
// var collection *mongo.Collection

// // connect with mongoDB

// // this method run only once
// func init() {
// 	// client options
// 	clientOption := options.Client().ApplyURI(connectionString)
// 	// create a new mongodb client instance using the above option and then establishing it to server
// 	client, err := mongo.Connect(context.TODO(), clientOption)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// check the connection
// 	err = client.Ping(context.Background(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB")

// 	collection := client.Database(dbName).Collection(colName)

// 	fmt.Println("Collection instance created", collection)

// }

// // MongoDB helpers - file

// // insert one record
// func insertOneMovie(movie model.Netflix) {
// 	inserted, err := collection.InsertOne(context.Background(), movie)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Inserted one movie in db with id: ", inserted.InsertedID)
// }

// // update one record
// func updateOneMovie(movieId string) {
// 	id, _ := primitive.ObjectIDFromHex(movieId) // movie id will converted in mongoDB format
// 	filter := bson.M{"_id": id}
// 	update := bson.M{"$set": bson.M{"watched": true}}

// 	result, err := collection.UpdateOne(context.Background(), filter, update)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Modified count: ", result.ModifiedCount)
// }

// // delete one record
// func deleteOneMovie(movieId string) {
// 	id, _ := primitive.ObjectIDFromHex(movieId)
// 	filter := bson.M{"_id": id}
// 	result, err := collection.DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Deleted count: ", result)
// }

// // delete all the movies
// func deleteAllMovies() int64 {
// 	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Deleted count: ", deleteResult.DeletedCount)
// 	return deleteResult.DeletedCount
// }

// // get all movies from database

// func getAllMovies() []primitive.M {
// 	cur, err := collection.Find(context.Background(), bson.D{{}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var movies []primitive.M
// 	//In MongoDB Go driver, primitive.M is a type that represents a BSON document (a set of key-value pairs) in a Go-friendly way

// 	for cur.Next(context.Background()) {
// 		var movie bson.M
// 		err := cur.Decode(&movie)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		movies = append(movies, movie)
// 	}
// 	defer cur.Close(context.Background())
// 	return movies
// }

// // Actual controller - file
// func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type","application/x-www-form-urlencode")
// 	allMovies := getAllMovies()
// 	json.NewEncoder(w).Encode(allMovies)
// }

package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Garry028/mongoApi/database"
	"github.com/Garry028/mongoApi/service"
)

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	client := database.Client
	collection := client.Database("netflix").Collection("watchlist")

	movies := service.GetAllMovies(collection)

	json.NewEncoder(w).Encode(movies)
	// Handle errors if needed
}
