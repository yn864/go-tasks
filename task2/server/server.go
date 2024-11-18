package server

import (
	"context"
	"fmt"
	"net/http"
	"web_app/server/handlers"
)

func Start(serverAddress string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", handlers.Version)
	mux.HandleFunc("/decode", handlers.Decode)
	mux.HandleFunc("/hard-op", handlers.HardOp)

	server := http.Server{Addr: serverAddress, Handler: mux}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("error starting server ", err)
		}
	}()

	return &server
}

func Close(ctx context.Context, server *http.Server) error {
	return server.Shutdown(ctx)
}
