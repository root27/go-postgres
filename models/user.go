package models

type User struct {
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}
