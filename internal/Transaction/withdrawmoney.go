package transaction

import (
	"database/sql"
	"fmt"
)

func WitdrawMoney(db *sql.DB) {
	var (
		sum float64
	)
	fmt.Print("Qancha yechib olmoqchisiz?:  ")
	fmt.Scanln(&sum)

	
}
