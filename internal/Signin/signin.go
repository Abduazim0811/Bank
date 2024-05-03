package signin

import (
	"database/sql"
	"fmt"

	md "github.com/Abduazim0811/Bank/internal/models"
	tr "github.com/Abduazim0811/Bank/internal/Transaction"
)

func Signin(db *sql.DB, mp map[string]string) {
	var (
		user md.Users
	)

	fmt.Print("Email: ")
	fmt.Scanln(&user.Email)
	fmt.Print("Password: ")
	fmt.Scanln(&user.Password)

	value, found := mp[user.Email]
	if found {
		if user.Password == value {
			fmt.Println("Siz muvaffaqiyatli kirdingiz!!")
			tr.Transaction(db)
		} else {
			fmt.Println("Parol xato!!")
		}
	} else {
		fmt.Println("Email xato!!")
	}

}
