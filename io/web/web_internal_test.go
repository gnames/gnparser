package web

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gnames/gnlib/domain/entity/gn"
	"github.com/gnames/gnlib/encode"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func handlerGET(path string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	e.Renderer = templates()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestHome(t *testing.T) {
	cfg := gnparser.NewConfig(gnparser.OptFormat("compact"))
	gnp := gnparser.New(cfg)
	gnps := NewGNparserService(gnp, 0)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	e.Renderer = templates()

	assert.Nil(t, home(gnps)(c))
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Contains(t, rec.Body.String(), "GNA Logo")
}

func TestDocAPI(t *testing.T) {
	c, rec := handlerGET("/doc/api")
	assert.Nil(t, docAPI()(c))
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Contains(t, rec.Body.String(), "Application Programming Interface")
}

func TestInfo(t *testing.T) {
	c, rec := handlerGET("/")

	assert.Nil(t, info()(c))
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Contains(t, rec.Body.String(), "OpenAPI")
}

func TestPing(t *testing.T) {
	cfg := gnparser.NewConfig(gnparser.OptFormat("compact"))
	gnp := gnparser.New(cfg)
	gnps := NewGNparserService(gnp, 0)
	c, rec := handlerGET("/ping")

	assert.Nil(t, ping(gnps)(c))
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Equal(t, rec.Body.String(), "pong")
}

func TestVer(t *testing.T) {
	cfg := gnparser.NewConfig(gnparser.OptFormat("compact"))
	gnp := gnparser.New(cfg)
	gnps := NewGNparserService(gnp, 0)
	c, rec := handlerGET("/version")

	assert.Nil(t, ver(gnps)(c))
	enc := encode.GNjson{}
	var response gn.Version
	err := enc.Decode(rec.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Regexp(t, `^v\d+\.\d+\.\d+`, response.Version)
}

func TestParseGET(t *testing.T) {
	cfg := gnparser.NewConfig(gnparser.OptFormat("compact"))
	gnp := gnparser.New(cfg)
	gnps := NewGNparserService(gnp, 0)

	var response []parsed.Parsed
	names := []string{
		"Not name", "Bubo bubo", "Pomatomus",
		"Pardosa moesta", "Plantago major var major",
		"Cytospora ribis mitovirus 2",
		"A-shaped rods", "Alb. alba",
		"Pisonia grandis", "Acacia vestita may",
	}
	request := strings.Join(
		names,
		"|",
	)
	namesQuery := url.QueryEscape(request)
	path := "/" + namesQuery

	c, rec := handlerGET(path)
	c.SetPath("/:names")
	c.SetParamNames("names")
	c.SetParamValues(namesQuery)

	assert.Nil(t, parseNamesGET(gnps)(c))

	enc := encode.GNjson{}
	err := enc.Decode(rec.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, len(response), len(names))
	for i, v := range response {
		switch i {
		case 0:
			assert.Equal(t, v.Verbatim, "Not name", v.Verbatim)
			assert.False(t, v.Parsed, v.Verbatim)
		case 1:
			assert.Equal(t, v.Verbatim, "Bubo bubo", v.Verbatim)
			assert.Equal(t, v.Canonical.Simple, "Bubo bubo")
		}
	}
}

func TestParsePOST(t *testing.T) {
	cfg := gnparser.NewConfig(gnparser.OptFormat("compact"))
	gnp := gnparser.New(cfg)
	gnps := NewGNparserService(gnp, 0)

	var response []parsed.Parsed
	names := []string{
		"Not name", "Bubo bubo", "Pomatomus",
		"Pardosa moesta", "Plantago major var major",
		"Cytospora ribis mitovirus 2",
		"A-shaped rods", "Alb. alba",
		"Pisonia grandis", "Acacia vestita may",
	}
	params := inputPOST{
		Names:       names,
		CSV:         false,
		WithDetails: false,
	}
	reqBody, err := encode.GNjson{}.Encode(params)
	assert.Nil(t, err)
	r := bytes.NewReader(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	assert.Nil(t, parseNamesPOST(gnps)(c))

	enc := encode.GNjson{}
	err = enc.Decode(rec.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, len(response), len(names))
	for i, v := range response {
		switch i {
		case 0:
			assert.Equal(t, v.Verbatim, "Not name", v.Verbatim)
			assert.False(t, v.Parsed, v.Verbatim)
		case 1:
			assert.Equal(t, v.Verbatim, "Bubo bubo", v.Verbatim)
			assert.Equal(t, v.Canonical.Simple, "Bubo bubo")
		}
	}

	params = inputPOST{
		Names:       names,
		CSV:         true,
		WithDetails: false,
	}
	reqBody, err = encode.GNjson{}.Encode(params)
	r = bytes.NewReader(reqBody)
	req = httptest.NewRequest(http.MethodPost, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	assert.Nil(t, parseNamesPOST(gnps)(c))
	assert.True(t, strings.HasPrefix(rec.Body.String(), "Id"))
}
