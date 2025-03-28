package service

import (
	"bufio"
	"fmt"
	"log"
	"openplat/dao"
	"openplat/model"
	"os"
)

// VideoInit 文件上传预处理
func VideoInit(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ArcInitUrl
	//videoInitReq := model.VideoInitReq{
	//	Name:  "test.mp4",
	//	Utype: "0",
	//}
	//videoInitReqJson, _ := json.Marshal(videoInitReq)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入reqJson串: ")
	reqJson, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// VideoArcComplete 文件分片合片
func VideoArcComplete(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ArcComplete
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入reqJson串: ")
	reqJson, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArcAddUrl 稿件提交 POST
func ArcAddUrl(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArcAddUrl
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
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
