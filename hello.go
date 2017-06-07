package main

import (
    "fmt"
    "net/http"
    "os"
	"html"
	
	"github.com/gorilla/mux"
)

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Hello, %q", html.EscapeString(req.URL.Path) )
}
