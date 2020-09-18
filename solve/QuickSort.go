package solve

import "fmt"

/**
快速排序 - 递归
*/
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	splitdata := arr[0]        // 暂定临界值为第一个数据
	low := make([]int, 0, 0)   // 小数据
	hight := make([]int, 0, 0) // 大数据
	mid := make([]int, 0, 0)   // 相等数据
	mid = append(mid, splitdata)

	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}

	low, hight = QuickSort(low), QuickSort(hight)
	return append(append(low, mid...), hight...)
}

/**
快速排序 - 非递归 将分治的数组押入栈中 循环调用
*/
func QuickSortNotR(arr []int) []int {
	length := len(arr)
	var stack Stack
	stack.Push(0, length-1)

	for {
		i := stack.Pull()
		if i == nil {
			break
		}

		if i.Left > i.Right {
			continue
		}

		m := i.Left
		n := i.Right

		key := arr[(i.Left+i.Right)/2]

		for m <= n {
			for arr[m] < key {
				m++
			}
			for arr[n] > key {
				n--
			}
			if m <= n {
				arr[m], arr[n] = arr[n], arr[m]
				m++
				n--
			}
		}
		if i.Left < n {
			stack.Push(i.Left, n)
		}
		if i.Right > m {
			stack.Push(m, i.Right)
		}
	}

	return arr
}

// 栈内元素结构
type Index struct {
	Left  int
	Right int
}

// 定义栈
type Stack struct {
	S []*Index
}

// 入栈
func (s *Stack) Push(l, r int) {
	s.S = append(s.S, &Index{
		Left:  l,
		Right: r,
	})
}

// 出栈
func (s *Stack) Pull() *Index {
	if len(s.S) == 0 {
		return nil
	}
	rt := s.S[0]
	s.S = s.S[1:]
	return rt
}

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	arrNew := QuickSort(arr)
	fmt.Println(arrNew)
}
