package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/io/fs"
	"github.com/labstack/echo/v4"
	"github.com/shurcooL/httpfs/html/vfstemplate"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func templates() *Template {
	var t *template.Template
	return &Template{
		templates: template.Must(
			vfstemplate.ParseGlob(fs.Files, t, "/templates/*.html"),
		),
	}
}

type Data struct {
	Input    string
	Parsed   []string
	HomePage bool
	Version  string
}

func NewData() *Data {
	return &Data{Version: gnparser.Version}
}

func home(gnps GNParserService) func(echo.Context) error {
	return func(c echo.Context) error {
		var parsed []string
		data := NewData()
		data.HomePage = true
		data.Input = c.QueryParam("q")
		names := strings.Split(data.Input, "\n")
		fmt.Println(len(names))
		for i := range names {
			if len(names[i]) == 0 {
				continue
			}
			p := gnps.ParseName(names[i]).Output(format.PrettyJSON)
			parsed = append(parsed, p)
		}
		data.Parsed = parsed
		return c.Render(http.StatusOK, "layout", data)
	}
}

func docAPI() func(echo.Context) error {
	return func(c echo.Context) error {
		data := NewData()
		return c.Render(http.StatusOK, "layout", data)
	}
}
