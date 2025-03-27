package dao

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func OauthReq(url string) {
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "api.bilibili.com")
	req.Header.Add("Connection", "keep-alive")

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
}

// BuildURL 拼接 URL 参数的函数
func BuildURL(baseURL string, params map[string]string) (string, error) {
	// 解析基础 URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// 获取 URL 的查询参数
	query := u.Query()

	// 遍历参数映射，将参数添加到查询参数中
	for key, value := range params {
		query.Add(key, value)
	}

	// 将更新后的查询参数设置回 URL
	u.RawQuery = query.Encode()

	// 返回拼接好的完整 URL
	return u.String(), nil
}
