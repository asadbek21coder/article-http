package models

import "time"

type Article struct {
	ID        int
	Content   Content
	Author    Person
	CreatedAt *time.Time
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

type Author struct {
	Person   Person
	Articles []Article
}

var People []Person
var Articles []Article
// var Authors []Author
