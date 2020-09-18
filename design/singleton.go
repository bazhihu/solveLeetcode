/**
单例模式
*/
package design

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

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

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
