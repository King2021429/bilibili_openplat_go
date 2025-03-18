package service

import (
	"encoding/json"
	"fmt"
	"openplat/dao"
	"openplat/model"
)

// VideoInit 文件上传预处理
func VideoInit(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ArcInitUrl
	videoInitReq := model.VideoInitReq{
		Name:  "test.mp4",
		Utype: "0",
	}
	videoInitReqJson, _ := json.Marshal(videoInitReq)
	resp, err = dao.ApiRequest(string(videoInitReqJson), url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
	if err != nil {
		fmt.Printf("VideoInit err:%+v", err)
	}
	return
}

// VideoArcComplete 合片
func VideoArcComplete(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ArcComplete
	resp, err = dao.ApiRequest("", url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
	if err != nil {
		fmt.Printf("VideoArcComplete err:%+v", err)
	}
	return
}

// ArcAddUrl 稿件提交 POST
func ArcAddUrl(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArcAddUrl
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArcTypeList 分区查询 GET
func ArcTypeList(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ArcTypeList
	return dao.ApiRequest("", url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArcAddFetch 稿件提交fetch模式
func ArcAddFetch(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArcAddFetch
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArcEdit 稿件编辑
func ArcEdit(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArcEdit
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArcDel 稿件删除
func ArcDel(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArcDel
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArcView 稿件查询
func ArcView(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArcView
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArcViewList 稿件列表查询
func ArcViewList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArcViewList
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}
