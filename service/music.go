package service

import (
	"openplat/dao"
	"openplat/model"
)

// MusicMetaList 音乐列表 获取媒体元数据
func MusicMetaList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.MusicMetaListUrl
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// MusicList 音乐列表 获取媒体资源信息
func MusicList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.MusicListUrl
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// MusicSearch 音乐搜索
func MusicSearch(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.MusicSearchUrl
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}
