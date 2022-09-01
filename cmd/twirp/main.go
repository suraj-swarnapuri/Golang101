package main

import (
	"net/http"

	server "github.com/suraj-swarnapuri/Golang101/internal/twirptest"
	pb "github.com/suraj-swarnapuri/Golang101/rpc/twirptest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	
    server := &server.HelloWorld{}
	twirpHandler := pb.NewHelloWorldServer(server)
    // You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    
    r.Mount(twirpHandler.PathPrefix(), twirpHandler)
    
    http.ListenAndServe(":8081", r)
}