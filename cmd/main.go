package main

import (
	"bufio"
	"fmt"
	"log"
	"openplat/service/user"
	"os"
)

// 入口启动函数
func main() {
	// 调用方法
	// 创建一个读取器来读取命令行输入
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

	// 输入 app_secret
	fmt.Print("请输入 app_secret: ")
	appSecret, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	appSecret = appSecret[:len(appSecret)-1]

	fmt.Print("请输入对应权限点 (1: 方法 1, 2: 方法 2, 3: 方法 3, q: 退出): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	// 去除输入末尾的换行符
	input = input[:len(input)-1]

	switch input {
	case "1":
		_, _ = user.AccountInfo(clientID, accessToken, appSecret)
	//case "2":
	//	method2()
	//case "3":
	//	method3()
	case "q":
		fmt.Println("退出程序")
		return
	default:
		fmt.Println("无效的命令，请重新输入")
	}

}
