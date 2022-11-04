package main

import (
	"fmt"
	"math"
)

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 { return }
	return math.Sqrt(f), true
}

func main()  {
	// 1. if-else测试
	if t, ok := mySqrt(-25.0); !ok {
		fmt.Println(t, ok)
	}

	// 2. switch测试
	var num1 int = 20
	// default可放在任何位置
	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100: 
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	case 20:
		fmt.Println("It's equal to 20")
	}
	
	// 3. for 测试
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
	fmt.Println()

	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 { break }
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")
}