package service

import (
	"encoding/json"
	"fmt"
	"openplat/dao"
	"openplat/model"
)

// ConversationSendMsg 客服系统 消息发送
func ConversationSendMsg(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ConversationSendMsgUrl
	queryReq := model.OpenMarketCustomerSendMsgReq{
		MsgType:        1,
		ConversationId: 115957154487296,
		Msg:            "{\"content\":\"888\"}",
		UserOpenId:     "",
	}
	queryReqJson, _ := json.Marshal(queryReq)
	resp, err = dao.ApiRequest(string(queryReqJson), url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
	if err != nil {
		fmt.Printf("SendMsg err:%+v", err)
	}

	// 解析返回值
	//queryResp := &model.QueryResp{}
	//err = json.Unmarshal(resp.Data, &queryResp)
	//if err != nil {
	//	fmt.Printf("SendMsg resp json unmarshal err:%+v", err)
	//}
	//if queryResp == nil {
	//	fmt.Println("queryResp Data nil")
	//}
	fmt.Printf("queryResp:%+v\n", resp)
	return
}

// ConversationCustomerUserFrom 客服系统 获取用户来源
func ConversationCustomerUserFrom(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ConversationCustomerUserFrom
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ConversationStaffStatusUpdate 客服系统 修改客服状态
func ConversationStaffStatusUpdate(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ConversationStaffStatusUpdate
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ConversationClose 客服系统 关闭
func ConversationClose(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ConversationClose
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}
