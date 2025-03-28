package service

import (
	"bufio"
	"fmt"
	"log"
	"openplat/dao"
	"openplat/model"
	"os"
)

// MusicMetaList 音乐列表 获取媒体元数据
func MusicMetaList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.MusicMetaListUrl
	reader := bufio.NewReader(os.Stdin)
	// 输入 access_token
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// MusicList 音乐列表 获取媒体资源信息
func MusicList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.MusicListUrl
	reader := bufio.NewReader(os.Stdin)
	// 输入 access_token
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// MusicSearch 音乐搜索
func MusicSearch(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.MusicSearchUrl
	reader := bufio.NewReader(os.Stdin)
	// 输入 access_token
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}
