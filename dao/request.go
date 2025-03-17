package dao

import (
	"fmt"
	"github.com/monaco-io/request"
	"openplat/model"
	"strconv"
	"time"
)

// ApiRequest http request demo方法
// reqJson 请求参数json
// requestUrl 请求url
// method 请求方法
// clientId 应用id
// accessToken 应用token
// appSecret 应用秘钥
// version 版本号 "1.0" "2.0"
func ApiRequest(reqJson, requestUrl, method, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	resp = model.BaseResp{}
	header := &model.CommonHeader{
		ContentType:       model.JsonType,
		ContentAcceptType: model.JsonType,
		Timestamp:         strconv.FormatInt(time.Now().Unix(), 10),
		SignatureMethod:   model.HmacSha256,
		SignatureVersion:  version,
		Authorization:     "",
		Nonce:             strconv.FormatInt(time.Now().UnixNano(), 10), //用于幂等,记得替换
		AccessKeyId:       clientId,
		ContentMD5:        Md5(reqJson),
		AccessToken:       accessToken,
	}
	header.Authorization = CreateSignature(header, appSecret)
	fmt.Println(requestUrl)
	fmt.Println(ToMap(header))

	cli := request.Client{
		Method: method,
		URL:    fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl),
		Header: ToMap(header),
		String: reqJson,
	}

	// 打印请求的cURL命令
	fmt.Println("cURL Command:")

	cliResp := cli.Send().Scan(&resp)
	if !cliResp.OK() {
		err = fmt.Errorf("[error] req:%+v resp:%+v err:%+v", reqJson, resp, cliResp.Error())
	}
	return
}
