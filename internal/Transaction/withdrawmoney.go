package transaction

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	md "github.com/Abduazim0811/Bank/internal/models"
)

func WitdrawMoney(db *sql.DB) {
	var (
		sum float64
		user md.Users
	)
	fmt.Print("Qancha yechib olmoqchisiz?:  ")
	fmt.Scanln(&sum)
	tx, err:=db.Begin()
	if err!=nil{
		log.Fatal(err)
	}
	defer tx.Rollback()

	name:="../internal/DB/withdrawmoney.sql"

	sqlfile,err:=os.ReadFile(name)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Print("Enter the user email: ")
	fmt.Scanln(&user.Email)

	_,err=tx.Exec(string(sqlfile), sum,user.Email)
	if err!=nil{
		log.Fatal(err)
	}

	err=tx.Commit()
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Hisobizdan pul yechildi")
}
