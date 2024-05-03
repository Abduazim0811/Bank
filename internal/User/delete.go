package user

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	md "github.com/Abduazim0811/Bank/internal/models"
)

func UserDelete(db *sql.DB,user md.Users){
	name:="../internal/DB/deleteuser.sql"
	sqlfile,err:=os.ReadFile(name)
	if err!=nil{
		log.Fatal(err)
	}
	_,err=db.Exec(string(sqlfile), user.Email)
	if err!=nil{
		log.Fatal(err)
	}
}

func Delete(db *sql.DB){
	var (
		num int
		user md.Users
	)
	fmt.Println("Siz akkauntni o'chirishni hohlaysizmi?")
	fmt.Println("[1]HA\t[2]YOQ")
	fmt.Scanln(&num)
	if num==1{
		UserDelete(db,user)
		fmt.Println("Akkauntingiz o'chirildi!!!")
	}else{
		return
	}
}
