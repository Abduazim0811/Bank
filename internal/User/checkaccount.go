package user

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	md "github.com/Abduazim0811/Bank/internal/models"
)

func CheckAccount(db *sql.DB) {
	var user md.Users
	name := "../internal/DB/userSelect.sql"
	sqlfile, err := os.ReadFile(name)
	if err != nil {
		log.Printf("Error reading SQL file: %v", err)
		return
	}

	fmt.Print("Enter the user email: ")
	fmt.Scanln(&user.Email)

	row := db.QueryRow(string(sqlfile), user.Email)
	err = row.Scan(&user.Firstname, &user.Lastname, &user.Fathers_name, &user.Email, &user.Password,&user.Balans)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with that email.")
		} else {
			log.Printf("Error querying user: %v", err)
		}
		return
	}

	fmt.Printf(`
		Firstname:    %s
		Lastname:     %s
		Fathers_name: %s
		Email:        %s
		Balans:       %.2f 
	`, user.Firstname, user.Lastname, user.Fathers_name, user.Email, user.Balans)
}
