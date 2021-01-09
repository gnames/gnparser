package web

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gnames/gnparser/io/fs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const withLogs = true

func Run(gnps GNParserService) {
	log.Printf("Starting the HTTP on port %d.", gnps.Port())
	e := echo.New()
	e.Renderer = templates()
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
	if withLogs {
		e.Use(middleware.Logger())
	}
	e.GET("/", home(gnps))
	e.GET("/doc/api", docAPI())
	e.GET("/api", info())
	e.GET("/api/v1", info())
	e.GET("/api/v1/ping", ping(gnps))
	e.GET("/api/v1/version", ver(gnps))
	e.GET("/api/v1/:names", parseNamesGET(gnps))
	e.GET("/api/:names", parseNamesGET(gnps))
	e.POST("/api/v1", parseNamesPOST(gnps))
	e.POST("/api", parseNamesPOST(gnps))

	assetHandler := http.FileServer(fs.Files)
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

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

https://app.swaggerhub.com/apis-docs/dimus/gnparser/1.0.0`,
		)
	}
}

func ping(gnps GNParserService) func(echo.Context) error {
	return func(c echo.Context) error {
		result := gnps.Ping()
		return c.String(http.StatusOK, result)
	}
}

func ver(gnps GNParserService) func(echo.Context) error {
	return func(c echo.Context) error {
		result := gnps.GetVersion()
		return c.JSON(http.StatusOK, result)
	}
}

func parseNamesGET(gnps GNParserService) func(echo.Context) error {
	return func(c echo.Context) error {
		nameStr, _ := url.QueryUnescape(c.Param("names"))
		names := strings.Split(nameStr, "|")
		res := gnps.ParseNames(names)
		return c.JSON(http.StatusOK, res)
	}
}

func parseNamesPOST(gnps GNParserService) func(echo.Context) error {
	return func(c echo.Context) error {
		var names []string
		if err := c.Bind(&names); err != nil {
			return err
		}
		res := gnps.ParseNames(names)
		return c.JSON(http.StatusOK, res)
	}
}
