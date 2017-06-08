package main

import (
    "fmt"
    "net/http"
    "os"
	
	"github.com/gorilla/mux"
)
 
func main() {

	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
    if err != nil {
      panic(err)
    }
}

func Index(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Hello!" )
}


func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}