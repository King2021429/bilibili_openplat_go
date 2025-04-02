package dao

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"openplat/model"
	"os"
	"strconv"
	"time"
)

// ArticleAdd 文章提交
func ArticleAdd(requestUrl, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	title = title[:len(title)-1]
	_ = writer.WriteField("title", title)

	fmt.Print("请输入category: ")
	category, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	category = category[:len(category)-1]
	_ = writer.WriteField("category", category)

	fmt.Print("请输入template_id: ")
	templateId, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	templateId = templateId[:len(templateId)-1]
	_ = writer.WriteField("template_id", templateId)

	fmt.Print("请输入summary: ")
	summary, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	summary = summary[:len(summary)-1]
	_ = writer.WriteField("summary", summary)

	fmt.Print("请输入content: ")
	content, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	content = content[:len(content)-1]
	_ = writer.WriteField("content", content)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl), payload)
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

// ArticleEdit 文章编辑
func ArticleEdit(requestUrl, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入id: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	id = id[:len(id)-1]
	_ = writer.WriteField("id", id)

	fmt.Print("请输入title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	title = title[:len(title)-1]
	_ = writer.WriteField("title", title)

	fmt.Print("请输入category: ")
	category, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	category = category[:len(category)-1]
	_ = writer.WriteField("category", category)

	fmt.Print("请输入template_id: ")
	templateId, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	templateId = templateId[:len(templateId)-1]
	_ = writer.WriteField("template_id", templateId)

	fmt.Print("请输入summary: ")
	summary, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	summary = summary[:len(summary)-1]
	_ = writer.WriteField("summary", summary)

	fmt.Print("请输入content: ")
	content, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	content = content[:len(content)-1]
	_ = writer.WriteField("content", content)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl), payload)
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

// ArticleDel 文章删除
func ArticleDel(requestUrl, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入ids: ")
	ids, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	ids = ids[:len(ids)-1]
	_ = writer.WriteField("ids", ids)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl), payload)
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

// ArticleDetail 查询文章详情
func ArticleDetail(requestUrl, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入id: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	id = id[:len(id)-1]
	_ = writer.WriteField("id", id)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(model.MethodGet, fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl), payload)
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

// ArticleList 查询文章列表
func ArticleList(requestUrl, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入pn: ")
	pn, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	pn = pn[:len(pn)-1]
	_ = writer.WriteField("pn", pn)

	fmt.Print("请输入ps: ")
	ps, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	ps = ps[:len(ps)-1]
	_ = writer.WriteField("ps", ps)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(model.MethodGet, fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl), payload)
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

// ArticleCard 获取视频、文章卡片信息
func ArticleCard(requestUrl, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入resource_id: ")
	resourceId, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	resourceId = resourceId[:len(resourceId)-1]
	_ = writer.WriteField("resource_id", resourceId)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(model.MethodGet, fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl), payload)
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

// AnthologyAdd 文集提交
func AnthologyAdd(requestUrl, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	name = name[:len(name)-1]
	_ = writer.WriteField("name", name)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(model.MethodPost, fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl), payload)
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

// AnthologyEdit 文集信息编辑
func AnthologyEdit(requestUrl, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	name = name[:len(name)-1]
	_ = writer.WriteField("name", name)

	fmt.Print("请输入list_id: ")
	listId, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	listId = listId[:len(listId)-1]
	_ = writer.WriteField("list_id", listId)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(model.MethodPost, fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl), payload)
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

// AnthologyBelong 文集下文章列表修改
func AnthologyBelong(requestUrl, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入list_id: ")
	listId, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	listId = listId[:len(listId)-1]
	_ = writer.WriteField("list_id", listId)

	fmt.Print("请输入article_ids: ")
	articleIds, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	articleIds = articleIds[:len(articleIds)-1]
	_ = writer.WriteField("article_ids", articleIds)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(model.MethodPost, fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl), payload)
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

// AnthologyDelete 文集删除
func AnthologyDelete(requestUrl, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入id: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	id = id[:len(id)-1]
	_ = writer.WriteField("id", id)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(model.MethodPost, fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl), payload)
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

// AnthologyDetail 查询文集详情
func AnthologyDetail(requestUrl, clientId, accessToken, appSecret, version string) (resp model.BaseResp, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入id: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入时出错: %v", err)
	}
	id = id[:len(id)-1]
	_ = writer.WriteField("id", id)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(model.MethodPost, fmt.Sprintf("%s%s", model.UatMainOpenPlatformHttpHost, requestUrl), payload)
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
