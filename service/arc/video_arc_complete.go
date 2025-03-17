package arc

import (
	"fmt"
	"openplat/dao"
	"openplat/model"
)

// VideoArcComplete 合片
func VideoArcComplete(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ArcComplete

	resp, err = dao.ApiRequest("", url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
	if err != nil {
		fmt.Printf("VideoArcComplete err:%+v", err)
	}
	return
}
