package model

import "time"

// User structure stores user data
type User struct {
	ID        int
	Name      string
	Lastname  string
	Age       int
	Birthdate time.Time
}
