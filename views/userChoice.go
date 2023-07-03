package views

import "fmt"

func UserChoice() (int, string) {
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

func Leaving(userName string) {
	fmt.Printf("Thank you %v for visiting\n", userName)
}
