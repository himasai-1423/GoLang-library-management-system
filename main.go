package main

import (
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

	views.FlowControl(coll, ctx)
}
