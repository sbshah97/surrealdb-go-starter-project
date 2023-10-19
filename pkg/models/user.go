package models

type User struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
