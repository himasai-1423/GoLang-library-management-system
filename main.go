package main

func main() {
	client, ctx, cancel, err := Connect("mongodb+srv://borahimasaireddy:himu2003@cluster0.daxzqzv.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}

	defer CloseDB(client, ctx, cancel)

	Ping(client, ctx)
}
