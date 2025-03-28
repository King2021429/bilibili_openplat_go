package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"openplat/dao"
	"openplat/model"
	"os"
)

// 唯一一个v1签名

// CommonAddShare 新增共享
func CommonAddShare(clientId, accessToken, appSecret string) (resp model.BaseResp, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入biz_code串: ")
	bizCode, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	bizCode = bizCode[:len(bizCode)-1]

	fmt.Print("请输入source: ")
	source, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	source = source[:len(source)-1]

	fmt.Print("请输入cover: ")
	cover, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	cover = cover[:len(cover)-1]

	// 创建一个 CommonMsg 对象
	commonMsg := model.CommonMsg{
		Source:  source,
		Cover:   cover,
		Title:   "投稿标题",
		TypeID:  172,
		TopicID: 1173894,
		VideoMaterialURL: []string{
			"https://1400335750.vod2.myqcloud.com/ff539f7evodcq1400335750/4d9970f01397757888517111156/QGeorKmhB2kA.mp4",
		},
	}

	// 将对象编码为 JSON 字符串
	jsonString, err := json.Marshal(commonMsg)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(jsonString)
	commonAddShareReq := model.CommonAddShareReq{
		CommonMsg: string(jsonString),
		BizCode:   bizCode,
		SceneCode: model.SceneCode,
	}
	commonAddShareReqJson, _ := json.Marshal(commonAddShareReq)
	return dao.ApiRequest(string(commonAddShareReqJson), model.ResourceAddShareUrl, model.MethodPost, clientId, accessToken, appSecret, model.BiliVersion)
}
