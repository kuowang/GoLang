package srm

import "fmt"

var k string // 全局使用正常声明

const U = "p"

var (
	n   string = "shineyork"
	inn int    = 00
)

const (
	kl = "sd"
	ck = "s"
)

// type fdao func() // 定义方法的类型
type ih interface { // 定义接口

}
type User struct { // 定义结构体
	Name string //结构体的属性
	age  int
}

func init() {
	fmt.Println("srm class .go")
}

func init() {
	fmt.Println("srm class .go 222")
}

func init() {
	fmt.Println("srm class .go 222222")
}

func (u User) tostring() {
	fmt.Println(u)
}

// 方法 根据 结构体而定义的
// 调用是一定要创建了结构体在.使用  实例化
func (u User) getAge() int {
	return u.age
}

// 函数
func hello() string {
	fmt.Println("hello world")
	return "p`"
}
