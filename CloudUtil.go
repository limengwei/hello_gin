package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/ChimeraCoder/gojson"
	"github.com/gin-gonic/gin"
)

type Pserson struct {
	Id     int    `json:"id"`
	Name   string `json:"login"`
	Avatar string `json:"avatar_url"`
	Url    string `json:"url"`
}

const (
	AppID    = "8NSr8skDzWT3wEBLWsKnVwMe-gzGzoHsz"
	AppKey   = "D4Wsyd1ePzgjib1VmLBUdDf6"
	BaseURL  = "https://api.leancloud.cn/1.1"
	EmptyStr = ""
)

func users(c *gin.Context) {
	body, _ := sendGet(BaseURL+"/classes/_User", nil)

	c.JSON(200, "users")

	person, err := json2struct.Generate(body, "Pserson", "main")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(person)
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
	client := &http.Client{}
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
	//defer res.Body.Close()

	return res.Body, nil
}
