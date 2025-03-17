package service

import (
	"encoding/json"
	"fmt"
	"openplat/dao"
	"openplat/model"
)

// SendMsg 新增共享
func SendMsg(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	queryReq := model.OpenMarketCustomerSendMsgReq{
		MsgType:        1,
		ConversationId: 115957154487296,
		Msg:            "{\"content\":\"888\"}",
		UserOpenId:     "",
	}
	queryReqJson, _ := json.Marshal(queryReq)
	resp, err = dao.ApiRequest(string(queryReqJson), model.ConversationSendMsgUrl)
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
