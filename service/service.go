package service

import (
	"context"
	"log"

	"github.com/Garry028/mongoApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertOneMovie inserts a movie into the database
func InsertOneMovie(collection *mongo.Collection, movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted one movie in db with id:", inserted.InsertedID)
}

// UpdateOneMovie updates a movie's "watched" status
func UpdateOneMovie(collection *mongo.Collection, movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Modified count:", result.ModifiedCount)
}

// DeleteOneMovie deletes a movie by ID
func DeleteOneMovie(collection *mongo.Collection, movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted count:", result.DeletedCount)
}

// DeleteAllMovies deletes all movies from the collection
func DeleteAllMovies(collection *mongo.Collection) int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted count:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// GetAllMovies retrieves all movies from the collection
func GetAllMovies(collection *mongo.Collection) []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}
