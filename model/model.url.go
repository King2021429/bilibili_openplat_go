package model

// USER_INFO
const (
	AccountInfoUrl = "/arcopen/fn/user/account/info"   // 获取用户公开信息 GET
	UserScopesUrl  = "/arcopen/fn/user/account/scopes" // 查询用户已授权权限列表 GET
)

// ARC
const (
	ArcInitUrl  = "/arcopen/fn/archive/video/init"                                                   //文件上传预处理
	ArcComplete = "/arcopen/fn/archive/video/complete?upload_token=7213c24789bf42b3a3482b7c7d9a597f" // 文件分片合片
	ArcAddUrl   = "/arcopen/fn/archive/add"                                                          // 稿件提交 POST
	ArcTypeList = "/arcopen/fn/archive/type/list"                                                    // 分区查询 GET "/arcopen/fn/archive/add-by-utoken", arcAddByUTokenMock) // 稿件提交(utoken)
	ArcAddFetch = "/arcopen/fn/archive/add-fetch"                                                    // 稿件提交fetch模式
	ArcEdit     = "/arcopen/fn/archive/edit"                                                         // 稿件编辑
	ArcDel      = "/arcopen/fn/archive/delete"                                                       // 稿件删除
	ArcView     = "/arcopen/fn/archive/view"                                                         // 稿件查询
	ArcViewList = "/arcopen/fn/archive/viewlist"                                                     // 稿件列表查询
)

const (
	// 文章
	ArticleAddUrl     = "/arcopen/fn/article/add" // 投稿 POST
	ArticleEdit       = "/arcopen/fn/article/edit"
	ArticleDelete     = "/arcopen/fn/article/delete"
	ArticleDetail     = "/arcopen/fn/article/detail"
	ArticleList       = "/arcopen/fn/article/list"
	ArticleCategories = "/arcopen/fn/article/category"
	ArticleCard       = "/arcopen/fn/article/card" // 获取视频、文章卡片信息 GET

	// 文集
	AnthologyAdd    = "/arcopen/fn/article/anthology/add"    // 文集提交
	AnthologyEdit   = "/arcopen/fn/article/anthology/edit"   // 文集信息编辑
	ArticleBelong   = "/arcopen/fn/article/belong"           // 文集下文章列表修改
	AnthologyDelete = "/arcopen/fn/article/anthology/delete" // 文集删除
	AnthologyList   = "/arcopen/fn/article/anthology/list"   // 文集列表查询
	AnthologyDetail = "/arcopen/fn/article/anthology/detail" // 文集详情查询
)

const (
	// RESOURCE
	ResourceAddShareUrl = "/arcopen/fn/resource/add_share" // 新增共享 POST
)

// USER_DATA
const (
	DataUserStatUrl = "/arcopen/fn/data/user/stat" // 获取用户数据 GET
)

// ARC_DATA
const (
	ArcStatUrl  = "/arcopen/fn/data/arc/stat"      // 获取单个稿件数据 GET
	ArcIncStats = "/arcopen/fn/data/arc/inc-stats" // 获取整体稿件增量数据 GET
)

// ATC_DATA
const (
	ArtStatUrl  = "/arcopen/fn/data/art/stat"      // 获取单一专栏数据 GET
	ArtIncStats = "/arcopen/fn/data/art/inc-stats" // 获取整体投稿增量数据 GET
)

const (
	// SHOP_STORE_INFO
	ShopInfoGetUrl = "/arcopen/fn/v2/market/shop/info/get" // 获取店铺信息 GET

	// SHOP_COMMODITY_INFO
	ProductAddUrl                  = "/arcopen/fn/market/common/product_add"                      // 商品发布 POST
	CommodityItemListUrl           = "/arcopen/fn/v2/market/commodity/item/list"                  // 查询商品列表 GET
	ProductGetPublishRuleUrl       = "/arcopen/fn/market/common/product_get_product_publish_rule" // 查询发品配置 POST
	ProductDetail                  = "/arcopen/fn/market/common/product/detail"
	ProductEdit                    = "/arcopen/fn/market/common/product/edit"
	ProductDel                     = "/arcopen/fn/market/common/product/del"
	ProductGetCateProperty         = "/arcopen/fn/market/common/product/get_cate_property"
	CommodityCategoryQualifiedList = "/arcopen/fn/market/commodity/category/qualified/list"

	// SHOP_ORDER_INFO 订单
	OrderBatchDecrypt = "/arcopen/fn/market/common/order_batch_decrypt" //订单敏感信息解密
	OrderReview       = "/arcopen/fn/market/common/order_review"        //流量卡审核信息回传
	OrderRemark       = "/arcopen/fn/market/common/order_remark"        // 商家备注
	OrderDetail       = "/arcopen/fn/market/common/order_detail"        // 订单详情查询
	OrderSearchList   = "/arcopen/fn/market/common/order_searchList"    // 订单列表查询

	// SHOP_LOGISTICS_INFO
	LogisticsAdd                 = "/arcopen/fn/market/common/logistics_add"                   // 发货回掉
	LogisticsEdit                = "/arcopen/fn/market/common/logistics_edit"                  // 物流编辑
	LogisticsCompanyList         = "/arcopen/fn/market/common/logistics_company_list"          // 获取快递公司列表
	AddressCreate                = "/arcopen/fn/market/common/address_create"                  // 地址库创建
	AddressList                  = "/arcopen/fn/market/common/address/list"                    // 批量查询地址库列表
	AddressGetProvince           = "/arcopen/fn/market/common/address_get_province"            // 获取全量省份信息
	AddressGetAreasByProvince    = "/arcopen/fn/market/common/address_get_areas_by_province"   // 根据省获取全量三级地址
	LogisticsFreightTemplateList = "/arcopen/fn/market/common/logistics_freight_template_list" // 运费模板列表查询

	// SHOP_AFTERSALES_INFO 售后
	AfterSaleQueryList      = "/arcopen/fn/market/common/after_sale_query_list"       // 售后列表查询
	AfterSaleQueryDetail    = "/arcopen/fn/market/common/after_sale_query_detail"     // 售后详情查询
	AfterSaleCheckAfterSale = "/arcopen/fn/market/common/after_sale_check_after_sale" // 售后审核
	AfterSaleConfirmReceipt = "/arcopen/fn/market/common/after_sale_confirm_receipt"  // 确认收货
	AfterSaleBarterShip     = "/arcopen/fn/market/common/after_sale_barter_ship"      // 换货发货
	AfterSaleStop           = "/arcopen/fn/market/common/after_sale_stop"             // 售后终止
	// SHOP_STOCK_INFO
	StockQuery  = "/arcopen/fn/v2/market/stock/query"  // 库存查询 GET
	StockUpdate = "/arcopen/fn/v2/market/stock/update" // 库存更新 POST
)

// SHOP_BA_CHAT
const (
	ConversationSendMsgUrl        = "/arcopen/fn/market/customer/conversation/send_msg"            // 客服系统 消息发送
	ConversationCustomerUserFrom  = "/arcopen/fn/market/customer/conversation/user_from"           // 客服系统 获取用户来源
	ConversationStaffStatusUpdate = "/arcopen/fn/market/customer/conversation/staff_status_update" // 客服系统 修改客服状态
	ConversationClose             = "/arcopen/fn/market/customer/conversation/close"               // 客服系统 关闭
)

// COPYRIGHT_MUSIC_DATA
const (
	MusicMetaListUrl = "/arcopen/fn/music/meta/list" // 音乐列表 获取媒体元数据
	MusicListUrl     = "/arcopen/fn/music/list"      // 音乐列表 获取媒体资源信息
	MusicSearch      = "/arcopen/fn/music/search"    // 音乐搜索
)

// 图片上传url模块
const (
	// ImageUploadArcUrl 稿件图片上传
	ImageUploadArcUrl = "https://member.bilibili.com/arcopen/fn/archive/cover/upload"

	// ImageUploadArticleUrl 专栏稿件上传
	ImageUploadArticleUrl = "https://member.bilibili.com/arcopen/fn/article/upload/image"

	// ImageUploadCommodityUrl 商品图片上传
	ImageUploadCommodityUrl = "https://member.bilibili.com/arcopen/fn/v2/market/commodity/image/upload"

	// ImageUploadCustomer 客服图片上传
	ImageUploadCustomer = "https://member.bilibili.com/arcopen/fn/market/customer/image_upload"
)
