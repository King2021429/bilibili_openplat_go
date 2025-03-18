package model

// OpenMarketCustomerSendMsgReq 外部请求开平 消息发送请求
type OpenMarketCustomerSendMsgReq struct {
	BizId          int64  `json:"biz_id" form:"biz_id" validate:"required"`     // 业务方类型id: 1-商家，2-带货,3-经营号
	ConversationId int64  `json:"conversation_id" form:"conversation_id"`       // 会话id
	StaffId        int64  `json:"staff_id" form:"staff_id" validate:"required"` // 客服id
	UserOpenId     string `json:"user_open_id"`                                 // 用户openid
	ShopId         int64  `json:"shop_id" form:"shop_id" validate:"required"`   // 商家id
	MsgType        int64  `json:"msg_type" form:"msg_type" validate:"required"` // 消息类型
	Msg            string `json:"msg" form:"msg" validate:"required"`           // 消息内容
}

// OpenMarketCustomerSendMsgResp 开平对外 消息发送返回结构体
type OpenMarketCustomerSendMsgResp struct {
	MsgKey string `json:"msg_key"` // 消息key
}

// OpenMarketCustomerConversationCloseReq 客服关闭会话请求
type OpenMarketCustomerConversationCloseReq struct {
	ConversationId int64 `json:"conversation_id" form:"conversation_id" validate:"required"` // 会话id
	StaffId        int64 `json:"staff_id" form:"staff_id" validate:"required"`               // 客服id
}

// OpenMarketCustomerConversationCloseResp 客服关闭会话返回参数
type OpenMarketCustomerConversationCloseResp struct{}

// OpenMarketCustomerStaffStatusUpdateReq 客服修改状态 请求参数
type OpenMarketCustomerStaffStatusUpdateReq struct {
	Status              int64 `json:"status" form:"status" validate:"required"`     // 客服状态,1-上线,2-忙碌,3-离线
	StaffId             int64 `json:"staff_id" form:"staff_id" validate:"required"` // 客服id
	IsCloseConversation int64 `json:"is_close_conversation"`                        // 离线是否关闭会话  1-是 0-否,默认0
}

// OpenMarketCustomerStaffStatusUpdateResp 客服修改状态 返回参数
type OpenMarketCustomerStaffStatusUpdateResp struct{}

// MarketCustomerImageUploadReq 客服系统 图片上传
type MarketCustomerImageUploadReq struct {
	StaffId int64 `json:"staff_id" form:"staff_id" validate:"required"` // 客服id
}

// MarketCustomerImageUploadResp 客服系统 图片上传 返回url
type MarketCustomerImageUploadResp struct {
	Url string `json:"url"` // 图片url
}
