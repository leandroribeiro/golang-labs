package main

import (
	"context"
	"fmt"
	"github.com/leandroribeiro/go-labs/crud-with-mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	DBName = "goLabs"
	notesCollection = "notes"
	URI = "mongodb://mestre:siga-o-mestre@127.0.0.1"
)

func main()  {
	ctx := context.Background()

	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		fmt.Println(err)
		return
	}

	db := client.Database(DBName)
	coll := db.Collection(notesCollection)
	fmt.Println(db.Name())
	fmt.Println(coll.Name())

	note := model.Note{}

	// Insert one document
	// ********************

	// An ID for MongoDB.
	note.ID = primitive.NewObjectID()
	note.Title = "First note"
	note.Body = "Some spam text"
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()
	result, err := coll.InsertOne(ctx, note)
	if err != nil {
		fmt.Println(err)
		return
	}    // ID of the inserted document.
	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)

	// Insert Many Documents.
	// ***********************

	var notes []interface{}

	note2 := model.Note{}
	note2.ID = primitive.NewObjectID()
	note2.Title = "Second note"
	note2.Body = "Some spam text"
	note2.CreatedAt = time.Now()
	note2.UpdatedAt = time.Now()

	note3 := model.Note{
		ID:        primitive.NewObjectID(),
		Title:     "Third note",
		Body:      "Some spam text",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	notes = append(notes, note2, note3)
	results, err := coll.InsertMany(ctx, notes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(results.InsertedIDs)

	// Update one document
	// ********************

	// Parsing a string ID to ObjectID from MongoDB.
	objID, err := primitive.ObjectIDFromHex("5e87b0cba1867e4fcfeea980")
	if err != nil {
		fmt.Println(err)
		return
	}

	resultUpdate, err := coll.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"body":       "Some updated text",
				"updated_at": time.Now(),
			},
		},
	)

	fmt.Println(resultUpdate.ModifiedCount) // output: 1


	// Delete one document.
	// ********************

	objID, err = primitive.ObjectIDFromHex("5e87b162566d211cb1b9b441")
	if err != nil {
		fmt.Println(err)
		return
	}

	resultDelete, err := coll.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resultDelete.DeletedCount) // output: 1


	// Find documents.
	// ********************

	objID, err = primitive.ObjectIDFromHex("5e87b296175afff4c9fc6671")
	if err != nil {
		fmt.Println(err)
		return
	}

	findResult := coll.FindOne(ctx, bson.M{"_id": objID})
	if err := findResult.Err(); err != nil {
		fmt.Println(err)
		return
	}

	n := model.Note{}
	err = findResult.Decode(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n.Body) // output: Some updated text

	// Find many documents
	// ********************

	notesResult := []model.Note{}

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(ctx) {
		cursor.Decode(&n)
		notesResult = append(notesResult, n)
	}

	for _, el := range notesResult {
		fmt.Println(el.Title)
	}
}
