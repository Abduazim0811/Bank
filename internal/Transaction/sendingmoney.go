package transaction

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	md "github.com/Abduazim0811/Bank/internal/models"
)

func SendingMoney(db *sql.DB) {
	var (
		user md.Users
		sum  float64
	)
	fmt.Print("Who will you transfer to?\nFirstname>>>>  ")
	fmt.Scanln(&user.Email)
	fmt.Print("How much will you transfer: ")
	fmt.Scanln(&sum)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	name := "../internal/DB/sendinguser.sql"

	sqlfile, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(string(sqlfile), sum, user.Email)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transfer money!!!!")

}
