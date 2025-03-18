package main

import (
	"bufio"
	"fmt"
	"log"
	"openplat/service"
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

	fmt.Print("请输入对应方法:" +
		"0: 计算签名" +
		"1: 查询用户已授权权限列表 " +
		"2: 方法 " +
		"3: 方法" +
		"q: 退出")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	// 去除输入末尾的换行符
	input = input[:len(input)-1]

	switch input {
	case "0":
		_, _ = service.Sign(clientID, accessToken, appSecret)
	case "1":
		_, _ = service.AccountInfo(clientID, accessToken, appSecret)
	case "2":
		_, _ = service.AccountScope(clientID, accessToken, appSecret)
	case "3":
		_, _ = service.ConversationSendMsg(clientID, accessToken, appSecret, "")
	case "4":
		_, _ = service.ConversationCustomerUserFrom(clientID, accessToken, appSecret, "")
	case "q":
		fmt.Println("退出程序")
		return
	default:
		fmt.Println("无效的命令，请重新输入")
	}

}
