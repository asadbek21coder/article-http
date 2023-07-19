package models

import "time"

type Article struct {
	ID        int
	Content   Content
	AuthorId  int
	CreatedAt time.Time
}

type ArticleGetAll struct {
	ID        int
	Content   Content
	Author    Person
	CreatedAt time.Time
}

type Content struct {
	Title string
	Body  string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

// type GetPersonExtended struct {
// 	ID        int
// 	FirstName string
// 	LastName  string
// 	Articles  []Article
// }

type Author struct {
	Person   Person
	Articles []ArticleSmall
}

type ArticleSmall struct {
	ID    int
	Title string
	// 
}

// "/author" [GET] method

type DeleteUserRequest struct {
	ID int
}

var People []Person
var Articles []Article
var ArticlesGetAll []ArticleGetAll

var Authors []Author
