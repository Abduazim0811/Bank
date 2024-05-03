package transaction

import (
	"database/sql"
	"fmt"
	"os"

	us "github.com/Abduazim0811/Bank/internal/User"
)

func Transaction(db *sql.DB) {
	var (
		num int
	)
	i := 1
	for i > 0 {
		fmt.Println(`
[1]Check account
[2]Edit account
[3]Put money
[4]Withdraw money
[5]Sending money
[6]Delete account
[7]Exit
	`)
	fmt.Print(">>>>>>: ")
		fmt.Scanln(&num)
		switch num {
		case 1:
			us.CheckAccount(db)
		case 2:
			us.Edit(db)
		case 3:
			PutMoney(db)
		case 4:
			WitdrawMoney(db)
		case 5:
			SendingMoney(db)
		case 6:
			us.Delete(db)
		case 7:
			os.Exit(0)
		default:
			fmt.Println("Wrong number!!!")
		}
	}

}
