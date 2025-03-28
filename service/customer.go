package service

import (
	"bufio"
	"fmt"
	"log"
	"openplat/dao"
	"openplat/model"
	"os"
)

// ConversationSendMsg 客服系统 消息发送
func ConversationSendMsg(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ConversationSendMsgUrl
	reader := bufio.NewReader(os.Stdin)
	// 输入 access_token
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ConversationCustomerUserFrom 客服系统 获取用户来源
func ConversationCustomerUserFrom(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ConversationCustomerUserFrom
	reader := bufio.NewReader(os.Stdin)
	// 输入 access_token
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ConversationStaffStatusUpdate 客服系统 修改客服状态
func ConversationStaffStatusUpdate(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ConversationStaffStatusUpdate
	reader := bufio.NewReader(os.Stdin)
	// 输入 access_token
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ConversationClose 客服系统 关闭
func ConversationClose(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ConversationClose
	reader := bufio.NewReader(os.Stdin)
	// 输入 access_token
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}
