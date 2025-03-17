package model

// USER_INFO
const (
	AccountInfoUrl = "/arcopen/fn/user/account/info"   // 获取用户公开信息 GET
	UserScopesUrl  = "/arcopen/fn/user/account/scopes" // 查询用户已授权权限列表 GET
)

// ARC
const (
	ArcInitUrl     = "/arcopen/fn/archive/video/init"                                                   //文件上传预处理
	ArcComplete    = "/arcopen/fn/archive/video/complete?upload_token=7213c24789bf42b3a3482b7c7d9a597f" // 文件分片合片
	arcAddUrl      = "/arcopen/fn/archive/add"                                                          // 稿件提交 POST
	arcTypeList    = "/arcopen/fn/archive/type/list"                                                    // 分区查询 GET
	arcCoverUpload = "/arcopen/fn/archive/cover/upload"                                                 // 封面上传 POST

)

const (
	// RESOURCE
	ResourceAddShareUrl = "/arcopen/fn/resource/add_share" // 新增共享 POST
)

// USER_DATA
const (
	DataUserStatUrl = "/arcopen/fn/data/user/stat" // 获取用户数据 GET
)

const (
	// SHOP_STORE_INFO
	ShopInfoGetUrl = "/arcopen/fn/v2/market/shop/info/get" // 获取店铺信息 GET

	// SHOP_COMMODITY_INFO
	ProductAddUrl            = "/arcopen/fn/market/common/product_add"                      // 商品发布 POST
	itemListUrl              = "/arcopen/fn/v2/market/commodity/item/list"                  // 查询商品列表 GET
	productGetPublishRuleUrl = "/arcopen/fn/market/common/product_get_product_publish_rule" // 查询发品配置 POST

	// SHOP_STOCK_INFO
	StockQuery  = "/arcopen/fn/v2/market/stock/query"  // 库存查询 GET
	StockUpdate = "/arcopen/fn/v2/market/stock/update" // 库存更新 POST
)

// SHOP_BA_CHAT
const (
	ConversationSendMsgUrl        = "/arcopen/fn/market/customer/conversation/send_msg"            // 客服系统 消息发送
	conversationCustomerUserFrom  = "/arcopen/fn/market/customer/conversation/user_from"           // 客服系统 获取用户来源
	conversationStaffStatusUpdate = "/arcopen/fn/market/customer/conversation/staff_status_update" // 客服系统 修改客服状态
	conversationClose             = "/arcopen/fn/market/customer/conversation/close"               // 客服系统 关闭
)

// COPYRIGHT_MUSIC_DATA
const (
	musicMetaListUrl = "/arcopen/fn/music/meta/list" // 音乐列表 获取媒体元数据
	musicListUrl     = "/arcopen/fn/music/list"      // 音乐列表 获取媒体资源信息
	musicSearch      = "/arcopen/fn/music/search"    // 音乐搜索
)

// 图片上传url模块
const (
	// ImageUploadArcUrl 稿件图片上传
	ImageUploadArcUrl = "https://member.bilibili.com/arcopen/fn/archive/cover/upload"

	// ImageUploadCommodityUrl 商品图片上传
	ImageUploadCommodityUrl = "https://member.bilibili.com/arcopen/fn/v2/market/commodity/image/upload"

	// ImageUploadCustomer 客服图片上传
	ImageUploadCustomer = "https://member.bilibili.com/arcopen/fn/market/customer/image_upload"
)
