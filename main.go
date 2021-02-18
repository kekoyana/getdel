package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		data := struct {
			Message string
		}{
			Message: "にゃーん",
		}
		return c.Render(http.StatusOK, "hello", data)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
