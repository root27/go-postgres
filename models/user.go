package models

type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
