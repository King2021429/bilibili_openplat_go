package service

import (
	"bufio"
	"fmt"
	"log"
	"openplat/dao"
	"openplat/model"
	"os"
	"strconv"
	"time"
)

func Sign(clientId, accessToken, appSecret string) {
	reader := bufio.NewReader(os.Stdin)
	// 输入 access_token
	fmt.Print("请输入reqJson串: ")
	reqJson, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]

	// 输入 version
	fmt.Print("请输入version: ")
	version, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	reqJson = reqJson[:len(reqJson)-1]
	header := &model.CommonHeader{
		ContentType:       model.JsonType,
		ContentAcceptType: model.JsonType,
		Timestamp:         strconv.FormatInt(time.Now().Unix(), 10),
		SignatureMethod:   model.HmacSha256,
		SignatureVersion:  version,
		Authorization:     "",
		Nonce:             strconv.FormatInt(time.Now().UnixNano(), 10), //用于幂等,记得替换
		AccessKeyId:       clientId,
		ContentMD5:        dao.Md5(reqJson),
		AccessToken:       accessToken,
	}
	header.Authorization = dao.CreateSignature(header, appSecret)
	fmt.Printf("\n请求头:%s", dao.ToMap(header))
	return
}

func Oauth(clientId, appSecret string) {
	// 基础 URL
	baseURL := "https://api.bilibili.com/x/account-oauth2/v1/token"

	reader := bufio.NewReader(os.Stdin)

	// 输入 client_id
	fmt.Print("请输入 code: ")
	code, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	code = code[:len(code)-1]
	// 参数字典
	params := map[string]string{
		"client_id":     clientId,
		"client_secret": appSecret,
		"grant_type":    "authorization_code",
		"code":          code,
	}

	// 调用拼接函数
	fullURL, err := dao.BuildURL(baseURL, params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 打印拼接好的完整 URL
	fmt.Println("Full URL:", fullURL)
	dao.OauthReq(fullURL)

}
