package models

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // Password won't be included in JSON responses
}

// For demo purposes, we'll use an in-memory store
var Users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "secret"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Password: "secret"},
}