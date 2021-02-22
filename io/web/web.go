// Package web provides RESTful API service and a website for gnparser.
package web

import (
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/io/fs"
	"github.com/labstack/echo/v4"
	"github.com/shurcooL/httpfs/html/vfstemplate"
)

// echoTempl implements echo.Renderer interface.
type echoTempl struct {
	templates *template.Template
}

// Render implements echo.Renderer interface.
func (t *echoTempl) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func templates() *echoTempl {
	var t *template.Template
	return &echoTempl{
		templates: template.Must(
			vfstemplate.ParseGlob(fs.Files, t, "/templates/*.html"),
		),
	}
}

// Data contains information required to render web-pages.
type Data struct {
	Input    string
	Parsed   []string
	HomePage bool
	Version  string
}

// NewData creates new Data for web-page templates.
func NewData() *Data {
	return &Data{Version: gnparser.Version}
}

func home(gnps GNparserService) func(echo.Context) error {
	gnp := gnps.ChangeConfig(gnparser.OptWithDetails(true))
	return func(c echo.Context) error {
		var parsed []string
		data := NewData()
		data.HomePage = true
		data.Input = c.QueryParam("q")
		names := strings.Split(data.Input, "\n")
		for i := range names {
			if len(names[i]) == 0 {
				continue
			}
			p := gnp.ParseName(names[i]).Output(gnfmt.PrettyJSON)
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
