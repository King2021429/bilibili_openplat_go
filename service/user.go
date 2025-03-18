package service

import (
	"fmt"
	"openplat/dao"
	"openplat/model"
)

// AccountInfo 查询用户账号信息
func AccountInfo(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.AccountInfoUrl
	resp, err = dao.ApiRequest("", url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
	if err != nil {
		fmt.Printf("AccountInfo err:%+v", err)
	}
	return
}

// AccountScope 查询用户已授权权限列表
func AccountScope(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.UserScopesUrl
	resp, err = dao.ApiRequest("", url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
	if err != nil {
		fmt.Printf("AccountScope err:%+v", err)
	}
	return
}
