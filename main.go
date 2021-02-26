package main

import (
	"fmt"
	"gitee.com/ikongjix/go_common/redis_db/master_db"
	"time"

	//"gitee.com/ikongjix/go_common/redis_db/master_db"
	"log"
	"net/http"
	"regexp"
)

const POST_CENTER_THREAD_GO_REDIS_CACHE = "pc:tg:" // 帖子redis 缓存

func a() {
	master_db.Lock("123123", 10)
	defer master_db.Unlock("123123")
	time.Sleep(5 * time.Second)
	fmt.Println("11111111")
}

func b() {
	for {
		if ok, _ := master_db.Lock("123123", 10); ok {
			break
		}
		time.Sleep(time.Second)
	}
	defer master_db.Unlock("123123")

	fmt.Println("2222222")
}

type ab struct {
}

func (a *ab) start() {
	go func() {
		c := 1
		for {
			log.Println(c)
			c++
		}
	}()
}

func main() {
	//go a()
	//time.Sleep(1*time.Second)
	//go b()
	//
	//time.Sleep(10*time.Second)
	//str := "[video]http://img.ikongji.com/attachment/newbie/post/202012/29/566337-1609224454180-2988.mp4[/video]sda[video]http://img.ikongji.com/attachment/newbie/post/202012/29/566337-1609224454180-2988.mp4[/video]"
	//data := regexp.MustCompile(`\[video\].*?\[\/video]`).FindAll([]byte(str), -1)
	//for k, _ := range data {
	//	fmt.Println(string(data[k]))
	//}

	//a := &ab{}
	//a.start()

	m := time.Now()
	time.Sleep(1 * time.Second)
	fmt.Println(time.Since(m))

	//var d sync.WaitGroup
	//d.Add(1)
	//go func() {
	//	defer d.Done()
	//	n := 0
	//	for {
	//
	//		n++
	//		log.Println("asda", n)
	//	}
	//}()
	//
	//go func() {
	//	defer d.Done()
	//	n := 0
	//	for {
	//		//runtime.Gosched()
	//		n++
	//		log.Println(n)
	//	}
	//}()
	//
	//d.Wait()
	//return

	//var a = "asdfghjkl"
	//
	//md5UrlByte := sha1.Sum([]byte(a + "post"))
	//md5Url := base64.StdEncoding.EncodeToString(md5UrlByte[:])
	//
	//fmt.Println(md5Url)
	//aa := fmt.Sprintf("%s%d", POST_CENTER_THREAD_GO_REDIS_CACHE, 77)
	//fmt.Println(aa)

	videoReg := regexp.MustCompile(`\[video\](.*?)\[\/video\]`)

	result := videoReg.FindAll([]byte("[video]http://img.ikongji.com/attachment/newbie/post/202101/13/742547-1610508280820-6853.JPEG[/video]"), -1)

	for k, _ := range result {
		fmt.Println(string(result[k]))
	}

	return

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "pong")
	})

	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "pong")
	})

	http.HandleFunc("hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello")
	})

	err := http.ListenAndServe(":9990", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
