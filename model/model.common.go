package model

type BaseResp struct {
	Code      int64       `json:"code"`
	Message   string      `json:"message"`
	RequestId string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

// 域名
const (
	// MainOpenPlatformHttpHost 主站开平
	MainOpenPlatformHttpHost = "https://member.bilibili.com"
)

const (
	SceneCode = "ARC_APP_SHARE"
)
