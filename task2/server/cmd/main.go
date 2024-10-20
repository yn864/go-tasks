package main

import (
	"context"
	"time"
	"web_app/server"
)

func main() {

	serv := server.Start(":8080")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server.Close(ctx, serv)
}
