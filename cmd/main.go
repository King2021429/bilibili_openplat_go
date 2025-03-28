package main

import (
	"bufio"
	"fmt"
	"log"
	"openplat/service"
	"os"
)

// 入口启动函数
func main() {
	// 调用方法
	// 创建一个读取器来读取命令行输入
	reader := bufio.NewReader(os.Stdin)

	// 输入 client_id
	fmt.Print("请输入 client_id: ")
	clientID, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	clientID = clientID[:len(clientID)-1]

	// 输入 access_token
	fmt.Print("请输入 access_token: ")
	accessToken, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	accessToken = accessToken[:len(accessToken)-1]

	// 输入 app_secret
	fmt.Print("请输入 app_secret: ")
	appSecret, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	appSecret = appSecret[:len(appSecret)-1]

	// TODO 直播
	fmt.Print("请输入对应方法:" +
		"0: 计算签名\n" +
		"1: 获取授权\n" +
		"2: USER_INFO-获取已授权用户基础公开信息\n" +
		"3: USER_INFO-查询用户已授权权限列表\n" +
		"4: ARC_BASE-文件上传预处理\n" +
		"7: ARC_BASE-文件分片合片\n" +
		"8: ARC_BASE-视频稿件提交\n" +
		"9: ARC_BASE-获取用于投稿的连接\n" +
		"10: ARC_BASE-查询单一视频稿件详情\n" +
		"11: ARC_BASE-查询用户视频稿件列表\n" +
		"12: ATC_BASE-文章提交\n" +
		"13: ATC_BASE-文章编辑\n" +
		"14: ATC_BASE-文章删除\n" +
		"15: ATC_BASE-文章详情\n" +
		"16: ATC_BASE-文章列表查询\n" +
		"17: ATC_BASE-文章分类查询\n" +
		"18: ATC_BASE-获取视频、文章卡片信息\n" +
		"19: ATC_BASE-文集提交\n" +
		"20: ATC_BASE-文集信息编辑\n" +
		"21: ATC_BASE-文集下文章列表修改\n" +
		"22: ATC_BASE-文集删除\n" +
		"23: ATC_BASE-文集列表查询\n" +
		"24: ATC_BASE-文集详情查询\n" +
		"26: USER_DATA-获取用户数据\n" +
		"27: ARC_DATA-获取单个稿件数据\n" +
		"28: ARC_DATA-获取整体稿件增量数据\n" +
		"29: ATC_DATA-获取单一专栏数据\n" +
		"30: ATC_DATA-获取整体投稿增量数据\n" +
		"31: SHOP_STORE_INFO-获取店铺信息\n" +
		"32: SHOP_COMMODITY_INFO-商品发布\n" +
		"33: SHOP_COMMODITY_INFO-商品编辑\n" +
		"34: SHOP_COMMODITY_INFO-商品删除\n" +
		"35: SHOP_COMMODITY_INFO-商品列表查询\n" +
		"36: SHOP_COMMODITY_INFO-商品详情查询\n" +
		"37: SHOP_COMMODITY_INFO-查询发品配置\n" +
		"38: SHOP_COMMODITY_INFO-查询商品类目属性列表\n" +
		"39: SHOP_COMMODITY_INFO-获取类目列表\n" +
		"40: SHOP_ORDER_INFO-订单列表查询\n" +
		"41: SHOP_ORDER_INFO-订单详情查询\n" +
		"42: SHOP_ORDER_INFO-订单敏感信息解密\n" +
		"43: SHOP_ORDER_INFO-流量卡审核信息回传\n" +
		"44: SHOP_ORDER_INFO-商家备注\n" +
		"45: SHOP_LOGISTICS_INFO-发货回掉\n" +
		"46: SHOP_LOGISTICS_INFO-物流编辑\n" +
		"47: SHOP_LOGISTICS_INFO-获取快递公司列表\n" +
		"48: SHOP_LOGISTICS_INFO-地址库创建\n" +
		"49: SHOP_LOGISTICS_INFO-批量查询地址库列表\n" +
		"50: SHOP_LOGISTICS_INFO-获取全量省份信息\n" +
		"51: SHOP_LOGISTICS_INFO-根据省获取全量三级地址\n" +
		"52: SHOP_LOGISTICS_INFO-运费模板列表查询\n" +
		"53: SHOP_AFTERSALES_INFO-售后列表查询\n" +
		"54: SHOP_AFTERSALES_INFO-售后详情查询\n" +
		"55: SHOP_AFTERSALES_INFO-售后审核\n" +
		"56: SHOP_AFTERSALES_INFO-确认收货\n" +
		"57: SHOP_AFTERSALES_INFO-换货发货\n" +
		"58: SHOP_AFTERSALES_INFO-售后终止\n" +
		"59: SHOP_STOCK_INFO-库存查询\n" +
		"60: SHOP_STOCK_INFO-库存更新\n" +
		"61: SHOP_BA_CHAT-客服系统 消息发送\n" +
		"62: SHOP_BA_CHAT-客服系统 获取用户来源\n" +
		"63: SHOP_BA_CHAT-客服系统 修改客服状态\n" +
		"64: SHOP_BA_CHAT-客服系统 关闭\n" +
		"65: COPYRIGHT_MUSIC_DATA-音乐列表 获取媒体元数据\n" +
		"66: COPYRIGHT_MUSIC_DATA-音乐列表 获取媒体资源信息\n" +
		"67: COPYRIGHT_MUSIC_DATA-音乐搜索\n" +
		"68:稿件图片上传\n" +
		"69:专栏稿件上传\n" +
		"70:商品图片上传\n" +
		"71:客服图片上传\n" +
		"q: 退出 ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	// 去除输入末尾的换行符
	input = input[:len(input)-1]

	switch input {
	case "0":
		// 计算签名
		service.Sign(clientID, accessToken, appSecret)
	case "1":
		// 获取授权
		service.Oauth(clientID, appSecret)
	case "2":
		// 获取用户信息
		_, _ = service.AccountInfo(clientID, accessToken, appSecret)
	case "3":
		// 查询用户已授权权限列表
		_, _ = service.AccountScope(clientID, accessToken, appSecret)
	case "4":
		// 文件上传预处理
		_, _ = service.VideoInit(clientID, accessToken, appSecret)
	case "7":
		// 文件分片合片
		_, _ = service.VideoArcComplete(clientID, accessToken, appSecret)
	case "8":
		// 视频稿件提交
		fmt.Print("请输入reqJson串: ")
		reqJson, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("读取输入时出错: %v", err)
		}
		reqJson = reqJson[:len(reqJson)-1]
		_, _ = service.ArcAddUrl(clientID, accessToken, appSecret, reqJson)
	case "9":
		// 获取用于投稿的连接
		_, _ = service.CommonAddShare(clientID, accessToken, appSecret, "")
	case "10":
		// 查询单一视频稿件详情
		_, _ = service.ArcView(clientID, accessToken, appSecret, "")
	case "11":
		// 查询用户视频稿件列表
		_, _ = service.ArcViewList(clientID, accessToken, appSecret, "")
	case "12":
		// 文章投稿
		reqJson := "" // 这里你可能需要根据实际情况构造请求的 JSON 字符串
		_, _ = service.ArticleAdd(clientID, accessToken, appSecret, reqJson)
	case "13":
		// 文章编辑
		fmt.Print("请输入reqJson串: ")
		reqJson, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("读取输入时出错: %v", err)
		}
		reqJson = reqJson[:len(reqJson)-1]
		_, _ = service.ArticleEdit(clientID, accessToken, appSecret, reqJson)
	case "14":
		// 文章删除
		fmt.Print("请输入reqJson串: ")
		reqJson, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("读取输入时出错: %v", err)
		}
		reqJson = reqJson[:len(reqJson)-1]
		_, _ = service.ArticleDelete(clientID, accessToken, appSecret, reqJson)
	case "15":
		// 文章详情
		fmt.Print("请输入reqJson串: ")
		reqJson, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("读取输入时出错: %v", err)
		}
		reqJson = reqJson[:len(reqJson)-1]
		_, _ = service.ArticleDetail(clientID, accessToken, appSecret, reqJson)
	case "16":
		_, _ = service.ArticleList(clientID, accessToken, appSecret, "")
	case "17":
		// 文章分类查询
		reqJson := ""
		_, _ = service.ArticleCategories(clientID, accessToken, appSecret, reqJson)
	case "18":
		// 获取视频、文章卡片信息
		reqJson := ""
		_, _ = service.ArticleCard(clientID, accessToken, appSecret, reqJson)
	case "19":
		// 文集提交
		reqJson := ""
		_, _ = service.AnthologyAdd(clientID, accessToken, appSecret, reqJson)
	case "20":
		// 文集信息编辑
		reqJson := ""
		_, _ = service.AnthologyEdit(clientID, accessToken, appSecret, reqJson)
	case "21":
		// 文集下文章列表修改
		reqJson := ""
		_, _ = service.ArticleBelong(clientID, accessToken, appSecret, reqJson)
	case "22":
		// 文集删除
		reqJson := ""
		_, _ = service.AnthologyDelete(clientID, accessToken, appSecret, reqJson)
	case "23":
		// 文集列表查询
		reqJson := ""
		_, _ = service.AnthologyList(clientID, accessToken, appSecret, reqJson)
	case "24":
		// 文集详情查询
		reqJson := ""
		_, _ = service.AnthologyDetail(clientID, accessToken, appSecret, reqJson)
	case "26":
		// 获取用户数据 GET
		reqJson := ""
		_, _ = service.UserData(clientID, accessToken, appSecret, reqJson)
	case "27":
		// 获取单个稿件数据 GET
		reqJson := ""
		_, _ = service.ArcStat(clientID, accessToken, appSecret, reqJson)
	case "28":
		// 获取整体稿件增量数据 GET
		reqJson := ""
		_, _ = service.ArcIncStats(clientID, accessToken, appSecret, reqJson)
	case "29":
		// ArtStat 获取单一专栏数据 GET
		reqJson := ""
		_, _ = service.ArtStat(clientID, accessToken, appSecret, reqJson)
	case "30":
		// 获取整体投稿增量数据 GET
		reqJson := ""
		_, _ = service.ArtIncStats(clientID, accessToken, appSecret, reqJson)
	case "31":
		// 获取店铺信息 GET
		reqJson := ""
		_, _ = service.ShopInfoGetUrl(clientID, accessToken, appSecret, reqJson)
	case "32":
		// 商品发布
		reqJson := ""
		_, _ = service.ProductAdd(clientID, accessToken, appSecret, reqJson)
	case "33":
		// 商品编辑
		reqJson := ""
		_, _ = service.ProductEdit(clientID, accessToken, appSecret, reqJson)
	case "34":
		// 商品删除
		reqJson := ""
		_, _ = service.ProductDel(clientID, accessToken, appSecret, reqJson)
	case "35":
		// 商品列表查询
		reqJson := ""
		_, _ = service.CommodityItemList(clientID, accessToken, appSecret, reqJson)
	case "36":
		// 商品详情查询
		reqJson := ""
		_, _ = service.ProductDetail(clientID, accessToken, appSecret, reqJson)
	case "37":
		// 查询发品配置
		reqJson := ""
		_, _ = service.ProductGetPublishRule(clientID, accessToken, appSecret, reqJson)
	case "38":
		// 查询商品类目属性列表
		reqJson := ""
		_, _ = service.ProductGetCateProperty(clientID, accessToken, appSecret, reqJson)
	case "39":
		// 获取类目列表
		reqJson := ""
		_, _ = service.CommodityCategoryQualifiedList(clientID, accessToken, appSecret, reqJson)
	case "40":
		// 订单列表查询
		reqJson := ""
		_, _ = service.OrderSearchList(clientID, accessToken, appSecret, reqJson)
	case "41":
		// 订单详情查询
		reqJson := ""
		_, _ = service.OrderDetail(clientID, accessToken, appSecret, reqJson)
	case "42":
		// 订单敏感信息解密
		reqJson := ""
		_, _ = service.OrderBatchDecrypt(clientID, accessToken, appSecret, reqJson)
	case "43":
		// 流量卡审核信息回传
		reqJson := ""
		_, _ = service.OrderReview(clientID, accessToken, appSecret, reqJson)
	case "44":
		// 商家备注
		reqJson := ""
		_, _ = service.OrderRemark(clientID, accessToken, appSecret, reqJson)
	case "45":
		// 发货回掉
		reqJson := ""
		_, _ = service.LogisticsAdd(clientID, accessToken, appSecret, reqJson)
	case "46":
		// 物流编辑
		reqJson := ""
		_, _ = service.LogisticsEdit(clientID, accessToken, appSecret, reqJson)
	case "47":
		// 获取快递公司列表
		reqJson := ""
		_, _ = service.LogisticsCompanyList(clientID, accessToken, appSecret, reqJson)
	case "48":
		// 地址库创建
		reqJson := ""
		_, _ = service.AddressCreate(clientID, accessToken, appSecret, reqJson)
	case "49":
		// 批量查询地址库列表
		reqJson := ""
		_, _ = service.AddressList(clientID, accessToken, appSecret, reqJson)
	case "50":
		// 获取全量省份信息
		reqJson := ""
		_, _ = service.AddressGetProvince(clientID, accessToken, appSecret, reqJson)
	case "51":
		// 根据省获取全量三级地址
		reqJson := ""
		_, _ = service.AddressGetAreasByProvince(clientID, accessToken, appSecret, reqJson)
	case "52":
		// 运费模板列表查询
		reqJson := ""
		_, _ = service.LogisticsFreightTemplateList(clientID, accessToken, appSecret, reqJson)
	case "53":
		// 售后列表查询
		reqJson := ""
		_, _ = service.AfterSaleQueryList(clientID, accessToken, appSecret, reqJson)
	case "54":
		// 售后详情查询
		reqJson := ""
		_, _ = service.AfterSaleQueryDetail(clientID, accessToken, appSecret, reqJson)
	case "55":
		// 售后审核
		reqJson := ""
		_, _ = service.AfterSaleCheckAfterSale(clientID, accessToken, appSecret, reqJson)
	case "56":
		// 确认收货
		reqJson := ""
		_, _ = service.AfterSaleConfirmReceipt(clientID, accessToken, appSecret, reqJson)
	case "57":
		// 换货发货
		reqJson := ""
		_, _ = service.AfterSaleBarterShip(clientID, accessToken, appSecret, reqJson)
	case "58":
		// 售后终止
		reqJson := ""
		_, _ = service.AfterSaleStop(clientID, accessToken, appSecret, reqJson)
	case "59":
		// StockQuery 库存查询
		_, _ = service.StockQuery(clientID, accessToken, appSecret)
	case "60":
		// StockUpdate 库存更新
		_, _ = service.StockUpdate(clientID, accessToken, appSecret)
	case "61":
		// 客服系统 消息发送
		_, _ = service.ConversationSendMsg(clientID, accessToken, appSecret, "")
	case "62":
		// 客服系统 获取用户来源
		_, _ = service.ConversationCustomerUserFrom(clientID, accessToken, appSecret, "")
	case "63":
		// 客服系统 修改客服状态
		_, _ = service.ConversationStaffStatusUpdate(clientID, accessToken, appSecret, "")
	case "64":
		// 客服系统 关闭
		_, _ = service.ConversationClose(clientID, accessToken, appSecret, "")
	case "65":
		// MusicMetaList 音乐列表 获取媒体元数据
		_, _ = service.MusicMetaList(clientID, accessToken, appSecret, "")
	case "66":
		// MusicList 音乐列表 获取媒体资源信息
		_, _ = service.MusicList(clientID, accessToken, appSecret, "")
	case "67":
		// MusicSearch 音乐搜索
		_, _ = service.MusicSearch(clientID, accessToken, appSecret, "")
	case "68":
		picUrl := ""
		// ImageUploadArc 稿件图片上传
		_, _ = service.ImageUploadArc(picUrl, clientID, accessToken, appSecret)
	case "69":
		picUrl := ""
		// ImageUploadArticle 专栏稿件上传
		_, _ = service.ImageUploadArticle(picUrl, clientID, accessToken, appSecret)
	case "70":
		picUrl := ""
		// ImageUploadCommodity 商品图片上传
		_, _ = service.ImageUploadCommodity(picUrl, clientID, accessToken, appSecret)
	case "71":
		picUrl := ""
		// ImageUploadCustomer 客服图片上传
		_, _ = service.ImageUploadCustomer(picUrl, clientID, accessToken, appSecret)
	case "q":
		fmt.Println("退出程序")
		return
	default:
		fmt.Println("无效的命令，请重新输入")
	}

}
