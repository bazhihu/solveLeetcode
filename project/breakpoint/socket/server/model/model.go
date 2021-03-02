package model

import (
	"database/sql"
	"errors"
	"fmt"
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
	_, err := db.Exec(sql)
	return true, err
}

func (bin *Binlog) ExecuteSql(db *sql.DB, row []string) (bool, error) {
	l := len(row)
	var (
		b   bool  = false
		err error = nil
	)
	switch row[l-1] {
	case "I":
		b, err = bin.insert(db, row[0:l-1])
	case "D":
		b, err = bin.delete(db, row[0:l-1])
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
	//id, errI := res.LastInsertId()
	return n > 0, nil
}

func (bin *Binlog) delete(db *sql.DB, row []string) (bool, error) {
	return true, nil
}
