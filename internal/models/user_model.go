package models

import (
	"time"
)

type User struct {
	Id          int32     `josn:"id,omitempty"`
	Firstname   string    `json:"firstname,omitempty"`
	Lastname    string    `json:"lastname,omitempty"`
	AddressId   int32     `json:"addressid,omitempty"`
	Username    string    `json:"username,omitempty"`
	Password    string    `json:"password,omitempty"`
	Salt        string    `json:"salt,omitempty"`
	CreatedDate time.Time `json:"createddate,omitempty"`
	UpdatedDate time.Time `json:"upadteddate,omitempty"`
	Email       string    `json:"email,omitempty"`
	BirthDate   time.Time `json:"birthdate,omitempty"`
}

func CreateUser(id int32,
	firstname string,
	lastname string,
	AddressId int32,
	Username string,
	Password string,
	Salt string,
	CreatedDate time.Time,
	UpdatedDate time.Time,
	Email string,
	BirthDate time.Time,
) *User {
	return &User{id, firstname, lastname, AddressId, Username, Password, Salt, CreatedDate, UpdatedDate, Email, BirthDate}
}
