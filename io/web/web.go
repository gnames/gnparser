// Package web provides RESTful API service and a website for gnparser.
package web

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/labstack/echo/v4"
)

// inputFORM is used to collect data from HTML form.
type inputFORM struct {
	Names             string `query:"names" form:"names"`
	Format            string `query:"format" form:"format"`
	WithDetails       string `query:"with_details" form:"with_details"`
	WithCultivars     string `query:"cultivars" form:"cultivars"`
	PreserveDiaereses string `query:"diaereses" form:"diaereses"`
}

// Data contains information required to render web-pages.
type Data struct {
	Input             string
	Parsed            []parsed.Parsed
	Format            string
	HomePage          bool
	Version           string
	WithDetails       bool
	WithCultivars     bool
	PreserveDiaereses bool
}

// NewData creates new Data for web-page templates.
func newData(isHome bool) *Data {
	return &Data{HomePage: isHome, Format: "html", Version: gnparser.Version}
}

func homePOST(gnps GNparserService) func(echo.Context) error {
	return func(c echo.Context) error {
		inp := new(inputFORM)
		data := newData(true)

		err := c.Bind(inp)
		if err != nil {
			return err
		}

		if strings.TrimSpace(inp.Names) == "" {
			return c.Redirect(http.StatusFound, "")
		}

		if strings.Count(inp.Names, "\n") < 20 {
			return redirectToHomeGET(c, inp)
		}

		return parsingResults(c, gnps, inp, data)
	}
}

func redirectToHomeGET(c echo.Context, inp *inputFORM) error {
	withDetails := inp.WithDetails == "on"
	withCultivars := inp.WithCultivars == "on"
	preserveDiaereses := inp.PreserveDiaereses == "on"
	q := make(url.Values)
	q.Set("names", inp.Names)
	q.Set("format", inp.Format)
	if withDetails {
		q.Set("with_details", inp.WithDetails)
	}
	if withCultivars {
		q.Set("cultivars", inp.WithCultivars)
	}
	if preserveDiaereses {
		q.Set("diaereses", inp.PreserveDiaereses)
	}

	url := fmt.Sprintf("/?%s", q.Encode())
	return c.Redirect(http.StatusFound, url)
}

func homeGET(gnps GNparserService) func(echo.Context) error {
	return func(c echo.Context) error {
		data := newData(true)

		inp := new(inputFORM)
		err := c.Bind(inp)
		if err != nil {
			return err
		}

		if strings.TrimSpace(inp.Names) == "" {
			return c.Render(http.StatusOK, "layout", data)
		}

		return parsingResults(c, gnps, inp, data)
	}
}

func parsingResults(
	c echo.Context,
	gnps GNparserService,
	inp *inputFORM,
	data *Data,
) error {
	var names []string
	data.WithDetails = inp.WithDetails == "on"
	data.WithCultivars = inp.WithCultivars == "on"
	data.PreserveDiaereses = inp.PreserveDiaereses == "on"

	format := inp.Format
	if format == "csv" || format == "tsv" || format == "json" {
		data.Format = format
	}

	data.Input = strings.TrimSpace(inp.Names)
	split := strings.Split(data.Input, "\n")
	if len(split) > 5_000 {
		split = split[0:5_000]
	}

	names = make([]string, len(split))
	for i := range split {
		names[i] = strings.TrimSpace(split[i])
	}
	data.Input = strings.Join(names, "\n")

	opts := []gnparser.Option{
		gnparser.OptWithDetails(data.WithDetails),
		gnparser.OptWithCultivars(data.WithCultivars),
		gnparser.OptWithPreserveDiaereses(data.PreserveDiaereses),
	}

	gnp := gnps.ChangeConfig(opts...)
	data.Parsed = gnp.ParseNames(names)

	switch data.Format {
	case "json":
		return c.JSON(http.StatusOK, data.Parsed)
	case "csv", "tsv":
		f := gnfmt.CSV
		if data.Format == "tsv" {
			f = gnfmt.TSV
		}

		res := make([]string, len(data.Parsed)+1)
		res[0] = parsed.HeaderCSV(f)
		for i := range data.Parsed {
			res[i+1] = data.Parsed[i].Output(f)
		}
		return c.String(http.StatusOK, strings.Join(res, "\n"))
	default:
		return c.Render(http.StatusOK, "layout", data)
	}
}

func docAPI() func(echo.Context) error {
	return func(c echo.Context) error {
		data := newData(false)
		return c.Render(http.StatusOK, "layout", data)
	}
}
