// test 1
// package main

// var a = "G"

// func main() {
//    n()
//    m()
//    n()
// }

// func n() { print(a) }

// func m() {
//    a := "O"
//    print(a)
// }



// test 2
// package main

// var a = "G"

// func main() {
//    n()
//    m()
//    n()
// }

// func n() {
//    print(a)
// }

// func m() {
//    a = "O"
//    print(a)
// }



// test 3
package main

var a string

func main_ex() {
   a = "G"
   print(a)
   f1()
}

func f1() {
   a := "O"
   print(a)
   f2()
}

func f2() {
   print(a)
}