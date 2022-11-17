package main

import (
	"fmt"
)

func main() {
	fmt.Println("main")
	// 1.如果我们想让数组元素类型为任意类型的话可以使用空接口作为类型,当使用值时我们必须先做一个类型判断.
	// 2.Go 语言中的数组是一种值类型
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	// 3.切片 (slice) 是对数组一个连续片段的引用（该数组我们称之为相关数组，通常是匿名的），所以切片是一个引用类型
	// 4.一个切片和相关数组的其他切片是共享存储的
	// 5.绝对不要用指针指向切片。切片本身已经是一个引用类型，所以它本身就是一个指针

	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	fmt.Println(items)
	// 6.改变切片长度的过程称之为切片重组 reslicing
	// 将切片扩展 1 位可以这么做：sl = sl[0:len(sl)+1]


	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo, slFrom)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

	var s string = "hello world"
	st := s[2:5]

	c := []byte(s)
	c[3] = 'p'
	s = string(c)
	
	fmt.Println(s, st)
	// 7. 只有在没有任何切片指向的时候，底层的数组内存才会被释放
}
