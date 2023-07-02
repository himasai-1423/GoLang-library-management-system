package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func BooksAvailable(coll *mongo.Collection, ctx context.Context) {
	fmt.Println("Here is the list of available books:- ")
	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var bookList bson.M
		if err = cursor.Decode(&bookList); err != nil {
			panic(err)
		}
		fmt.Println(bookList)
	}
}
