package signup

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	md "github.com/Abduazim0811/Bank/internal/models"
	tr "github.com/Abduazim0811/Bank/internal/Transaction"
)

func Signup(db *sql.DB) {
	var (
		user md.Users
	)

	fmt.Print("Firstname: ")
	fmt.Scanln(&user.Firstname)
	fmt.Print("Lastname: ")
	fmt.Scanln(&user.Lastname)
	fmt.Print("Fathersname: ")
	fmt.Scanln(&user.Fathers_name)
	fmt.Print("Email: ")
	fmt.Scanln(&user.Email)
	fmt.Print("Password: ")
	fmt.Scanln(&user.Password)

	name:="../internal/DB/insertusers.sql"

	sqlfile,err:=os.ReadFile(name)
	if err!=nil{
		log.Fatal(err)
	}

	_,err=db.Exec(string(sqlfile), user.Firstname,user.Lastname,user.Fathers_name,user.Email,user.Password)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("----Siz muvvaqiyatli kirdingiz!!!----")
	tr.Transaction(db)
	
}
