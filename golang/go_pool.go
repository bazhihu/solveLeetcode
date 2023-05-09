package main

import (
	"context"
	"fmt"
	"math/rand"
)

// 本质是 生产者消费者模型
// 控制 goroutine数量，防止暴涨

type Job struct {
	//id
	Id int
	// 需要计算得随机数
	RandNum int
}

type Result struct {
	// 对象实列
	job *Job
	//求和
	sum int
}

func call(num int) int {
	// 随机数每一位相加
	var sum int
	for num != 0 {
		tmp := num % 10
		sum += tmp
		num /= 10
	}
	return sum
}

// 创建工作池
// 参数1： 开几个协程
func createPool(ctx context.Context, num int, jobChan chan *Job, resultChan chan *Result, call func(int) int) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(ctx context.Context, jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历Job 管道所有数据，进行相加
			for {
				select {
				case <-ctx.Done():
					return
				case job := <-jobChan:
					r_num := job.RandNum
					sum := call(r_num)
					r := &Result{
						job: job,
						sum: sum,
					}
					resultChan <- r
				}
			}
		}(ctx, jobChan, resultChan)
	}
}

func main() {
	// 两个管道 一个job管道，一个结果管道
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 创建工作池
	createPool(ctx, 50, jobChan, resultChan, call)

	// 开启打印协程
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("JOB ID:%V RANDNUM:%V RESULT:%D\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)

	// 生产者
	// 循环创建job，输入到管道
	var id int
	for {
		id++
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}

}
