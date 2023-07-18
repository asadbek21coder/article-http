package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/asadbek21coder/article-http/models"
)

func HandlePerson(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodPost:
		createPerson(w, r)
	case http.MethodGet:
		getPeople(w, r)
	case http.MethodPut:
		updatePerson(w, r)
	case http.MethodDelete:
		deletePerson(w, r)

	}
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var people []models.Person
	var newPerson models.Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	read, _ := os.ReadFile("db/people.json")
	json.Unmarshal(read, &people)
	if len(people) == 0 {
		newPerson.ID = 1
	} else {
		max := people[0].ID
		for i := 0; i < len(people); i++ {
			if people[i].ID > max {
				max = people[i].ID
			}
		}
		newPerson.ID = max + 1
	}
	fmt.Println(newPerson.FirstName + " Created")

	people = append(people, newPerson)
	// models.People = append(models.People, newPerson)
	data, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("db/people.json", []byte(string(data)), 0)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	read, _ := os.ReadFile("db/people.json")

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(read)
	fmt.Fprint(w, string(read))
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	var people []models.Person
	var newPerson models.Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	read, _ := os.ReadFile("db/people.json")
	json.Unmarshal(read, &people)
	index := -1
	for i := 0; i < len(people); i++ {
		if newPerson.ID == people[i].ID {
			index = i
		}
	}

	if index == -1 {
		fmt.Println("No person with such id")
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(people[index].FirstName + " Updated to " + newPerson.FirstName)
	people = append(people[:index], people[index+1:]...)
	people = append(people, newPerson)
	data, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("db/people.json", []byte(string(data)), 0)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newPerson)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	var req models.DeleteUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Can`t parse from json")
		return
	}

	index := -1
	for i := 0; i < len(models.People); i++ {
		if models.People[i].ID == req.ID {
			index = i
		}
	}

	if index == -1 {
		fmt.Println("No person with such id")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No person with such id")
		return
	}
	fmt.Println(models.People[index].FirstName, " deleted")

	models.People = append(models.People[:index], models.People[index+1:]...)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted successfully\n")
}
