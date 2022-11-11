package main

import (
	"fmt"
	"io"
	"log"
)


func f(int, int) (x int) {
	fmt.Println("没有形参名字")
	x = 4
	return 
}

func deferTest() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func multiDefer() {
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func inter(s string) string{
	fmt.Println("inter:", s)
	return s
}

func b() {
	defer un(trace(inter("b")))
	fmt.Println("in b")
	a()
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func main(){
	fmt.Println("function: ", f(2, 2));
	// 1.函数重载指可以编写多个同名函数，只要它们拥有不同的形参或者不同的返回值，在Go里面函数重载是不被允许的
	// 2.在函数调用时，像切片(slice)、字典(map)、接口(interface)、通道(channel)这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）。
	// 3.当函数需要返回多个非命名返回值时，需要使用 () 把它们括起来，一个的情况可以不使用括号
	// 4.命名返回值作为结果形参被初始化为相应类型的零值，当需要返回的时候，我们只需要一条简单的不带参数的return语句。
	// 	 需要注意的是，即使只有一个命名返回值，也需要使用 () 括起来
	// 5.虽然 return 或 return var 都是可以的，但是 return var = expression（表达式）会引发一个编译错误
	// 6.变长参数可以作为对应类型的 slice 进行二次传递
	// 7.如果一个变长参数的类型没有被指定，则可以使用默认的空接口interface{}，这样就可以接受任何类型的参数。
	// 	 该方案不仅可以用于长度未知的参数，还可以用于任何不确定类型的参数。
	deferTest()
	// 8.当有多个defer行为被注册时，它们会以逆序执行（类似栈，即后进先出）
	multiDefer()
	// 注意defer后执行的嵌套函数，defer只跟踪了最外层函数
	b()

	func1("Go")
}