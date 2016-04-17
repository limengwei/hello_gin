package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ConnMap map[string]*websocket.Conn

func main() {
	openDB()

	ConnMap = make(map[string]*websocket.Conn)

	r := gin.Default()
	r.LoadHTMLGlob("res/*.html")
	r.Static("/static", "res/static")
	r.NoRoute(notfound)

	authorized := r.Group("/", interceptor)
	authorized.GET("/", index)
	authorized.GET("/art/:cate", art)
	authorized.Any("/editor", editor)
	authorized.Any("/sp", spider)
	authorized.Any("/chat", chat)
	authorized.Any("/ws/:uname", OnWebSocket)
	authorized.Any("/gitstar")

	r.GET("/logout", logout)
	r.Any("/login", login)
	r.Any("/oauth_github_callback", oauth_github_callback)

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
			setCookie("test", c)

			c.Redirect(302, "/")
		} else {
			c.HTML(200, "login.html", gin.H{"msg": "登陆失败"})
		}
	}

	if c.Request.Method == "GET" {
		c.HTML(200, "login.html", nil)
	}
}

//github授权回调
func oauth_github_callback(c *gin.Context) {
	var err error
	var resp *http.Response
	var res []byte

	code, exists := c.GetQuery("code")
	if exists {

		//请求token
		resp, err = http.Get("https://github.com/login/oauth/access_token?client_id=b270b5c94db796d87e22&client_secret=de82aebab00a7348ff658ef4432574d2ba087862&code=" + code)
		if err != nil {
			fmt.Println(err)
			return
		}

		//解析结果
		res, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		//获取用户信息
		resp, err = http.Get("https://api.github.com/user?" + string(res))
		if err != nil {
			fmt.Println(err)
			return
		}

		//解析结果
		res, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(res))
		c.Redirect(302, "/")

		defer resp.Body.Close()
	}
}

var arts []string = []string{"AAA", "BBB", "CCC", "CCC", "CCC", "CCC", "CCC"}

//首页（最新文章）
func index(c *gin.Context) {
	l, r := getPagetOption(c)
	cate := c.Params.ByName("cate")
	fmt.Println("--cate--", cate)

	//	as := make([]Article, 0)
	//	err := db.Find(&as)
	//	if err != nil {
	//		fmt.Println("--err--", err)
	//	} else {
	//		fmt.Println("--err--", len(as))
	//	}

	c.HTML(200, "index.html", gin.H{"titles": arts, "l": l, "r": r})
}

//文章（分类文章）
func art(c *gin.Context) {
	l, r := getPagetOption(c)
	cate := c.Params.ByName("cate")
	fmt.Println("--cate--", cate)
	//	as := make([]Article, 0)

	//	db.Find(&as)

	c.HTML(200, "index.html", gin.H{"titles": arts, "l": l, "r": r})
}

//编辑
func editor(c *gin.Context) {
	if c.Request.Method == "POST" {
		fmt.Println("title:", c.PostForm("title"))
		fmt.Println("content:", c.PostForm("content"))

		c.Redirect(302, "/")
	}

	if c.Request.Method == "GET" {
		c.HTML(200, "editor.html", nil)
	}
}

//生成cookie
func setCookie(cookieVal string, c *gin.Context) {
	cookiename := "cookieee"
	cookieval := cookieVal
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
