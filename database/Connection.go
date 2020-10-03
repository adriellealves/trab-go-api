package database

import(
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)
func SetConnection() *sql.DB {
	connectionDB, err := sql.Open("mysql", "root:@/go_api")

	if err != nil {
		fmt.Println("Error to connect", err)
	}

	return connectionDB
}