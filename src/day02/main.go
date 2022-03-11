package main_test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var inputReader = bufio.NewReader(os.Stdin)

func main() {
	oper := []string{"计算器", "if", "switch", "数组排序", "map"}

	// for循环
	for {
		fmt.Println("============>>>>> start ============>>>>>")
		fmt.Println("项目运行请输入你想看的操作")
		for okey, ovalue := range oper {
			fmt.Printf("(%d) : %s\n", okey, ovalue)
		}
		fmt.Println("(x) : 结束")
		fmt.Println("请输入你的操作")

		input, _ := inputReader.ReadString('\n') // 输入回车就结束
		flag, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("信息输入有误请重新输入")
			continue
		}
		choice(flag)
	}
}
func choice(choice int) {
	switch choice {
	case 0:
		// 计算器
	case 1:
		// if .....
	case 2:
		// switch
	case 3:
		// 数组排序
	case 4:
		// map -- key和valu转化
	}
}
