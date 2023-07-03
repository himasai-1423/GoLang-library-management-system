package main

import (
	"fmt"
	controller "lib-mng-sys/controllers"
	model "lib-mng-sys/models"
	"lib-mng-sys/views"
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
		choice, userName := views.UserChoice()

		switch choice {
		case 1:
			controller.RentBook(coll, ctx, userName)
		case 2:
			controller.ReturnBook(coll, ctx, userName)
		case 3:
			controller.BooksAvailable(coll, ctx)
		default:
			Leaving(userName)
			shouldExit = true
		}
	}
}

func Leaving(userName string) {
	fmt.Printf("Thank you %v for visiting\n", userName)
}
