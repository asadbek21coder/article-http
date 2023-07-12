package main

import (
	"fmt"
	"net/http"

	"github.com/asadbek21coder/article-http/handlers"
)

const PORT = "8000"

func main() {
	// fmt.Println("Hello world")

	// Endpoints
	http.HandleFunc("/", handlers.GetHomePage)

	// Articles CRUD
	http.HandleFunc("/article", handlers.HandleArticle)

	// Person CRUD
	http.HandleFunc("/person", handlers.HandlePerson)

	// Author
	http.HandleFunc("/author", handlers.GetAuthors)

	//  Building server
	fmt.Println("Server is working on port: ", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
