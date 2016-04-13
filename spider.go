package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

//爬虫
func spider(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(200, "spider.html", nil)
	}

	if c.Request.Method == "POST" {
		url, exist := c.GetPostForm("urll")
		if exist {
			fmt.Println(url)
			crawl(url, c)
		} else {
			c.HTML(200, "spider.html", nil)
		}
	}

}

var ch chan int

func crawl(url string, c *gin.Context) {
	u_base := "http://www.cnblogs.com/#p"
	for i := 1; i < 2; i++ {
		u_pager := u_base + strconv.Itoa(i)
		crawlList(u_pager, c)
		time.Sleep(5 * time.Second)
	}
}

//分页

func crawlList(url string, c *gin.Context) {

	var titles []string
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
	}
	selection := doc.Find("#post_list").First().Find(".post_item")
	ch = make(chan int, selection.Size())

	selection.Each(func(i int, s *goquery.Selection) {
		time.Sleep(1 * time.Second)
		title := s.Find(".titlelnk").First()
		//summary := s.Find(".post_item_summary").First()
		titles = append(titles, title.Text())
		detailUrl, exists := title.Attr("href")

		if exists {

			//crawlDetail(detailUrl, c)
			a := new(Article)
			a.Title = title.Text()
			//a.summary = summary.Text()

			crawlDetail(a, detailUrl, c)
		}

		fmt.Printf("【%d----%s----%s】\n", i, title.Text(), detailUrl)

	})
}

//解析文章内容
func crawlDetail(a *Article, url string, c *gin.Context) {
	//	doc, err := goquery.NewDocument(url)
	//	if err != nil {
	//		fmt.Println(err)
	//	}

	//body := doc.Find("#cnblogs_post_body").First().Text()

	//fmt.Println("--文章长度--", len(body)) //字符串截取

	db.Debug().Create(a)
	//	fmt.Println("--NewRecord--", db.NewRecord(a)) //插入数据库
	//db.Exec("insert into article (title,body) values(?,?)", a.title, a.body)
}
