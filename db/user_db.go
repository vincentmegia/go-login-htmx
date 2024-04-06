package db

import (
	"fmt"

	"github.com/vincentmegia/go-login-htmx/models"
)

func GetAllUsers() ([]models.User, error) {
	database, error := OpenDatabase("./customers.db")
	if error != nil {
		fmt.Println("Failed to open customers database")
		return nil, error
	}
	rows, error := database.Query(`SELECT * FROM Users`)
	database.Close()
	if error != nil {
		fmt.Println("Failed to query users db")
		return nil, error
	}

	users := []models.User{}
	for rows.Next() {
		var user models.User
		error := rows.Scan(&user.Id, &user.Firstname, &user.Lastname,
			&user.AddressId, &user.Username, &user.Password,
			&user.Salt, &user.CreatedDate, &user.UpdatedDate,
			&user.Email, &user.BirthDate)

		if error != nil {
			fmt.Println("failed to read row, skiping", error)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
