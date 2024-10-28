package main

import (
	"context"
	"fmt"
	"time"
	"web_app/server"
)

func main() {

	serv := server.Start(":8080")
	fmt.Println("sever started")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	<-ctx.Done()
	err := server.Close(ctx, serv)

	if err != nil {
		fmt.Println("couldn't shut server down")
	}

	fmt.Println("server closed")

}
