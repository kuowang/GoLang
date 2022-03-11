package main

import "fmt"
func main() {


	var p *string
	var a string
	a = "第一次写内容"
	p = &a
	fmt.Println(p)
	fmt.Printf("%T",p)



	var usernames []string = []string{"shineyork", "cara"} //切片
	fmt.Println(len(usernames))
	var username [5]string =[5]string{"diyige","第二个"}


	for i := 0;i < len(usernames);i++ {
		fmt.Println(usernames[i])
		fmt.Println(username[i])
	}
}
