package point

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"net"
	"os"
	"solveLeetcode/project/binlog/model"
	"strconv"
	"time"
)

func ServerConn(conn net.Conn, db *sql.DB) {
	var (
		filename string
		postfix  string
	)
	for {
		var buf = make([]byte, 10)
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
			off := getFileStat()

			stringoff := strconv.FormatInt(off, 10)
			_, err := conn.Write([]byte(stringoff))
			if err != nil {
				log.Fatalf("server write faild: %s\n", err)
			}
			continue
		case "<--end":
			// 如果接收到客户端同志所有文件内容发送完毕消息则退出
			log.Fatalf("receive over \n")
			return
		}
		writeFile(buf[:n])
	}

	// 处理数据
	var bin = model.Binlog{}

	switch postfix {
	case "sql":
		fileInfo, err := os.Stat(filename)
		if err != nil {
			log.Fatalf("err :%+v", err)
			return
		}
		log.Fatalf("size :%d", fileInfo.Size())
		var buf []byte

		fs, err := os.Open(filename)
		if err != nil {
			log.Fatalf("err :%+v", err)
			return
		}
		defer fs.Close()

		n, err := fs.Read(buf[:])
		if err == io.EOF || err != nil {
			log.Fatalf("read file err %+v", err)
			return
		}
		bin.CreateTable(db, string(buf[:n]))
	case "csv":
		go func() {
			for {
				time.Sleep(1 * time.Second)
				data, _ := model.CacheMap.Load(1)
				bin.ExecuteSql(db, l.(string))
			}
		}()
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
			bin.ExecuteSqlCache(db, row)
		}
	}

}

// 把接收到的内容 写入文件
func writeFile(content []byte) {
	if len(content) != 0 {
		fp, err := os.OpenFile("test_1.txt", (os.O_CREATE | os.O_WRONLY | os.O_APPEND), 0755)
		defer fp.Close()
		if err != nil {
			log.Fatal("open file faild: %s\n", err)
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
func getFileStat() int64 {
	fileInfo, err := os.Stat("test_1.txt")
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
