package service

import (
	"openplat/dao"
	"openplat/model"
)

func ImageUpload(picPath, clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ImageUploadCustomer
	return dao.PicRequest(url, picPath, clientId, accessToken, appSecret, model.BiliVersionV2)
}
