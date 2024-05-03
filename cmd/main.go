package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	signin "github.com/Abduazim0811/Bank/internal/Signin"
	signup "github.com/Abduazim0811/Bank/internal/Signup"

	_ "github.com/lib/pq"
)

func SelectUser(db *sql.DB) map[string]string {
	var us map[string]string = make(map[string]string)
	name := "../internal/DB/selectuser.sql"
	sqlfile, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(string(sqlfile))
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var email, password string

		err = rows.Scan(&email, &password)
		if err != nil {
			log.Fatal(err)
		}
		us[email] = password
	}

	return us
}

func main() {
	var (
		num int
	)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Abdu0811", "project")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	defer db.Close()

	name := "../internal/DB/createuser.sql"
	sqlfile, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(sqlfile))
	if err != nil {
		log.Fatal(err)
	}

	Us := SelectUser(db)

	fmt.Println("[1]Users\t[2]Admin")
	fmt.Scanln(&num)
	switch num {
	case 1:
		fmt.Println("[1]Signin\t[2]Signup\t[3]Exit")
		fmt.Scanln(&num)
		if num == 1 {
			signin.Signin(db, Us)
		} else if num == 2 {
			signup.Signup(db)
		} else {
			os.Exit(0)
		}

	case 2:
		fmt.Println("Tez orada admin panel ishlaydi!!!")
		// admin.Admin()
	}

}
