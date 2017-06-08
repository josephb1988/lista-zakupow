package main

import (
    "net/http"
    "os"
)
 
func main() {

	router := NewRouter()
    err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
    if err != nil {
      panic(err)
    }
}