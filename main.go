package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 入口启动函数
func main() {
	// 调用方法
	// 创建一个读取器来读取命令行输入
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入对应权限点 (1: 方法 1, 2: 方法 2, 3: 方法 3, q: 退出): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("读取输入时出错: %v", err)
		}
		// 去除输入末尾的换行符
		input = input[:len(input)-1]

		switch input {
		case "1":
			method1()
		case "2":
			method2()
		case "3":
			method3()
		case "q":
			fmt.Println("退出程序")
			return
		default:
			fmt.Println("无效的命令，请重新输入")
		}
	}
}

// 方法 1：接收用户输入的值
// 方法 1：接收用户输入的值
// 方法 1：输入 client_id 和 access_token
func method1() {
	reader := bufio.NewReader(os.Stdin)

	// 输入 client_id
	fmt.Print("请输入 client_id: ")
	clientID, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	clientID = clientID[:len(clientID)-1]

	// 输入 access_token
	fmt.Print("请输入 access_token: ")
	accessToken, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	accessToken = accessToken[:len(accessToken)-1]

	// 使用输入的值执行逻辑
	fmt.Printf("方法 1 执行成功！\nclient_id: %s\naccess_token: %s\n", clientID, accessToken)
}

func method2() {
	fmt.Println("调用了方法 2")
}

func method3() {
	fmt.Println("调用了方法 3")
}
