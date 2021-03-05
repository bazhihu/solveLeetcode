package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

/*
	硬编码更新mysql
	1、 建表语句

	2、插入删除sql
*/

type syncBinlog interface {
	CreateTable(db *sql.DB, sql string) (bool, error)
	ExecuteSql(db *sql.DB, row []string) (bool, error)
	insert(db *sql.DB, row []string) (bool, error)
	delete(db *sql.DB, row []string) (bool, error)
}

type Binlog struct {
	TableName string
}

func (bin *Binlog) CreateTable(db *sql.DB, sql string) (bool, error) {
	if len(sql) < 1 {
		return false, nil
	}
	tableNameReg := regexp.MustCompile("CREATE TABLE `(\\w+)`")
	strArr := tableNameReg.FindAllString(sql, 1)
	if len(strArr) < 1 {
		return false, nil
	}
	// set tableName
	bin.TableName = strings.Split(strArr[0], "`")[1]

	log.Println("||||||||||||||||sql", sql, bin.TableName)
	_, err := db.Exec(sql)
	return true, err
}

func (bin *Binlog) ExecuteSqlCache(db *sql.DB, row []string, Ch chan<- string) (bool, error) {
	r := strings.Join(row, ",")
	CacheMap.Store(r, r)
	Ch <- r
	return true, nil
}

func (bin *Binlog) ExecuteSql(db *sql.DB, row string) (bool, error) {
	rows := strings.Split(row, ",")
	l := len(rows)
	var (
		b   bool  = false
		err error = nil
	)

	log.Println("||||||||_________ tablename", bin.TableName)

	switch rows[l-1] {
	case "I":
		b, err = bin.insert(db, rows[0:l-1])
	case "D":
		b, err = bin.delete(db, rows[0:l-1])
	}
	return b, err
}

func (bin *Binlog) insert(db *sql.DB, row []string) (bool, error) {
	if len(row) < 1 {
		return false, errors.New("no value")
	}

	var (
		plate  = make([]string, len(row))
		params = make([]interface{}, len(row))
	)

	for k, _ := range row {
		plate[k] = "?"
		params[k] = row[k]
	}

	var sqlText = fmt.Sprintf("INSERT INTO %s VALUES (%s)", bin.TableName, strings.Join(plate, ","))
	stmt, err := db.Prepare(sqlText)
	if nil != err {
		return false, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(params...)
	if nil != err {
		return false, err
	}
	n, err := res.RowsAffected()
	if nil != err {
		return false, err
	}
	log.Println("||||||||||||||||sqlText", sqlText)
	//id, errI := res.LastInsertId()
	return n > 0, nil
}

func (bin *Binlog) delete(db *sql.DB, row []string) (bool, error) {
	// default id is row[0]
	if len(row) < 1 {
		return false, nil
	}

	var sqlText = fmt.Sprintf("DELETE FROM %s WHERE id = ?", bin.TableName)
	stmt, err := db.Prepare(sqlText)
	if nil != err {
		return false, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(row[0])
	if nil != err {
		return false, err
	}
	n, err := res.RowsAffected()
	if nil != err {
		return false, err
	}
	log.Println("||||||||||||||||sqlText", sqlText)
	return n > 0, nil
}
