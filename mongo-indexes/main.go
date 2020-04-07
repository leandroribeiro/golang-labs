package main

import (
	"context"
	"fmt"
	"github.com/leandroribeiro/go-labs/mongo-indexes/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"time"
)

const (
	DBName          = "goLabs"
	notesCollection = "notesWithIndexes"
	URI             = "mongodb://mestre:siga-o-mestre@127.0.0.1"
)

func main() {
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

	// Index Keys
	// **************

	indexOptions := options.Index().SetUnique(true)
	indexKeys := bsonx.MDoc{
		"title": bsonx.Int32(-1),
	}

	noteIndexModel := mongo.IndexModel{
		Options: indexOptions,
		Keys:    indexKeys,
	}

	indexName, err := coll.Indexes().CreateOne(ctx, noteIndexModel)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(indexName) // Output: title_-1

	// Text Indexes
	// ***************

	textIndexModel := mongo.IndexModel{
		Options: options.Index().SetBackground(true),
		Keys: bsonx.MDoc{
			"title": bsonx.String("text"),
			"body":  bsonx.String("text"),
		},
	}

	indexName, err = coll.Indexes().CreateOne(ctx, textIndexModel)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(indexName) // Output: title_text_body_text

	// Making queries
	// ****************

	var notes = []interface{}{
		model.Note{
			ID:        primitive.NewObjectID(),
			Title:     "First note",
			Body:      "Concurrency is not parallelism",
			CreatedAt: time.Now(),
		}, model.Note{
			ID:        primitive.NewObjectID(),
			Title:     "Second note",
			Body:      "A little copying is better than a little dependency",
			CreatedAt: time.Now(),
		}, model.Note{
			ID:        primitive.NewObjectID(),
			Title:     "Third note",
			Body:      "Don't communicate by sharing memory, share memory by communicating",
			CreatedAt: time.Now(),
		}, model.Note{
			ID:        primitive.NewObjectID(),
			Title:     "Fourth note",
			Body:      "Don't just check errors, handle them gracefully",
			CreatedAt: time.Now(),
		},
	}

	_, err = coll.InsertMany(ctx, notes)
	if err != nil {
		fmt.Println(err)
		return
	}

	n := model.Note{}
	fmt.Println("First example")
	cursor, err := coll.Find(ctx, bson.M{
		"$text": bson.M{
			"$search": "note",
		},
	})
	for cursor.Next(ctx) {
		cursor.Decode(&n)
		fmt.Println(n)
	}
	fmt.Println("Second example")
	cursor, err = coll.Find(ctx, bson.M{
		"$text": bson.M{
			"$search": "gracefully",
		},
	})
	for cursor.Next(ctx) {
		cursor.Decode(&n)
		fmt.Println(n)
	}

}
