// Package web provides RESTful API service and a website for gnparser.
package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnlib/ent/nomcode"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/labstack/echo/v4"
)

// inputFORM is used to collect data from HTML form.
type inputFORM struct {
	Names             string `query:"names"           form:"names"`
	Code              string `query:"code"            form:"code"`
	Format            string `query:"format"          form:"format"`
	WithDetails       string `query:"with_details"    form:"with_details"`
	PreserveDiaereses string `query:"diaereses"       form:"diaereses"`
	CompactAuthors    string `query:"compact_authors" form:"compact_authors"`
	FlattenOutput     string `query:"flatten"         form:"flatten"`

	// WithCultivars is deprecated and overriden by Code
	WithCultivars string `query:"cultivars" form:"cultivars"`
}

// Data contains information required to render web-pages.
type Data struct {
	Input             string
	Parsed            []parsed.Parsed
	Code              string
	Format            string
	HomePage          bool
	Version           string
	WithDetails       bool
	PreserveDiaereses bool
	CompactAuthors    bool
	FlattenOutput     bool

	// WithCultivars is deprecated by Code field
	WithCultivars bool
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
	compactAuthors := inp.CompactAuthors == "on"
	flattenOutput := inp.FlattenOutput == "on"
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
	if compactAuthors {
		q.Set("compact_authors", inp.CompactAuthors)
	}
	if flattenOutput {
		q.Set("flatten", inp.FlattenOutput)
	}
	q.Set("code", inp.Code)

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
	data.CompactAuthors = inp.CompactAuthors == "on"
	data.FlattenOutput = inp.FlattenOutput == "on"
	data.Code = inp.Code

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
	if l := len(names); l > 0 {
		slog.Info("Parsed",
			"namesNum", l,
			"example", names[0],
			"parsedBy", "WEB GUI",
		)
	}
	data.Input = strings.Join(names, "\n")

	opts := []gnparser.Option{
		gnparser.OptWithDetails(data.WithDetails),
		gnparser.OptWithPreserveDiaereses(data.PreserveDiaereses),
		gnparser.OptWithCompactAuthors(data.CompactAuthors),
		gnparser.OptWithFlatOutput(data.FlattenOutput),
	}

	if data.WithCultivars {
		opts = append(opts, gnparser.OptCode(nomcode.Cultivars))
	}

	code := nomcode.New(data.Code)
	if code != nomcode.Unknown {
		// overrides data.WithCultivars
		opts = append(opts, gnparser.OptCode(code))
	}

	gnp := gnps.ChangeConfig(opts...)
	data.Parsed = gnp.ParseNames(names)

	switch data.Format {
	case "json":
		res := make([]string, len(data.Parsed))
		for i := range data.Parsed {
			res[i] = data.Parsed[i].Output(gnfmt.CompactJSON, data.FlattenOutput)
		}
		jsonArray := "[" + strings.Join(res, ",") + "]"
		return c.JSONBlob(http.StatusOK, []byte(jsonArray))
	case "csv", "tsv":
		f := gnfmt.CSV
		if data.Format == "tsv" {
			f = gnfmt.TSV
		}

		res := make([]string, len(data.Parsed)+1)
		res[0] = parsed.HeaderCSV(f, data.WithDetails)
		for i := range data.Parsed {
			res[i+1] = data.Parsed[i].Output(f, data.FlattenOutput)
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
