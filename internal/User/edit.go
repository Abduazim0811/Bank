package user

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	md "github.com/Abduazim0811/Bank/internal/models"
)

func Update(db *sql.DB, column string,user md.Users) {
	var str string
	fmt.Print(column,": ")
	fmt.Scanln(&str)
	name := "../internal/DB/updatefirstname.sql"
	sqlfile, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(string(sqlfile),column,str,user.Email)
	if err != nil {
		log.Fatal(err)
	}
}

func Edit(db *sql.DB) {
	var (
		num  int
		user md.Users
		str string
	)
	fmt.Println("Nimani o'zgartirmoqchisiz?")
	fmt.Println(`[1]firstname
[2]lastname
[3]fathersname
[4]email
[5]password`)
	fmt.Scanln(&num)
	switch num {
	case 1:
		str="firstname"
		Update(db,str,user)
	case 2:
		str="lastname"
		Update(db,str,user)
	case 3:
		str="fathersname"
		Update(db,str,user)
	case 4:
		str="email"
		Update(db,str,user)
	case 5:
		str="password"
		Update(db,str,user)
	}
}
