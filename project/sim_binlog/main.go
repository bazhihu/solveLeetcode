package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"os"
	"os/signal"
	"solveLeetcode/project/sim_binlog/model"
	"solveLeetcode/project/sim_binlog/point"
	"syscall"
	"time"
)

func initDB(user, pwd, host, port, dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pwd, host, port, dbName))
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, err
}

var (
	Ch  = make(chan string, 10)
	bin = model.Binlog{}
)

func main() {
	var (
		host   = flag.String("slave.host", "127.0.0.1", "host")
		port   = flag.String("slave.port", "3306", "port")
		user   = flag.String("slave.user", "root", "user")
		pwd    = flag.String("slave.password", "123456", "password")
		dbName = flag.String("slave.db", "herman", "dbname")
	)
	flag.Parse()

	db, err := initDB(*user, *pwd, *host, *port, *dbName)
	if err != nil {
		log.Fatalf("initDB error %d\n", err)
	}

	// 建立监听
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("error listen")
	}

	defer l.Close()

	log.Println("waiting accept.")

	var errChan = make(chan error)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatalf("accept faild: %s\n", err)
				errChan <- err
			}
			point.ServerConn(conn, db, &bin, Ch)
		}
	}()

	go func() {
		// 缓存数据异步入库
		for {
			key := <-Ch
			if value, ok := model.CacheMap.Load(key); ok {
				var sql string
				switch value.(type) {
				case string:
					sql = value.(string)
				}
				if len(sql) > 0 {
					_, err := bin.ExecuteSql(db, sql)
					if err != nil {
						log.Println("sql----------err", err)
					}
					model.CacheMap.Delete(key)
				}
			}
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	log.Fatalf("server out err: %d", <-errChan)
}
