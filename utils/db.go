package utils

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Conn_DB() *sql.DB { // return mysql connection
	user := "netadmin_s96lu"
	password := "netadmin_s96lu"
	_db := "socialnet"
	_ip := "192.168.157.20"

	db, _ := sql.Open("mysql", user+":"+password+"@tcp("+_ip+":3306)/"+_db)
	err := db.Ping()
	if err != nil {panic(err)}
	return db
}
