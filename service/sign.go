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

func Oauth(clientId, accessToken, appSecret string) {

}
