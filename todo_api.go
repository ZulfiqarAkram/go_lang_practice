package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type todoItem struct {
	ID   int    `json:id`
	Text string `json:text`
}

var todoDB []todoItem

func main() {
	todoDB = append(todoDB, todoItem{
		ID:   1,
		Text: "hello",
	}, todoItem{
		ID:   2,
		Text: "world",
	})

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/todo", DisplayItems).Methods(http.MethodGet)
	api.HandleFunc("/todo", AddItem).Methods(http.MethodPost)
	api.HandleFunc("/todo/{id}", UpdateItem).Methods(http.MethodPut)
	api.HandleFunc("/todo/{id}", DeleteItem).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTodo todoItem
	json.NewDecoder(r.Body).Decode(&newTodo)
	newTodo.ID = len(todoDB) + 1
	todoDB = append(todoDB, newTodo)
	fmt.Println("Item added.")
	json.NewEncoder(w).Encode(newTodo)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todoDB {
		id, err := strconv.ParseInt(params["id"], 16, 64)
		if err != nil {
			fmt.Println(err)
		}
		if item.ID == int(id) {
			todoDB = append(todoDB[:index], todoDB[index+1:]...)
			fmt.Println("item removed.")
			return
		}
	}

}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range todoDB {
		id, err := strconv.ParseInt(params["id"], 16, 64)
		if err != nil {
			fmt.Println(err)
		}
		if item.ID == int(id) {
			todoDB = append(todoDB[:index], todoDB[index+1:]...)
			var todoToBeUpdate todoItem
			json.NewDecoder(r.Body).Decode(&todoToBeUpdate)
			todoToBeUpdate.ID = int(id)
			todoDB = append(todoDB, todoToBeUpdate)
			json.NewEncoder(w).Encode(todoToBeUpdate)
			fmt.Println("item updated.")
			return
		}
	}

}

func DisplayItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todoDB)

}
