package point

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"net"
	"os"
	"solveLeetcode/project/sim_binlog/model"
	"strconv"
	"strings"
)

func getData(conn net.Conn) (filename, postfix string) {
	for {
		var buf = make([]byte, 200)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("server io EOF")
				return
			}
			log.Fatalf("server read faild: %s\n", err)
		}

		log.Printf("recevice %d bytes, content is 【%s】\n", n, string(buf[:n]))

		switch string(buf[:n]) {
		case "start-->":
			_, err := conn.Write([]byte("filename"))
			if err != nil {
				log.Fatalf("server write faild: %s\n", err)
			}

			buf := make([]byte, 200)
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatalf("receive server info faild:%s \n", err)
			}

			filename = string(buf[:n])
			off := getFileStat(filename)
			postfix = strings.Split(filename, ".")[1]

			stringoff := strconv.FormatInt(off, 10)
			_, err = conn.Write([]byte(stringoff))
			if err != nil {
				log.Fatalf("server write faild: %s\n", err)
			}
			continue
		case "<--end":
			// 如果接收到客户端同志所有文件内容发送完毕消息则退出
			//log.Fatalf("receive over \n")
			log.Println("+++++++++++++")
			return
		default:
			writeFile(filename, buf[:n])
		}
	}
	log.Println("--------------")
	return
}

func ServerConn(conn net.Conn, db *sql.DB, bin *model.Binlog, ch chan<- string) {
	var (
		filename string
		postfix  string
	)
	defer conn.Close()
	filename, postfix = getData(conn)

	if postfix != "" {
		go func(filename, postfix string) {

			switch postfix {
			case "sql":
				createSql := readFile(filename)
				log.Println("|||||||||||||", createSql)
				_, err := bin.CreateTable(db, createSql)
				if err != nil {
					log.Println("sql----------CreateTable-err", err)
				}
			case "csv":
				fs, err := os.Open(filename)
				if err != nil {
					log.Fatalf("err :%+v", err)
					return
				}
				defer fs.Close()

				r := csv.NewReader(fs)
				for {
					row, err := r.Read()
					if err != nil && err != io.EOF {
						log.Fatalf("err: %+v", err)
					}
					if err == io.EOF {
						return
					}
					bin.ExecuteSqlCache(db, row, ch)
				}
			}
		}(filename, postfix)
	}
}

// 把接收到的内容 写入文件
func writeFile(filename string, content []byte) {
	if len(content) != 0 {
		fp, err := os.OpenFile(filename, (os.O_CREATE | os.O_WRONLY | os.O_APPEND), 0755)
		defer fp.Close()
		if err != nil {
			log.Fatalf("open file faild: %s : %s\n", err, filename)
		}
		_, err = fp.Write(content)
		if err != nil {
			log.Fatal("append content to file faild: %s\n", err)
		}
		log.Printf("append content: 【%s】 success\n", string(content))
	}
}

// 获取已接受的内容的大小
// 断点续传需要把已接收内容大小 通知客户端从哪儿开始发送文件内容
func getFileStat(fileName string) int64 {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("file size: %d\n", 0)
			return int64(0)
		}
		log.Fatalf("get file stat faild: %s\n", err)
	}
	log.Printf("file size: %d\n", fileInfo.Size())
	return fileInfo.Size()
}

// 一次性读取文件
func readFile(filename string) string {

	fs, err := os.Open(filename)
	if err != nil {
		log.Fatalf("err :%+v", err)
		return ""
	}
	defer fs.Close()
	var chunk []byte
	buf := make([]byte, 1024)

	for {
		n, err := fs.Read(buf[:])
		if err != nil && err != io.EOF {
			log.Fatalf("read file err %+v", err)
			return ""
		}
		if n == 0 {
			break
		}
		chunk = append(chunk, buf[:n]...)
	}

	return string(chunk)
}
