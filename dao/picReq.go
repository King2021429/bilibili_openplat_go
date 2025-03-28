package dao

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"openplat/model"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// PicRequest 图片上传请求方法
func PicRequest(requestUrl, picUrl, clientId, accessToken, appSecret, version string, flag int64) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	// picUrl eg: "/Users/shenyue/Downloads/test.jpg"

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入picUrl串: ")
	picUrl, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	picUrl = picUrl[:len(picUrl)-1]

	file, errFile1 := os.Open(picUrl)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base(picUrl))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}

	var staffId string
	if flag == 4 {
		fmt.Print("请输入staff_id: ")
		staffId, err = reader.ReadString('\n')
		if err != nil {
			log.Fatalf("读取输入时出错: %v", err)
		}
		staffId = staffId[:len(staffId)-1]
		_ = writer.WriteField("staff_id", staffId)
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", requestUrl, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	bodyStrMds := Md5("")

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-bili-timestamp", timestamp)
	req.Header.Add("x-bili-signature-method", "HMAC-SHA256")
	req.Header.Add("x-bili-signature-version", version)
	req.Header.Add("x-bili-signature-nonce", nonce)
	req.Header.Add("x-bili-accesskeyid", clientId)
	req.Header.Add("x-bili-content-md5", bodyStrMds)
	req.Header.Add("access-token", accessToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	header := &model.CommonHeader{
		ContentType:       writer.FormDataContentType(),
		ContentAcceptType: model.JsonType,
		Timestamp:         timestamp,
		SignatureMethod:   model.HmacSha256,
		SignatureVersion:  version,
		Authorization:     "",
		Nonce:             nonce, //用于幂等,记得替换
		AccessKeyId:       clientId,
		ContentMD5:        bodyStrMds,
		AccessToken:       accessToken,
	}

	req.Header.Add("Authorization", CreateSignature(header, appSecret))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	return
}
