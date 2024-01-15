package models

type Address struct {
	Street     string `json:"street,omitempty"`
	PostalCode int32  `json:"postalcode,omitempty"`
}
type User struct {
	Id        int32   `josn:"id,omitempty"`
	Firstname string  `json:"firstname,omitempty"`
	Lastname  string  `json:"lastname,omitempty"`
	Address   Address `json:"address,omitempty"`
}

func CreateUser(id int32, firstname string, lastname string, street string, postalcode int32) User {
	return User{id, firstname, lastname, Address{street, postalcode}}
}
