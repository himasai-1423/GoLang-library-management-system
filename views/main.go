package main

import (
	"fmt"
	controller "lib-mng-sys/controllers"
	model "lib-mng-sys/models"
)

func main() {
	// # Establishing connection
	client, ctx, cancel, err := model.Connect("mongodb+srv://borahimasaireddy:himu2003@cluster0.daxzqzv.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	defer model.CloseDB(client, ctx, cancel)
	model.Ping(client, ctx)

	// # Adding data
	coll := client.Database("ABCLibrary").Collection("Books")
	// AddData(coll, ctx)

	shouldExit := false
	for !shouldExit {
		choice, userName := userChoice()

		switch choice {
		case 1:
			controller.RentBook(coll, ctx)
		case 2:
			controller.ReturnBook(coll, ctx)
		case 3:
			controller.BooksAvailable(coll, ctx)
		default:
			leaving(userName)
			shouldExit = true
		}
	}
}

func userChoice() (int, string) {
	fmt.Println("\nWelcome to Library Management System!")

	fmt.Println("\n\nPlease enter your name: ")
	var userName string
	fmt.Scan(&userName)

	fmt.Printf("\n Hello! %v,", userName)
	fmt.Println("\nPlease choose any of the option below:-")
	fmt.Println("1. Rent a book ")
	fmt.Println("2. Return a book")
	fmt.Println("3. Read a book ")
	fmt.Println("Any other key to exit")

	var choice int
	fmt.Scan(&choice)
	return choice, userName
}

func leaving(userName string) {
	fmt.Printf("Thank you %v for visiting\n", userName)
}
