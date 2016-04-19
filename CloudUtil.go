package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"login"`
	Avatar string `json:"avatar_url"`
	Url    string `json:"url"`
}

type Result struct {
	Results []User `json:"results"`
}

const (
	AppID    = "8NSr8skDzWT3wEBLWsKnVwMe-gzGzoHsz"
	AppKey   = "D4Wsyd1ePzgjib1VmLBUdDf6"
	BaseURL  = "https://api.leancloud.cn/1.1"
	EmptyStr = ""
)

func users(c *gin.Context) {
	body, _ := sendGet(BaseURL+"/classes/Person", nil)

	bbody, _ := ioutil.ReadAll(body)

	fmt.Println(string(bbody))

	var result Result

	err := json.Unmarshal(bbody, &result)

	if err != nil {
		fmt.Println(err)
		return
	}

	var user User

	user = result.Results[0]

	fmt.Println(user)

}

//Get请求
func sendGet(url string, params url.Values) (io.Reader, error) {
	if params != nil {
		url = url + "?" + params.Encode()
	}
	fmt.Println(url)
	return sendRequest("GET", url, nil)
}

//Post请求
func sendPost(url string, params url.Values) (io.Reader, error) {

	return sendRequest("POST", url, params)
}

//通用请求
func sendRequest(method string, url string, params url.Values) (io.Reader, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(method, url, ioutil.NopCloser(strings.NewReader(params.Encode())))
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
