package service

import (
	"encoding/json"
	"fmt"
	"openplat/dao"
	"openplat/model"
)

// MsgQuery 获取聊天记录
func MsgQuery() (resp model.BaseResp, err error) {
	queryReq := model.OpenMarketCustomerSendMsgReq{}
	queryReqJson, _ := json.Marshal(queryReq)
	resp, err = dao.ApiRequest(string(queryReqJson), model.ConversationCustomerMsgQuery)
	if err != nil {
		fmt.Printf("MsgQuery err:%+v", err)
	}

	// 解析返回值
	//queryResp := &model.QueryResp{}
	//err = json.Unmarshal(resp.Data, &queryResp)
	//if err != nil {
	//	fmt.Printf("MsgQuery resp json unmarshal err:%+v", err)
	//}
	//if queryResp == nil {
	//	fmt.Println("MsgQuery queryResp Data nil")
	//}
	fmt.Printf("MsgQuery queryResp:%+v\n", resp)
	return
}
