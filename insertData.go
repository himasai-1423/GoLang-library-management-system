package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddData(coll *mongo.Collection, ctx context.Context) {
	insRes, err := coll.InsertOne(ctx, bson.D{
		{Key: "name", Value: "One Piece"},
		{Key: "author", Value: "Eiichiro Oda"},
		{Key: "bookId", Value: 2},
		{Key: "genre", Value: bson.A{"Fantasy", "Action", "Adventure", "History"}},
		{Key: "quantity", Value: 12},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(insRes.InsertedID)

	insertRes, err := coll.InsertMany(ctx, []interface{}{
		bson.D{
			{Key: "name", Value: "Attack On Titan"},
			{Key: "author", Value: "Hajime Isayam"},
			{Key: "bookId", Value: 3},
			{Key: "genre", Value: bson.A{"Fantasy", "Action", "Adventure", "History"}},
			{Key: "quantity", Value: 5},
		},
		bson.D{
			{Key: "name", Value: "Harry Potter"},
			{Key: "author", Value: "J.K. Rowling"},
			{Key: "bookId", Value: 4},
			{Key: "genre", Value: bson.A{"Fantasy", "Magic", "Adventure"}},
			{Key: "quantity", Value: 12},
		},
		bson.D{
			{Key: "name", Value: "Iron Man"},
			{Key: "author", Value: "Stan Lee"},
			{Key: "bookId", Value: 5},
			{Key: "genre", Value: bson.A{"Sci-Fi", "Action"}},
			{Key: "quantity", Value: 12},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(insertRes.InsertedIDs)

}
