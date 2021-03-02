package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
func initDB(user, pwd, host, port, dbName  string) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s)", user, pwd, host, port, dbName)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func main()  {
	var (
		host = flag.String("slave.host", "127.0.0.1", "host")
		port = flag.String("slave.port", "3306", "port")
		user = flag.String("slave.user", "root", "user")
		pwd = flag.String("slave.password", "123456", "password")
		dbName = flag.String("slave.db", "test", "dbname")
	)
	flag.Parse()


}

