package web

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnlib/ent/nomcode"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed static
var static embed.FS

type inputREST struct {
	Names             []string `json:"names"`
	CSV               bool     `json:"csv"`
	WithDetails       bool     `json:"withDetails"`
	PreserveDiaereses bool     `json:"preserveDiaereses"`
	NoSpacedInitials  bool     `json:"noSpacedInitials"`
	FlatOutput        bool     `json:"flatOutput"`
	Code              string   `json:"code"`

	// WithCultivars is deprecated by Code and overriden by it
	WithCultivars bool `json:"withCultivars"`
}

// Run starts the GNparser web service and servies both RESTful API and
// a website.
func Run(gnps GNparserService) {
	var err error

	e := echo.New()

	e.Renderer, err = NewTemplate()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	e.GET("/", homeGET(gnps))
	e.POST("/", homePOST(gnps))
	e.GET("/doc/api", docAPI())
	e.GET("/api", info())
	e.GET("/api/v1", info())
	e.GET("/api/v1/ping", ping(gnps))
	e.GET("/api/v1/version", ver(gnps))
	e.GET("/api/v1/:names", parseNamesGET(gnps))
	e.GET("/api/:names", parseNamesGET(gnps))
	e.POST("/api/v1/", parseNamesPOST(gnps))
	e.POST("/api/", parseNamesPOST(gnps))

	fs := http.FileServer(http.FS(static))
	e.GET("/static/*", echo.WrapHandler(fs))

	addr := fmt.Sprintf(":%d", gnps.Port())
	s := &http.Server{
		Addr:         addr,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}

func info() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(
			http.StatusOK,
			`OpenAPI for gnparser is described at

https://apidoc.globalnames.org/gnparser`,
		)
	}
}

func ping(gnps GNparserService) func(echo.Context) error {
	return func(c echo.Context) error {
		result := gnps.Ping()
		return c.String(http.StatusOK, result)
	}
}

func ver(gnps GNparserService) func(echo.Context) error {
	return func(c echo.Context) error {
		version := gnps.GetVersion()
		return c.JSON(http.StatusOK, version)
	}
}

func parseNamesGET(gnps GNparserService) func(echo.Context) error {
	return func(c echo.Context) error {
		nameStr, _ := url.QueryUnescape(c.Param("names"))
		csv := c.QueryParam("csv") == "true"
		det := c.QueryParam("with_details") == "true"
		cultivars := c.QueryParam("cultivars") == "true"
		diaereses := c.QueryParam("diaereses") == "true"
		initials := c.QueryParam("no_spaced_initials") == "true"
		flatten := c.QueryParam("flatten") == "true"
		codeStr := c.QueryParam("code")

		code := getCode(codeStr, cultivars)

		gnp := gnps.ChangeConfig(opts(code, csv, det, diaereses, initials, flatten)...)
		names := strings.Split(nameStr, "|")
		res := gnp.ParseNames(names)
		if l := len(names); l > 0 {
			slog.Info("Parsed",
				"namesNum", l, "example", names[0],
				"parsedBy", "REST API", "method", "GET",
			)
		}
		return formatNames(c, res, gnp)
	}
}

func parseNamesPOST(gnps GNparserService) func(echo.Context) error {
	return func(c echo.Context) error {
		var input inputREST
		if err := c.Bind(&input); err != nil {
			return err
		}

		if l := len(input.Names); l > 0 {
			slog.Info("Parsed",
				"namesNum", l,
				"example", input.Names[0],
				"parsedBy", "REST API",
				"method", "POST",
			)
		}
		code := getCode(input.Code, input.WithCultivars)

		gnp := gnps.ChangeConfig(
			opts(code,
				input.CSV,
				input.WithDetails,
				input.PreserveDiaereses,
				input.NoSpacedInitials,
				input.FlatOutput,
			)...)
		res := gnp.ParseNames(input.Names)
		return formatNames(c, res, gnp)
	}
}

func getCode(codeStr string, cultivars bool) nomcode.Code {
	code := nomcode.Unknown
	if cultivars {
		code = nomcode.Cultivars
	}
	code2 := nomcode.New(codeStr)
	if code2 == nomcode.Unknown {
		return code
	}
	return code2
}

func formatNames(
	c echo.Context,
	res []parsed.Parsed,
	gnp gnparser.GNparser,
) error {
	f := gnp.Format()
	switch f {
	case gnfmt.CSV, gnfmt.TSV:
		resCSV := make([]string, 0, len(res)+1)
		resCSV = append(resCSV, parsed.HeaderCSV(f, gnp.WithDetails()))
		for i := range res {
			resCSV = append(resCSV, res[i].Output(f, gnp.FlatOutput()))
		}
		return c.String(http.StatusOK, strings.Join(resCSV, "\n"))
	default:
		resJSON := make([]string, len(res))
		for i := range res {
			resJSON[i] = res[i].Output(f, gnp.FlatOutput())
		}
		str := "[" + strings.Join(resJSON, ",") + "]"
		return c.JSONBlob(http.StatusOK, []byte(str))
	}
}

func opts(code nomcode.Code, csv, details, diaereses,
	initials, flatten bool) []gnparser.Option {
	res := []gnparser.Option{
		gnparser.OptWithDetails(details),
		gnparser.OptCode(code),
		gnparser.OptWithPreserveDiaereses(diaereses),
		gnparser.OptWithCompactAuthors(initials),
		gnparser.OptWithFlatOutput(flatten),
	}
	if csv {
		res = append(res, gnparser.OptFormat(gnfmt.CSV))
	} else {
		res = append(res, gnparser.OptFormat(gnfmt.CompactJSON))
	}

	return res
}
