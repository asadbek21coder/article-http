package handlers

import (
	"encoding/json"
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

	// max ni aniqlab +1 id ga beradi
	newPerson.ID = len(models.People) + 1
	models.People = append(models.People, newPerson)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

func getPeople(w http.ResponseWriter, r *http.Request) {

}

func updatePerson(w http.ResponseWriter, r *http.Request) {

}

func deletePerson(w http.ResponseWriter, r *http.Request) {

}
