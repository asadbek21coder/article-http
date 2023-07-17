package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asadbek21coder/article-http/models"
)

func HandlePerson(w http.ResponseWriter, r *http.Request) {
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
	var newPerson models.Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if len(models.People) == 0 {
		newPerson.ID = 1
	} else {
		max := models.People[0].ID
		for i := 0; i < len(models.People); i++ {
			if models.People[i].ID > max {
				max = models.People[i].ID
			}
		}
		newPerson.ID = max + 1
	}
	fmt.Println(newPerson.FirstName + " Created")
	models.People = append(models.People, newPerson)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.People)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {

	var newPerson models.Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	index := -1
	for i := 0; i < len(models.People); i++ {
		if newPerson.ID == models.People[i].ID {
			index = i
		}
	}

	if index == -1 {
		fmt.Println("No person with such id")
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(models.People[index].FirstName + " Updated to " + newPerson.FirstName)
	models.People = append(models.People[:index], models.People[index+1:]...)
	models.People = append(models.People, newPerson)
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
