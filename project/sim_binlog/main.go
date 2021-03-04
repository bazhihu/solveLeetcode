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

func main() {
	var (
		host   = flag.String("slave.host", "127.0.0.1", "host")
		port   = flag.String("slave.port", "3306", "port")
		user   = flag.String("slave.user", "root", "user")
		pwd    = flag.String("slave.password", "123456", "password")
		dbName = flag.String("slave.db", "test", "dbname")
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
			defer conn.Close()
			point.ServerConn(conn, db)
		}
	}()

	go func() {
		// 缓存数据异步入库
		var bin = model.Binlog{}
		for {
			time.Sleep(1 * time.Second)
			keys := make([]interface{}, 0)
			model.CacheMap.Range(func(key, value interface{}) bool {
				var sql string
				switch value.(type) {
				case string:
					sql = value.(string)
				}
				if len(sql) > 0 {
					keys = append(keys, key)
					bin.ExecuteSql(db, sql)
				}
				return true
			})
			for k, _ := range keys {
				model.CacheMap.Delete(keys[k])
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
