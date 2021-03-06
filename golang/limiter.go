package main

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"time"
)

/*
	限流器

*/

func main() {
	l := rate.NewLimiter(1, 5)
	log.Println("流速：", l.Limit(), "桶大小:", l.Burst())

	for i := 0; i < 100; i++ {
		// 阻塞等待直到，取到一个token
		log.Println("before wait")
		c, _ := context.WithTimeout(context.Background(), time.Second*2)
		if err := l.Wait(c); err != nil {
			log.Println("limiter wait err:" + err.Error())
		}

		log.Println("after wait")
		// 返回需要等待多久 才有新的token 这样可以等待指定时间执行任务

		r := l.Reserve()
		log.Println("reserve Delay:", r.Delay())

		// 判断当前是否可以取到token
		a := l.Allow()
		log.Println("allow:", a)
	}

}
