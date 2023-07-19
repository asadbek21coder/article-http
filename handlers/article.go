package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/asadbek21coder/article-http/models"
)

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
	var newArticle models.Article
	json.NewDecoder(r.Body).Decode(&newArticle)
	newArticle.CreatedAt = time.Now()
	read, _ := os.ReadFile("db/article.json")
	var articles []models.Article
	json.Unmarshal(read, &articles)
	if len(articles) == 0 {
		newArticle.ID = 1
	}
	max := articles[0].ID
	for i := 0; i < len(articles); i++ {
		if articles[i].ID > max {
			max = articles[i].ID
		}
	}
	newArticle.ID = max + 1
	articles = append(articles, newArticle)
	data, _ := json.Marshal(articles)
	os.WriteFile("db/article.json", data, 0)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newArticle)
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	jsonData, _ := os.ReadFile("db/article.json")

	var articles []models.Article
	var articlesResponse []models.ArticleGetAll
	json.Unmarshal(jsonData, &articles)
	for i := 0; i < len(articles); i++ {
		var people []models.Person
		peopleJson, _ := os.ReadFile("db/people.json")
		json.Unmarshal(peopleJson, &people)

		newArticle := models.ArticleGetAll{
			ID:      articles[i].ID,
			Content: articles[i].Content,
			Author: models.Person{
				ID: articles[i].AuthorId,
			},
			CreatedAt: articles[i].CreatedAt,
		}
		for j := 0; j < len(people); j++ {
			if people[j].ID == articles[i].AuthorId {
				newArticle.Author.FirstName = people[j].FirstName
				newArticle.Author.LastName = people[j].LastName
			}
		}

		articlesResponse = append(articlesResponse, newArticle)

	}
	res, _ := json.Marshal(articlesResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(res))
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	var newArticle models.Article
	json.NewDecoder(r.Body).Decode(&newArticle)

	read, _ := os.ReadFile("db/article.json")

	var articles []models.Article
	json.Unmarshal(read, &articles)

	index := -1
	for i := 0; i < len(articles); i++ {
		if articles[i].ID == newArticle.ID {
			index = i
		}
	}

	articles = append(articles[:index], articles[index+1:]...)
	newArticle.CreatedAt = time.Now()
	articles = append(articles, newArticle)
	jsonArticles, _ := json.Marshal(articles)

	os.WriteFile("db/article.json", jsonArticles, 0)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newArticle)

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {

}
