package main

import (
	"fmt"
	"strings"
)


func baseType() {
	// fmt.Println("baseType")
	// 1. 要注意的是，Go 是强类型语言，因此不会进行隐式转换，任何不同类型之间的转换都必须显式说明.
	// Go 不存在像 C 那样的运算符重载，表达式的解析顺序是从左至右。

	// 2. Go 对于值之间的比较有非常严格的限制，只有两个类型相同的值才可以进行比较,
	// 如果值的类型是接口（interface），它们也必须都实现了相同的接口。如果其中一个值是常量，那么另外一个值的类型必须和该常量类型相兼容的。
	// 如果以上条件都不满足，则其中一个值的类型必须在被转换为和另外一个值的类型相同之后才可以进行比较。

	// 3. Go中不允许不同类型之间的混合使用，但是对于常量的类型限制非常少
	// fmt.Printf("%6.2e\n", 3.4)

	// 4. Println通常用来单独输出一行，含有默认格式化类型,如果需要指定格式化方式，可使用Printf
	// c := complex(20.0, 3.4)
	// fmt.Printf("%v %f\n", c, c)
	// fmt.Println(real(c), imag(c))

	// 5. 最小宽度, 不够部分可以选择补0, 最大宽度, 超出的部分会被截断
	fmt.Printf("|%05s|", "aa") // |000aa|
	fmt.Printf("|%.5s|\n", "xxxXXxx") // |xxxXX|

	// 6. 对于整数和浮点数，你可以使用一元运算符 ++（递增）和 --（递减），但只能用于后缀
	//    同时，带有 ++ 和 -- 的只能作为语句，而非表达式，因此 n = i++ 这种写法是无效的

	var str string = "ggg gg   gggg  ggg"
	fmt.Printf("%d\n", strings.Count(str, "gg"))

	res := strings.Fields(str)
	fmt.Println(res, len(res))
}

func main(){
	baseType()
}