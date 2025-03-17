package user

import (
	"fmt"
	"openplat/dao"
	"openplat/model"
)

// AccountInfo 查询用户已授权权限列表
func AccountInfo(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.AccountInfoUrl
	resp, err = dao.ApiRequest("", url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
	if err != nil {
		fmt.Printf("AccountInfo err:%+v", err)
	}
	return
}
