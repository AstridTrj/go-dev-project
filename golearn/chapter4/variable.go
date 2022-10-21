package main

import (
	"fmt"
	// "runtime"
	// "os"
)
/*
1. 当一个变量被声明之后，系统自动赋予它该类型的零值,int为 0，float32(64)为 0.0，bool为 false，string为空字符串，指针为nil。
所有的内存在 Go 中都是经过初始化的。
2. var a 这种语法是不正确的，因为编译器没有任何可以用于自动推断类型的依据
3. 值类型: int、float、bool、string、数组、struct，使用这些类型的变量直接指向存在内存中的值。
   引用类型：指针、slice、map、chan、接口、func，变量存储的是一个地址，这个地址存储最终的值。
4. 全局变量是允许声明但不使用
5. 如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a。
6. 变量除了可以在全局声明中初始化，也可以在 init() 函数中初始化。这是一类非常特殊的函数，它不能够被人为调用，
而是在每个包完成初始化后自动执行，并且执行优先级比 main() 函数高。每个源文件可以包含多个 init() 函数，
同一个源文件中的 init() 函数会按照从上到下的顺序执行，如果一个包有多个源文件包含 init() 函数的话，
则官方鼓励但不保证以文件名的顺序调用。初始化总是以单线程并且按照包的依赖关系顺序执行。
*/

var t int = 2

func main1(){
	fmt.Println("hello world")
	// 变量定义、声明与赋值
	// var goos string = runtime.GOOS
	// fmt.Printf("The operating system is: %s\n", goos)
	// path := os.Getenv("PATH")
	// fmt.Printf("Path is %s\n", path)

	// 指针
	var i int = 3
	var a, b *int
	a = &i
	b = a
	fmt.Println(&a, &b)

}