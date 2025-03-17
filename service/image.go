package service

import (
	"openplat/dao"
	"openplat/model"
)

func ImageUpload(picPath string) (resp model.BaseResp, err error) {
	url := model.ImageUploadCustomer
	return dao.PicRequest(url, picPath)
}
