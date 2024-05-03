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
		log.Fatal(err)
	}
	row := db.QueryRow(string(sqlfile), user.Email)
	err = row.Scan(&user.Firstname, &user.Lastname, &user.Fathers_name, &user.Email, &user.Password, &user.Balans)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(`
		Firstname: 	  %s\n
		Lastname:     %s\n
		Fathers_name: %s\n
		Email:		  %s\n
		Password:	  %s\n
		Balans:		  %f\n	 
	`, user.Firstname, user.Lastname, user.Fathers_name, user.Email, user.Password, user.Balans)
	
}
