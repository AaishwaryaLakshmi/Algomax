package models

type Course struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var Courses = []Course{
	{ID: 1, Title: "Math 101", Description: "Introduction to Mathematics"},
	{ID: 2, Title: "History 101", Description: "World History Overview"},
}
