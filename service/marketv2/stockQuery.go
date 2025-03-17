package service

import (
	"fmt"
	"openplat/dao"
	"openplat/model"
)

// StockQuery 新增共享
func StockQuery(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.StockQuery
	resp, err = dao.ApiRequest("", url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
	if err != nil {
		fmt.Printf("StockQuery err:%+v", err)
	}

	// 解析返回值
	queryResp := &model.StockInfo{}
	if queryResp == nil {
		fmt.Println("StockQuery Data nil")
	}
	fmt.Printf("StockQuery:%+v\n", queryResp)
	return
}
