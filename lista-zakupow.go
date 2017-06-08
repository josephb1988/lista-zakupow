package main

import (
    "fmt"
    "net/http"
    "os"
	
	"github.com/gorilla/mux"
)
 
func main() {

	router := NewRouter()
    err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
    if err != nil {
      panic(err)
    }
}