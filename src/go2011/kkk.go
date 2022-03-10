package main

import ( // 后面补充  go mod

	"fmt"
	"go2011/srm" // $gopath/go2011/srm

	sms "github.com/shineyork/go-sms" // $gopath/github.com/shineyork/go-sms
)

var class string

// 是在main方法之前执行

/*
init 就相当于是一个构造函数 - 》针对go的包而显示的
作用：是用来初始化变量，参数等等；

1. init 可以定义多个 main只能是一个
2. 执行顺序：
		根据引入的包的顺序执行 init
		在包目录下的.go 中的init；根据文件名的顺序执行
*/
func init() {

	fmt.Println("main")
	class = "p"
}

// main只能有一个
func main() {

	srm.Index()
	sms.Email()

	// 计算器 结构体；-> 以及一系列的方法
	var i int
	fmt.Scan(&i) // 获取到控制台输入的信息
	fmt.Println("获取到的内容", i)

	// fmt.Scan(&i)

	// srm.user
	// tu := srm.User{}
	// tu.Name = "p"
	// // tu.age = 9
	// // tu.tostring()
	// fmt.Println(tu)
	// 创建结构体
	// 实例化
	// tu := user{}
	// tu.name = "p"
	// tu.age = 9
	// tu.tostring()
	// fmt.Println(tu.getAge())

	// var k string = "0"
	// k = "p"
	// fmt.Println(k)

	// psds := "0" // 在局部变量中运用 ， 定义类型根据值的类型赋值； 会隐士进行类型转化
	// fmt.Println(psds)
	// // psds = "0ooo"
	// psds = "sdsd"
	// fmt.Println(psds)
	// i := 0
	// fmt.Println(i)
	//
	// fmt.Println(U)
}
