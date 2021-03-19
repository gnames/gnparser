// Package web provides RESTful API service and a website for gnparser.
package web

import (
	"net/http"
	"strings"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnparser"
	"github.com/labstack/echo/v4"
)

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
