package service

import (
	"encoding/json"
	"fmt"
	"openplat/dao"
	"openplat/model"
)

// StockUpdate 库存更新
func StockUpdate(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.StockUpdate
	queryReq := model.StockUpdateReq{}
	queryReqJson, _ := json.Marshal(queryReq)
	resp, err = dao.ApiGetRequest(string(queryReqJson), url)
	if err != nil {
		fmt.Printf("StockUpdate err:%+v", err)
	}

	// 解析返回值
	queryResp := &model.StockInfo{}

	if queryResp == nil {
		fmt.Println("StockUpdate Data nil")
	}
	fmt.Printf("StockUpdate:%+v\n", queryResp)
	return
}
