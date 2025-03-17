package arc

import (
	"encoding/json"
	"fmt"
	"openplat/dao"
	"openplat/model"
)

// VideoInit 文件上传预处理
func VideoInit(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ArcInitUrl
	videoInitReq := model.VideoInitReq{
		Name:  "test.mp4",
		Utype: "0",
	}
	videoInitReqJson, _ := json.Marshal(videoInitReq)
	resp, err = dao.ApiRequest(string(videoInitReqJson), url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
	if err != nil {
		fmt.Printf("VideoInit err:%+v", err)
	}
	return
}
