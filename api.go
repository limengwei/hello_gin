package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"login"`
	Avatar string `json:"avatar_url"`
	Url    string `json:"url"`
}

type Result struct {
	Results []User `json:"results"`
	Count   int    `json:"count"`
}

const (
	AppID   = "8NSr8skDzWT3wEBLWsKnVwMe-gzGzoHsz"
	AppKey  = "D4Wsyd1ePzgjib1VmLBUdDf6"
	BaseURL = "https://api.leancloud.cn/1.1"
)

//查询用户 是否存在
func userCount(u User) int {
	params := url.Values{}
	params.Set("count", "1")
	params.Set("limit", "0")
	params.Set("where", "{\"id\":"+strconv.Itoa(u.Id)+"}")

	body, _ := sendGet(BaseURL+"/classes/Person", params)

	bbody, _ := ioutil.ReadAll(body)

	fmt.Println(string(bbody))

	var result Result

	err := json.Unmarshal(bbody, &result)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.Count)

	if result.Count == 0 {
		userRegister(u)
	}

}

//创建用户
func userRegister(u User) {

	b, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, _ := sendPost(BaseURL+"/classes/Person", string(b))

	bbody, _ := ioutil.ReadAll(body)

	fmt.Println(string(bbody))

}

//Get请求
func sendGet(url string, params url.Values) (io.Reader, error) {
	if params != nil {
		url = url + "?" + params.Encode()
	}
	fmt.Println(url)
	return sendRequest("GET", url, "")
}

//Post请求
func sendPost(url string, jsonParams string) (io.Reader, error) {

	return sendRequest("POST", url, jsonParams)
}

//通用请求
func sendRequest(method string, url string, paramsEncode string) (io.Reader, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(method, url, ioutil.NopCloser(strings.NewReader(paramsEncode)))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-LC-Id", AppID)
	req.Header.Set("X-LC-Key", AppKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res.Body, nil
}
