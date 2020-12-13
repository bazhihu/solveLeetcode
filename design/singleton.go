/**
单例模式
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type Singleton struct {
	a int64 `json:"a"`
}

func (A *Singleton) echo() {
	A.a++
	fmt.Println(A.a)
}

var singleton *Singleton
var once sync.Once

// 保证只执行一次
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}

// 双重锁 避免多次加锁
var mu sync.Mutex

func GetInstanceTwo() *Singleton {
	if singleton == nil {
		mu.Lock()
		defer mu.Unlock()
		if singleton == nil {
			singleton = &Singleton{}
		}
	}
	return singleton
}

func main() {
	for i := 0; i <= 10; i++ {
		go func() {
			sing := GetInstance()
			sing.echo()
		}()
	}

	time.Sleep(3 * time.Second)
}
