package test

import (
	"errors"
	"fmt"
	"sync"
)

// 交替打印
func supersede() {
	// 协程交叉打印
	var printChar = func(chanA <-chan struct{}, chanB chan<- struct{}) {
		for i := 0; i < 11; i++ {
			j := 65 + i
			select {
			case <-chanA:
				fmt.Printf("%s", string(rune(j)))
				chanB <- struct{}{}
				break
			}
		}
	}
	var printNum = func(chanB <-chan struct{}, chanA chan<- struct{}, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			j := 1 + i
			select {
			case <-chanB:
				fmt.Printf("%d", j)
				chanA <- struct{}{}
				break
			}
		}
		wg.Done()
	}
	var (
		wg           = &sync.WaitGroup{}
		chanA, chanB = make(chan struct{}), make(chan struct{})
	)

	defer func() {
		close(chanA)
		close(chanB)
	}()

	wg.Add(1)

	go printChar(chanA, chanB)
	go printNum(chanB, chanA, wg)
	chanA <- struct{}{}
	wg.Wait()
}

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

type User struct {
	id   int
	name string
}

func (self *User) ToString() string {
	return fmt.Sprintf("User:%p,%v", self, self)
}
