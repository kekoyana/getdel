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
	Tweets := []db.Tweet{}
	conn.Order("id desc").Find(&Tweets)
	return c.Render(http.StatusOK, "index", Tweets)
}

func postTweet(c echo.Context) error {
	data := struct {
		Message string
	}{
		Message: "にゃーん",
	}
	return c.Render(http.StatusOK, "index", data)
}

func delTweet(c echo.Context) error {
	data := struct {
		Message string
	}{
		Message: "にゃーん",
	}
	return c.Render(http.StatusOK, "index", data)
}

func main() {
	db.DbInit()

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", getTweet)
	e.GET("/post", getTweet)
	e.GET("/del", delTweet)
	e.Logger.Fatal(e.Start(":1323"))
}
