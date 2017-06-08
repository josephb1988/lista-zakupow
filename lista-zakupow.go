package main

import (
    "fmt"
    "net/http"
    "os"
	"html"
	
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", hello)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Hello, %q", html.EscapeString(req.URL.Path) )
}
