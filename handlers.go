package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!")
}
var todos = Todos{
    Todo{Name: "Write presentation"},
    Todo{Name: "Host meetup"},
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    iTodoID, err:= strconv.Atoi(todoId)

    if ( err != nil ){
        panic(err)
    }
    
    fmt.Fprintln(w, "Todo show:", todoId)
    fmt.Fprintln(w, "Todo show:", todoId)

    if ( len(todos) > iTodoID ){
        if err := json.NewEncoder(w).Encode(todos[iTodoID]); err != nil {
            panic(err)
        }
    }else{
    fmt.Fprintln(w, "No such todo")
    }
}