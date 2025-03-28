package service

import (
	"openplat/dao"
	"openplat/model"
)

// ImageUploadArc 稿件图片上传
func ImageUploadArc(picPath, clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ImageUploadArcUrl
	return dao.PicRequest(url, picPath, clientId, accessToken, appSecret, model.BiliVersionV2, 1)
}

// ImageUploadArticle 专栏稿件上传
func ImageUploadArticle(picPath, clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ImageUploadArticleUrl
	return dao.PicRequest(url, picPath, clientId, accessToken, appSecret, model.BiliVersionV2, 2)
}

// ImageUploadCommodity 商品图片上传
func ImageUploadCommodity(picPath, clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ImageUploadCommodityUrl
	return dao.PicRequest(url, picPath, clientId, accessToken, appSecret, model.BiliVersionV2, 3)
}

// ImageUploadCustomer 客服图片上传
func ImageUploadCustomer(picPath, clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	url := model.ImageUploadCustomer
	return dao.PicRequest(url, picPath, clientId, accessToken, appSecret, model.BiliVersionV2, 4)
}
