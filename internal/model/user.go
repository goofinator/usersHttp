package model

import "time"

// User structure stores user data
type User struct {
	ID        int       `json:"Id,omitempty"`
	Name      string    `json:"Name"`
	Lastname  string    `json:"Lastname"`
	Age       int       `json:"Age,omitempty"`
	Birthdate time.Time `json:"Birthdate"`
}
