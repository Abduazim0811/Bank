package transaction

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	md "github.com/Abduazim0811/Bank/internal/models"
)

func PutMoney(db *sql.DB) {
	var sum float64
	var user md.Users

	fmt.Print("Enter the amount of distribution: ")
	fmt.Scanln(&sum)

	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return
	}
	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			log.Printf("Failed to rollback transaction: %v", err)
		}
	}()

	name := "../internal/DB/updatemoney.sql"
	sqlfile, err := os.ReadFile(name)
	if err != nil {
		log.Printf("Error reading SQL file: %v", err)
		return
	}

	fmt.Print("Enter the user email: ")
	fmt.Scanln(&user.Email)

	if _, err := tx.Exec(string(sqlfile), sum, user.Email); err != nil {
		log.Printf("Error executing update: %v", err)
		return
	}
	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)
		return
	}

	fmt.Println("Money has been added!")
}
