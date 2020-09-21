package main

import "fmt"

/**
堆简单操作
堆化
插入
删除
*/

var heap []int // 堆
var max int    // 最大存储
var count int  // 当前储存

func init() {
	max = 40
	heap = make([]int, max)
}

// 从下往上 堆化
func heapifyOne() {
	i := count - 1
	for {
		if i/2 < 0 || heap[i] <= heap[i/2] {
			break
		}
		heap[i], heap[i/2] = heap[i/2], heap[i]
		i = i / 2
	}
	return
}

// 从上往下 堆化
func heapifyTwo() {
	i, maxPos := 0, 0
	for {
		if (i+1)*2-1 <= count-1 && heap[(i+1)*2-1] > heap[i] {
			maxPos = (i+1)*2 - 1
		}

		if (i+1)*2 <= count-1 && heap[(i+1)*2] > heap[maxPos] {
			maxPos = (i + 1) * 2
		}

		if i == maxPos {
			break
		}

		heap[i], heap[maxPos] = heap[maxPos], heap[i]
		i = maxPos
	}
	return
}

// 插入
func insert(i int) {
	if count >= max {
		return
	}
	count++
	// 直接插入到堆底，然后进行堆化
	heap[count-1] = i
	heapifyOne()
	return
}

// 删除
func delete() {
	if count <= 0 {
		return
	}
	result := heap[0]
	// 将最后一个元素，提到堆顶，然后堆化
	heap[0] = heap[count-1]
	heap[count-1] = 0
	count--
	heapifyTwo()
	fmt.Println("delete", result)
}

// 堆排序 将堆顶大数值 假装删除置换到数组末尾
func sorting() {
	for {
		if count <= 0 {
			break
		}
		heap[0], heap[count-1] = heap[count-1], heap[0]
		count--

		heapifyTwo()
	}
}

func main() {
	insert(1)
	insert(6)
	insert(5)
	insert(2)
	insert(3)
	insert(10)
	insert(7)
	insert(8)
	insert(10)

	fmt.Println("heap", heap, "count", count)

	//delete()
	//delete()
	//delete()
	//delete()

	sorting()

	fmt.Println("heap", heap, "count", count)
}
