package views

import (
	"context"
	controller "lib-mng-sys/controllers"

	"go.mongodb.org/mongo-driver/mongo"
)

func FlowControl(coll *mongo.Collection, ctx context.Context) {
	shouldExit := false
	for !shouldExit {
		choice, userName := UserChoice()

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
