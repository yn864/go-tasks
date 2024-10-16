package main

import (
	"context"
	// "fmt"
	// "net/http"
	"time"
	"web_app/client"
	"web_app/server"
)

func main() {

	serv := server.Start(":8080")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := client.NewClient("http://localhost:8080")

	client.GetVersion()
	client.DecodeMessage("SGVsbG8gV29ybGQh")
	client.HardOp()

	server.Close(ctx, serv)
}
