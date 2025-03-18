package service

import (
	"openplat/dao"
	"openplat/model"
)

// ProductAdd 商品发布
func ProductAdd(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ProductAddUrl
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// CommodityItemList 查询商品列表
func CommodityItemList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.CommodityItemListUrl
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ProductGetPublishRule 查询发品配置
func ProductGetPublishRule(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ProductGetPublishRuleUrl
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ProductDetail 获取商品详情
func ProductDetail(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ProductDetail
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ProductEdit 编辑商品
func ProductEdit(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ProductEdit
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ProductDel 删除商品
func ProductDel(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ProductDel
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// ProductGetCateProperty 获取商品分类属性
func ProductGetCateProperty(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.ProductGetCateProperty
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// CommodityCategoryQualifiedList 获取商品分类合格列表
func CommodityCategoryQualifiedList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.CommodityCategoryQualifiedList
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// OrderBatchDecrypt 订单敏感信息解密
func OrderBatchDecrypt(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.OrderBatchDecrypt
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// OrderReview 流量卡审核信息回传
func OrderReview(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.OrderReview
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// OrderRemark 商家备注
func OrderRemark(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.OrderRemark
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// OrderDetail 订单详情查询
func OrderDetail(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.OrderDetail
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// OrderSearchList 订单列表查询
func OrderSearchList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.OrderSearchList
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// LogisticsAdd 发货回掉
func LogisticsAdd(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.LogisticsAdd
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// LogisticsEdit 物流编辑
func LogisticsEdit(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.LogisticsEdit
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// LogisticsCompanyList 获取快递公司列表
func LogisticsCompanyList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.LogisticsCompanyList
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AddressCreate 地址库创建
func AddressCreate(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AddressCreate
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AddressList 批量查询地址库列表
func AddressList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AddressList
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AddressGetProvince 获取全量省份信息
func AddressGetProvince(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AddressGetProvince
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AddressGetAreasByProvince 根据省获取全量三级地址
func AddressGetAreasByProvince(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AddressGetAreasByProvince
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// LogisticsFreightTemplateList 运费模板列表查询
func LogisticsFreightTemplateList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.LogisticsFreightTemplateList
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AfterSaleQueryList 售后列表查询
func AfterSaleQueryList(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AfterSaleQueryList
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AfterSaleQueryDetail 售后详情查询
func AfterSaleQueryDetail(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AfterSaleQueryDetail
	return dao.ApiRequest(reqJson, url, model.MethodGet, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AfterSaleCheckAfterSale 售后审核
func AfterSaleCheckAfterSale(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AfterSaleCheckAfterSale
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AfterSaleConfirmReceipt 确认收货
func AfterSaleConfirmReceipt(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AfterSaleConfirmReceipt
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AfterSaleBarterShip 换货发货
func AfterSaleBarterShip(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AfterSaleBarterShip
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}

// AfterSaleStop 售后终止
func AfterSaleStop(clientId, accessToken, appSecret string, reqJson string) (resp model.BaseResp, err error) {
	url := model.AfterSaleStop
	return dao.ApiRequest(reqJson, url, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersionV2)
}
