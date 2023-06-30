package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	//establishing connection
	client, ctx, cancel, err := Connect("mongodb+srv://borahimasaireddy:himu2003@cluster0.daxzqzv.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	defer CloseDB(client, ctx, cancel)
	Ping(client, ctx)

	coll := client.Database("ABCLibrary").Collection("Books")
	// insertRes, err := coll.InsertOne(ctx, bson.D{
	// 	{"name", "One Piece"},
	// 	{"author", "Eiichiro Oda"},
	// 	{"bookId", 2},
	// 	{"genre", bson.A{"Fantasy", "Action", "Adventure", "History"}},
	// 	{"quantity", 12},
	// })

	insertRes, err := coll.InsertMany(ctx, []interface{}{
		bson.D{
			{"name", "Attack On Titan"},
			{"author", "Hajime Isayam"},
			{"bookId", 3},
			{"genre", bson.A{"Fantasy", "Action", "Adventure", "History"}},
			{"quantity", 5},
		},
		bson.D{
			{"name", "Harry Potter"},
			{"author", "J.K. Rowling"},
			{"bookId", 4},
			{"genre", bson.A{"Fantasy", "Magic", "Adventure"}},
			{"quantity", 12},
		},
		bson.D{
			{"name", "Iron Man"},
			{"author", "Stan Lee"},
			{"bookId", 5},
			{"genre", bson.A{"Sci-Fi", "Action"}},
			{"quantity", 12},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(insertRes.InsertedIDs)

	// for {
	// 	choice, userName := userChoice()

	// 	switch choice {
	// 	case 1:
	// 		rentBook()
	// 	case 2:
	// 		returnBook()
	// 	case 3:
	// 		availableBooks()
	// 	default:
	// 		leaving(userName)
	// 	}
	// }
}

// func userChoice() (int, string) {
// 	fmt.Println("\nWelcome to Library Management System!")

// 	fmt.Println("\n\nPlease enter your name: ")
// 	var userName string
// 	fmt.Scan(&userName)

// 	fmt.Println("\nChoose any option:-")
// 	fmt.Println("1. Rent a book ")
// 	fmt.Println("2. Return a book")
// 	fmt.Println("3. Read a book ")
// 	fmt.Println("Any other key to exit")

// 	var choice int
// 	fmt.Scan(&choice)
// 	return choice, userName
// }

// func rentBook() {

// 	// call availableBooks
// 	// check availability
// }

// func returnBook() {
// 	// take book
// 	// call availableBooks
// }

// func availableBooks() {
// 	// fetch books list
// }

// func leaving(userName string) {
// 	fmt.Printf("Thank you %v for visiting", userName)
// }
