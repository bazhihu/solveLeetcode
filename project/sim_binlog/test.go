package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	sql := "CREATE TABLE `test` (" +
		"`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '编号'," +
		"`a` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'a'," +
		"`b` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'b'," +
		"`c` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'c'," +
		"`d` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'd'," +
		"`e` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'e'," +
		"`f` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'f'," +
		"PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试表';"
	tableNameReg := regexp.MustCompile("CREATE TABLE `(\\w+)`")
	strArr := tableNameReg.FindAllString(sql, 1)

	fmt.Println(strArr)

	fmt.Println(strings.Split(strArr[0], "`")[1])

	//fs, err := os.Open("data.csv")
	//if err != nil {
	//	log.Fatalf("err :%+v", err)
	//	return
	//}
	//defer fs.Close()
	//
	//r := csv.NewReader(fs)
	//for {
	//	row, err := r.Read()
	//	if err != nil && err != io.EOF {
	//		log.Fatalf("err: %+v", err)
	//	}
	//	if err == io.EOF {
	//		return
	//	}
	//	log.Println(row)
	//}
}
