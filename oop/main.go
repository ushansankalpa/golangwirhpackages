package main

import (
	"oop/author"
	"oop/employee"
)

func main() {
	e := employee.New("Ushan", "Sankalpa", 30, 20)
	e.LeavesRemaining()

	author1 := author.Author{
		FirstName: "ushan",
		LastName:  "sankalpa",
		Bio:       "Golang Enthusiast",
	}
	blogPost1 := author.BlogPost{
		Title:   "Inheritance in Go",
		Content: "Go supports composition instead of inheritance",
		Author:  author1,
	}
	blogPost1.Details()
}
