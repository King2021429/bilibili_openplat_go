package service

import (
	"openplat/dao"
	"openplat/model"
)

// UserData  USER_DATA 获取用户数据 GET
func UserData(clientId string, accessToken string, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.DataUserStatUrl
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArcStat 获取单个稿件数据 GET
func ArcStat(clientId string, accessToken string, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArcStatUrl
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
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
