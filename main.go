package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, e echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	fmt.Println("iam working")

	t := &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}

	e := echo.New()
	e.Static("/static", "static")
	e.Renderer = t
	e.GET("/", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index", nil)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
