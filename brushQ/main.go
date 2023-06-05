// @Title 刷题
// @Description
// @Author  herman  2023/3/9 20:48
// @Update  herman  2023/3/9 20:48
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// ad
func Q1() {

}

// 交替打印
func supersede() {
	// 协程交叉打印
	var printChar = func(chanA <-chan struct{}, chanB chan<- struct{}) {
		for i := 0; i < 11; i++ {
			j := 65 + i
			select {
			case <-chanA:
				fmt.Printf("%s", string(j))
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

func main() {
	fmt.Println("hello world")

	var s = "asadad"
	sR := []rune(s)
	//b := strings.Split(s, "")
	//fmt.Println(b)

	for i := 0; i < len(sR); i++ {
		if sR[i] == 97 {
			fmt.Printf("%d\n", sR[i])
		}
		//fmt.Printf("rrr/d", sR[i])
	}

	// 字符串和数字转换
	var i int
	i = 65544
	b := strconv.FormatInt(int64(i), 10)
	fmt.Printf("%s\n", reflect.TypeOf(b))

	ii, _ := strconv.Atoi(b)
	fmt.Printf("%d\n", ii)

	iis, _ := strconv.ParseInt(b, 10, 32)
	fmt.Printf("%d\n", iis)

	iif, _ := strconv.ParseFloat(b, 32)
	fmt.Printf("%f\n", iif)

	// 交替打印
	supersede()
	fmt.Println("\n")
	var a string = `2,Tina,37,"足球，""篮球",old
	3,adsas,37,"足球，""篮球",old
	`
	a = strings.ReplaceAll(a, "\",", "")
	a = strings.ReplaceAll(a, "\"\"", "\"")
	a = strings.ReplaceAll(a, "\n", "\t")
	fmt.Printf("%s", a)

	// regexp.MatchString("")

	var sss = "abcdefg"

	for i := range sss {
		fmt.Println(sss[i])
	}

	fmt.Println(sss[:3])
}
