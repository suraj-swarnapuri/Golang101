package main

import (
	"net/http"

	"github.com/suraj-swarnapuri/Golang101/internal/server"
)
func main(){
	http.HandleFunc("/hello",server.Hello)

	http.ListenAndServe(":8081",nil)
}