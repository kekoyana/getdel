package main

import (
	"html/template"
	"io"
	"net/http"

	db "getdel/lib"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getTweet(c echo.Context) error {
	conn := db.DbConn()
	tweet := db.Tweet{}
	conn.Create(&tweet)
	Tweets := []db.Tweet{}
	conn.Order("id desc").Find(&Tweets)
	return c.Render(http.StatusOK, "index", Tweets)
}

func postTweet(c echo.Context) error {
	conn := db.DbConn()
	tweet := db.Tweet{}
	//tweet.Message = c.FormValue("body")
	tweet.Message = "Hoge"
	conn.Create(&tweet)
	return c.Redirect(http.StatusFound, "/")
}

func delTweet(c echo.Context) error {
	conn := db.DbConn()
	tweet := db.Tweet{}
	conn.First(&tweet)
	conn.Delete(&tweet)
	return c.Redirect(http.StatusFound, "/")
}

func main() {
	db.DbInit()

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", getTweet)
	e.GET("/post", postTweet)
	e.GET("/del", delTweet)
	e.Logger.Fatal(e.Start(":1323"))
}
