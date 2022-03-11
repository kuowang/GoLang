package main_test

import (
	"fmt"
	// _ "go2011/day01/srm"
)

const (
	s = iota
	l = 100
	i = iota
	k = iota
)

func test() {
	// var a int8 = -128        // -128 ~ 127
	// fmt.Printf("(%d) \n", a) // (-128)

	var kh string = "0"
	fmt.Println("&kh", &kh)
	f := &kh
	*f = "p"
	fmt.Println(*f)
	fmt.Println(kh)
	fmt.Println(f)
	fmt.Println(&f)

	// fk := kh
	// fmt.Println(fk)

	u := 9
	sum(&u)
	fmt.Println(u)
	// a = a - 1
	// fmt.Println(a)
	//
	// var a1 int8 = 1
	// var a2 int = 1
	// fmt.Println(a1 == int8(a2))
	//
	// // uint8 == byte
	// var u8 uint8 = 1
	// var b1 byte = 1
	// fmt.Println(u8 == b1)

	// name := "   kk   "
	// fmt.Println(strings.Trim(name, " "))
	//
	// // string与int
	// // _ 代表省略
	// // 对import使用只会引入包下面的 init初始化可以不引用调用
	// sum1 := "1" // mysql , 参数中获取
	// num1, _ := strconv.ParseInt(sum1, 10, 64)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// fmt.Println(num1 + 1)
	//
	// fmt.Println(string(100))
	// fmt.Println()
}

func sum(a *int) {
	*a = *a + 10
}
