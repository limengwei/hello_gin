package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("res/*.html")
	r.Static("/static", "res/static")

	authorized := r.Group("/", interceptor)

	authorized.GET("/", index)

	//	authorized.GET("/blog", home)
	//	authorized.GET("/editor", editor)
	//	authorized.POST("/putlog", putlog)

	r.Any("/login", login)

	r.Run()
}

func interceptor(c *gin.Context) {
	_, err := c.GetCookie("cookieee")

	if err != nil {
		c.HTML(200, "login.html", gin.H{"msg": "登陆失败"})
		return
	}
	c.Redirect(200, "/")
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func login(c *gin.Context) {
	if c.Request.Method == "POST" {
		if c.PostForm("name") == "admin" {
			c.SetCookie("cookieee", "cookiee-val", 1024, "", "localhost", true, true)
			c.Redirect(200, "/index")
		} else {
			c.HTML(200, "login.html", gin.H{"msg": "登陆失败"})
		}
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func home(c *gin.Context) {
	titles := []string{"AAA", "BBB", "CCC"}
	c.HTML(200, "index.html", gin.H{"titles": titles})
}

func editor(c *gin.Context) {

	c.HTML(200, "editor.html", nil)
}

func putlog(c *gin.Context) {

	fmt.Println("title:", c.PostForm("title"))
	fmt.Println("content:", c.PostForm("content"))

	c.Redirect(302, "/")
}
