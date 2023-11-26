package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var Users = []User{
	{ID: 1, Username: "admin", Password: "adminpass", Role: "admin"},
	{ID: 2, Username: "teacher", Password: "teacherpass", Role: "teacher"},
	{ID: 3, Username: "student", Password: "studentpass", Role: "student"},
}
