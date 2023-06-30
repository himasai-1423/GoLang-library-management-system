package main

import "fmt"

func main() {
	//establishing connection
	client, ctx, cancel, err := Connect("mongodb+srv://borahimasaireddy:himu2003@cluster0.daxzqzv.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	defer CloseDB(client, ctx, cancel)
	Ping(client, ctx)

	fmt.Println("\nWelcome to Library Management System!")

	for {
		choice := userChoice()

		switch choice {
		case 1:
			rentBook()
		case 2:
			returnBook()
		case 3:
			availableBooks()
		default:
			leaving()
		}
	}
}

func userChoice() int {
	fmt.Println("\nChoose any option:-")
	fmt.Println("1. Rent a book ")
	fmt.Println("2. Return a book")
	fmt.Println("3. Read a book ")
	fmt.Println("Any other key to exit")

	var choice int
	fmt.Scan(&choice)
	return choice
}

func rentBook() {

}

func returnBook() {

}

func availableBooks() {

}

func leaving() {

}
