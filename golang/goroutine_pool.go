package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

/*
	本质： 生产者消费者模型
	解决：可以有效控制goroutine 数量，防止暴涨
	缺点：无法保证顺序 -- 解决方法：串行 + 锁


	例子：
		1、计算一个数字的各个位数之和，例如 数字123 结果为 1+2+3 = 6
		2、随机生成数字进行计算

*/

// 目标
type Job struct {
	Id string
	// 需要计算的随机数
	RandNum int
}

// 结果
type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

func main() {
	// 需要两个管道
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)

	// 工作池
	createPool(64, jobChan, resultChan)

	// 开启消费打印
	go func(resultChan chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)

	go func() {
		// 放入随机数 job
		for {
			r_num := rand.Int()
			job := &Job{
				Id:      rand.Int(),
				RandNum: r_num,
			}
			jobChan <- job
		}
	}()

	var errChan = make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	log.Fatalf("err %v\n", <-errChan)
}

// 创建工作池
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据协程 个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum

				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}

				resultChan <- &Result{
					job: job,
					sum: sum,
				}
			}
		}(jobChan, resultChan)
	}
}
