package repositories

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
		var address_id int32
		error := rows.Scan(&user.Firstname, &user.Lastname, &user.Id, &address_id)
		if error != nil {
			fmt.Println("failed to read row, skiping", error)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
