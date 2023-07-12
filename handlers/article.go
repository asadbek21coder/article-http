package handlers

import "net/http"

func HandleArticle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createArticle(w, r)
	case http.MethodGet:
		getArticles(w, r)
	case http.MethodPut:
		updateArticle(w, r)
	case http.MethodDelete:
		deleteArticle(w, r)

	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {

}

func getArticles(w http.ResponseWriter, r *http.Request) {

}

func updateArticle(w http.ResponseWriter, r *http.Request) {

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {

}
