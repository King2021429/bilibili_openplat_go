package dao

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func oauthReq() {
	url := "https://uat-api.bilibili.com/x/account-oauth2/v1/token?client_id=93a0f774fae84e6c&client_secret=3af0cfc8d9fb4d698b08c8a3782cb86b&grant_type=authorization_code&code=60f5ac0c1c9a4c1280e8f36f6043083b"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "uat-api.bilibili.com")
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
