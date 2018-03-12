package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    
    "github.com/gorilla/mux"
)
 
type Item struct {
    ID string   `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
}

var ShoppingList []Item


func GetItemEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range ShoppingList {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Item{})
}

func GetItemsEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(ShoppingList)
}

func AddItemEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var item Item
    _ = json.NewDecoder(req.Body).Decode(&item)
    item.ID = params["id"]
    ShoppingList = append(ShoppingList, item)
    json.NewEncoder(w).Encode(ShoppingList)
}

func RemoveItemEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range ShoppingList {
        if item.ID == params["id"] {
            ShoppingList = append(ShoppingList[:index], ShoppingList[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(ShoppingList)
}


func main() {
    router := mux.NewRouter()
    ShoppingList = append(ShoppingList, Item{ID: "1", Name: "Chleb"})
    ShoppingList = append(ShoppingList, Item{ID: "2", Name: "Maslo"})
    router.HandleFunc("/items", GetItemsEndpoint).Methods("GET")
    router.HandleFunc("/items/{id}", GetItemEndpoint).Methods("GET")
    router.HandleFunc("/items/{id}", AddItemEndpoint).Methods("POST")
    router.HandleFunc("/items/{id}", RemoveItemEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
