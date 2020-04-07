package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	Firstname string             `json:"firstname, omitempty" bson:"firstname, omitempty"`
	Lastname  string             `json:"lastname, omitempty" bson:"lastname, omitempty"`
}

var client *mongo.Client

func CreatePersonEndpoint(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Add("content-type", "application/json")

	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	person.ID = primitive.NewObjectID()

	collection := client.Database("api-lab1").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := collection.InsertOne(ctx, person)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(`{"message":"` + err.Error() + `"} `))
		return
	}

	json.NewEncoder(responseWriter).Encode(result)
}

func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("content-type", "application/json")
	var people []Person
	collection := client.Database("api-lab1").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{}, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"} `))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx){
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"} `))
		return
	}
	json.NewEncoder(w).Encode(people)

}

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	collection := client.Database("api-lab1").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"} `))
		return
	}
	json.NewEncoder(w).Encode(person)
}

func main() {
	println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://mestre:siga-o-mestre@127.0.0.1")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}
