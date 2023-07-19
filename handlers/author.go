package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/asadbek21coder/article-http/models"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var people []models.Person
	readPeople, _ := os.ReadFile("db/people.json")
	json.Unmarshal(readPeople, &people)

	for i := 0; i < len(people); i++ {
		var newAuthor models.Author

		newAuthor.Person = models.Person{
			ID:        people[i].ID,
			FirstName: people[i].FirstName,
			LastName:  people[i].LastName,
		}
		// fmt.Println(people[i].ID)
		var articles []models.Article
		readdArticles, _ := os.ReadFile("db/article.json")
		json.Unmarshal(readdArticles, &articles)

		for j := 0; j < len(articles); j++ {
			var articleSmall models.ArticleSmall
			if articles[j].AuthorId == people[i].ID {
				articleSmall.ID = articles[j].ID
				articleSmall.Title = articles[j].Content.Title
				// 
				newAuthor.Articles = append(newAuthor.Articles, articleSmall)
			}
		}
		models.Authors = append(models.Authors, newAuthor)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Authors)
}
