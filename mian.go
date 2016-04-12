package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("res/*.html")
	r.Static("/static", "res/static")
	r.NoRoute(notfound)

	authorized := r.Group("/", interceptor)
	authorized.GET("/", index)
	authorized.GET("/art/:cate", art)
	authorized.Any("/editor", editor)
	authorized.Any("/sp", spider)

	r.GET("/logout", logout)
	r.Any("/login", login)

	r.Run(":80")
}

//拦截器
func interceptor(c *gin.Context) {
	cookie, err := c.Cookie("cookieee")

	if err != nil {
		c.Redirect(302, "/login")
	} else {
		if cookie != "1024" {
			c.Redirect(302, "/login")
		}
	}

}

//登录
func login(c *gin.Context) {
	if c.Request.Method == "POST" {
		if c.PostForm("name") == "admin" && c.PostForm("pwd") == "admin" {
			setCookie(c)

			c.Redirect(302, "/")
		} else {
			c.HTML(200, "login.html", gin.H{"msg": "登陆失败"})
		}
	}

	if c.Request.Method == "GET" {
		c.HTML(200, "login.html", nil)
	}

}

var arts []string = []string{"AAA", "BBB", "CCC", "CCC", "CCC", "CCC", "CCC"}

//首页（最新文章）
func index(c *gin.Context) {
	l, r := getPagetOption(c)
	cate := c.Params.ByName("cate")
	fmt.Println("--cate--", cate)

	c.HTML(200, "index.html", gin.H{"titles": arts, "l": l, "r": r})
}

//文章（分类文章）
func art(c *gin.Context) {
	l, r := getPagetOption(c)
	cate := c.Params.ByName("cate")
	fmt.Println("--cate--", cate)

	c.HTML(200, "index.html", gin.H{"titles": arts, "l": l, "r": r})
}

//编辑
func editor(c *gin.Context) {
	if c.Request.Method == "POST" {
		fmt.Println("title:", c.PostForm("title"))
		fmt.Println("content:", c.PostForm("content"))

		arts = append(arts, c.PostForm("title"))
		c.Redirect(302, "/")
	}

	if c.Request.Method == "GET" {
		c.HTML(200, "editor.html", nil)
	}
}

//生成cookie
func setCookie(c *gin.Context) {
	cookiename := "cookieee"
	cookieval := "1024"
	maxAge := 102400
	path := ""
	domain := "localhost"
	secure := false
	httpOnly := false
	c.SetCookie(cookiename, cookieval, maxAge, path, domain, secure, httpOnly)
}

//退出登录(cookie置空)
func logout(c *gin.Context) {
	cookiename := "cookieee"
	cookieval := ""
	maxAge := 102400
	path := ""
	domain := "localhost"
	secure := false
	httpOnly := false
	c.SetCookie(cookiename, cookieval, maxAge, path, domain, secure, httpOnly)

	c.Redirect(302, "/login")
}

//分页
func getPagetOption(c *gin.Context) (int, int) {

	var ip, l, r int
	p, _ := c.GetQuery("p")

	ip, _ = strconv.Atoi(p)

	l = ip - 1
	r = ip + 1

	if l < 1 {
		l = 1
		r = 2
	}
	return l, r
}

//404
func notfound(c *gin.Context) {
	c.HTML(404, "404.html", nil)
}
