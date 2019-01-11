package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type person struct {
	ID        string   `json:"id"`
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Address   *address `json:"address"`
}

var people []person

func getPeople(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(people)
}

func getPerson(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(responseWriter).Encode(item)
			return
		}
	}
	json.NewEncoder(responseWriter).Encode(&person{})
}

func createPerson(responseWriter http.ResponseWriter, request *http.Request) {
	var newPerson person
	_ = json.NewDecoder(request.Body).Decode(&newPerson)
	people = append(people, newPerson)
}

func deletePerson(responseWriter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people/{id}", getPerson).Methods("GET")
	router.HandleFunc("/people/create", createPerson).Methods("POST")
	router.HandleFunc("/people/delete/{id}", deletePerson).Methods("DELETE")

	people = append(people, person{ID: "1", FirstName: "Ivan", LastName: "Ivanov", Address: &address{City: "City X", State: "State X"}})
	people = append(people, person{ID: "2", FirstName: "Petr", LastName: "Petrov", Address: &address{City: "City Y", State: "State Y"}})

	http.ListenAndServe(":37005", router)
}
