package main

import (
	"fmt"
	"gitee.com/ikongjix/go_common/tools/time_tools"
	"time"
)

// 时间向下取整
func timeIsRoundedDown() {
	timeNow := time_tools.GetUnix()
	getKey := func(now int64) string {
		remainder := now % 5
		return time.Unix(now-remainder, 0).Format("20060102:150405")
	}

	key := getKey(timeNow)
	fmt.Println(key)
	fmt.Println(time.Unix(timeNow, 0).Format("20060102:150405"))
}

func main() {

}
