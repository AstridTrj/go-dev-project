package main

import (
	"fmt"
)


func main()  {
	fmt.Println("chapter8")

	// 1. 在声明的时候不需要知道map的长度，map是可以动态增长的.
	// 2. 通过key在map中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢100倍；所以如果你很在乎性能的话还是建议用切片来解决问题
	// 3. map 也可以用函数作为自己的值
	// 4.不要使用 new()，永远用 make() 来构造 map
	// 5. 注意:如果你错误地使用 new() 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址
	mapNew := new(map[string]int)
	// mapNew["k1"] = 4  // 错误使用
	(*mapNew) = make(map[string]int)
	(*mapNew)["k1"] = 4

	var mapDef map[string]float32
	// mapDef["k2"] = 3.447 // 错误：map使用前需先初始化，否则为nil

	fmt.Println(mapNew, mapDef)
}