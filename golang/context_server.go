package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var banner = `
____________________________________O/_______
                                    O\`

// 服务器端接受客户端请求，随机延迟响应

func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)

	fmt.Println("number:", number)
	if number == 0 {
		time.Sleep(time.Second * 10) //故意制造延迟响应
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprint(w, "quick response")
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println(banner)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
