package service

import (
	"bufio"
	"fmt"
	"log"
	"openplat/dao"
	"openplat/model"
	"os"
)

// ArticleAdd 文章投稿
func ArticleAdd(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArticleAddUrl
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArticleEdit 文章编辑
func ArticleEdit(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArticleEdit
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArticleDelete 文章删除
func ArticleDelete(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArticleDelete
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArticleDetail 文章详情
func ArticleDetail(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArticleDetail
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArticleList 文章列表
func ArticleList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArticleList
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArticleCategories 文章分类
func ArticleCategories(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArticleCategories
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArticleCard 获取视频、文章卡片信息
func ArticleCard(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArticleCard
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AnthologyAdd 文集提交
func AnthologyAdd(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AnthologyAdd
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入reqJson串: ")
	reqJson, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AnthologyEdit 文集信息编辑
func AnthologyEdit(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AnthologyEdit
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ArticleBelong 文集下文章列表修改
func ArticleBelong(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ArticleBelong
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AnthologyDelete 文集删除
func AnthologyDelete(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AnthologyDelete
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AnthologyList 文集列表查询
func AnthologyList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AnthologyList
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AnthologyDetail 文集详情查询
func AnthologyDetail(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AnthologyDetail
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}
