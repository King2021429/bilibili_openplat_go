package service

import (
	"bufio"
	"fmt"
	"log"
	"openplat/dao"
	"openplat/model"
	"os"
)

// UserData  USER_DATA 获取用户数据 GET
func UserData(clientId string, accessToken string, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.DataUserStatUrl
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArcStat 获取单个稿件数据 GET
func ArcStat(clientId string, accessToken string, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArcStatUrl
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入resource_id: ")
	resourceId, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	resourceId = resourceId[:len(resourceId)-1]
	// 参数字典
	params := map[string]string{
		"resource_id": resourceId,
	}

	// 调用拼接函数
	fullURL, err := dao.BuildURL(url, params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 打印拼接好的完整 URL
	fmt.Println("Full URL:", fullURL)
	return dao.ApiRequest(reqJson, fullURL, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArcIncStats 获取整体稿件增量数据 GET
func ArcIncStats(clientId string, accessToken string, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArcIncStats
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArtStat 获取单一专栏数据 GET
func ArtStat(clientId string, accessToken string, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArtStatUrl
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArtIncStats 获取整体投稿增量数据 GET
func ArtIncStats(clientId string, accessToken string, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArtIncStats
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}
