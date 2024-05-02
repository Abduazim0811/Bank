package signup

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	md "github.com/Abduazim0811/Bank/internal/models"
	// tr "github.com/Abduazim0811/Bank/internal/Transaction"
)

func Signup(db *sql.DB) {
	var (
		user md.Users
		num int
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

	_,err=db.Exec(string(sqlfile))
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("----Siz muvvaqiyatli kirdingiz!!!----")

	fmt.Println("[1]Transaction\t[2]Accaunt")
	fmt.Scanln(&num)
	switch num{
	case 1:
		
	case 2:

	}
}
