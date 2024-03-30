package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodbapi/model"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb+srv://pavanmongodb:pavango@cluster0.r1ivk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const db = "prime"
const colname = "watchlist"

//mongodb into variable through pointer
var collection *mongo.Collection

//connect with db
func init() {
	//Client option
	clientoption := options.Client().ApplyURI(connectionstring)

	//Connect to DB
	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection done")

	collection = client.Database(db).Collection(colname)

	//collecton instance
	fmt.Println("collection instance eady")
}

//Mognodb helpers

func insertonevalue(movie model.Prime) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted 1 value:", inserted.InsertedID)
}

//Actual helper

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-allow-Methods", "POST")

	var movie model.Prime
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertonevalue(movie)
	json.NewEncoder(w).Encode(movie)

}
