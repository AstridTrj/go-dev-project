package main


import (
	"fmt"
)

// 类型和作用在它上面定义的方法必须在同一个包里定义，这就是为什么不能在 int、float32(64) 或类似这些的类型上定义方法
// func (p *list.List) Iter() {
// 	// ...
// }

type Test struct{
	a int
	b float32
}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main(){
	fmt.Println("chapter 10")
	// 1. t := new(T)，变量 t 是一个指向 T 的指针, 声明 var t T 也会给 t 分配内存，并零值化内存，但是这个时候 t 是类型 T
	var t Test
	fmt.Println(t) // {0 0}
	t1 := new(Test)
	t1.a = 3
	// 2. 在Go语言中这叫选择器。无论变量是一个结构体类型还是一个结构体类型指针，都使用同样的选择器符来引用结构体的字段,如t.a与t1.a等效,Go 会自动做转换
	fmt.Println(t1, t1.a) // &{0 0}
	// 3. 可以通过在值的前面放上字段名来初始化字段
	t2 := Test{b: 5, a: 3}
	fmt.Println(t2)
	// 4. Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体
	// 5. 试图 make() 一个结构体变量，会引发一个编译错误
	// error: invalid argument: cannot make Test; type must be slice, map, or channel
	// tMake := make(Test)
	// 6. 结构体中的字段除了有名字和类型外，还可以有一个可选的标签 (tag)：它是一个附属于字段的字符串，
	//     可以是文档或其他的重要标记。标签的内容不可以在一般的编程中使用，只有包 reflect 能获取它
	// 7. 在一个结构体中对于每一种数据类型只能有一个匿名字段


	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()

	// 8. Go 开发者不需要写代码来释放程序中不再使用的变量和结构占用的内存，在 Go 运行时中有一个独立的进程，即垃圾收集器 (GC)，
	//     会处理这些事情，它搜索不再使用的变量然后释放它们的内存。可以通过 runtime 包访问 GC 进程。
}