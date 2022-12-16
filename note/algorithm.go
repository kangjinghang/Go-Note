package note

import (
	"fmt"
	"math/rand"
	"time"
)

// 7.1 递归
var fibonacciRes []int

func fibonacci(n int) int {
	if n < 3 {
		return 1
	}
	if fibonacciRes[n] == 0 {
		fibonacciRes[n] = fibonacci(n-2) + fibonacci(n-1)
	}
	return fibonacciRes[n]
}

func Recursion() {
	fmt.Println()
	fmt.Println("\n7.1 递归")
	n := 45
	fibonacciRes = make([]int, n+1)
	fmt.Printf("第%v位斐波那契数是%v\n", n, fibonacci(n))
}

// 7.2 闭包
func closureFunc() func(n int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次调用接收到n=%v\n", n)
		i++
		fmt.Printf("匿名工具函数被第%v次调用\n", i)
		return i
	}
}

func Closure() {
	fmt.Println()
	fmt.Println("\n7.2 闭包")
	f := closureFunc()
	f(2)
	f(4)
	f = closureFunc() // 闭包里的i重置成0
	f(6)
}

// 7.3 排序
// 7.3.1 冒泡排序
func bubbleSort(s []int) {
	lastIndex := len(s) - 1
	for i := 0; i < lastIndex; i++ {
		for j := 0; j < lastIndex-i; j++ {
			if s[j] > s[j+1] {
				t := s[j]
				s[j] = s[j+1]
				s[j+1] = t
			}
		}
	}
}

// 7.3.2 选择排序
func selectionSort(s []int) {
	lastIndex := len(s) - 1
	for i := 0; i < lastIndex; i++ {
		max := lastIndex - i
		for j := 0; j < lastIndex-i; j++ {
			if s[j] > s[max] {
				max = j
			}
		}
		if max != lastIndex-i {
			t := s[lastIndex-i]
			s[lastIndex-i] = s[max]
			s[max] = t
		}
	}
}

// 7.3.3 插入排序
func insertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		t := s[i]
		j := i - 1
		for ; j >= 0 && s[j] > t; j-- {
			s[j+1] = s[j]
		}
		if j != i-1 {
			s[j+1] = t
		}
	}
}

// 7.3.4 快速排序
func quickSort(s []int, leftIndex, rightIndex int) {
	if leftIndex < rightIndex {
		pivot := s[rightIndex]
		var rs []int
		l := leftIndex
		for i := leftIndex; i < rightIndex; i++ {
			if s[i] > pivot {
				rs = append(rs, s[i])
			} else {
				s[l] = s[i]
				l++
			}
		}
		s[l] = pivot
		copy(s[l+1:], rs)
		if leftIndex < l-1 {
			quickSort(s, leftIndex, l-1)
		}
		if l+1 < rightIndex {
			quickSort(s, l+1, rightIndex)
		}
	}
}

func Sort() {
	n := 20
	s := make([]int, n)
	seedNum := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		rand.Seed(seedNum)
		s[i] = rand.Intn(10001)
		seedNum++
	}
	fmt.Println("排序前:", s)
	// bubbleSort(s)
	// selectionSort(s)
	insertionSort(s)
	fmt.Println("排序后:", s)
}
