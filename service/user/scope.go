package user

import (
	"fmt"
	"openplat/dao"
	"openplat/model"
)

// AccountScope 查询用户已授权权限列表
func AccountScope(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.UserScopesUrl
	resp, err = dao.ApiRequest("", url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
	if err != nil {
		fmt.Printf("AccountScope err:%+v", err)
	}
	return
}
