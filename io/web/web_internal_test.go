package web

import (
  "bytes"
  "net/http"
  "net/http/httptest"
  "net/url"
  "strings"
  "testing"

  "github.com/gnames/gnfmt"
  "github.com/gnames/gnlib/ent/gnvers"
  "github.com/gnames/gnparser"
  "github.com/gnames/gnparser/ent/parsed"
  "github.com/labstack/echo/v4"
  "github.com/stretchr/testify/assert"
)

func handlerGET(path string) (echo.Context, *httptest.ResponseRecorder) {
  req := httptest.NewRequest(http.MethodGet, path, nil)
  rec := httptest.NewRecorder()
  e := echo.New()
  e.Renderer, _ = NewTemplate()
  c := e.NewContext(req, rec)
  return c, rec
}

func TestHome(t *testing.T) {
  var err error
  cfg := gnparser.NewConfig(gnparser.OptFormat("compact"))
  gnp := gnparser.New(cfg)
  gnps := NewGNparserService(gnp, 0)

  req := httptest.NewRequest(http.MethodGet, "/", nil)
  rec := httptest.NewRecorder()
  e := echo.New()
  c := e.NewContext(req, rec)
  e.Renderer, err = NewTemplate()
  assert.Nil(t, err)

  assert.Nil(t, homePOST(gnps)(c))
  assert.Equal(t, rec.Code, http.StatusFound)
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
  enc := gnfmt.GNjson{}
  var response gnvers.Version
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

  enc := gnfmt.GNjson{}
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

func TestParseParamsGET(t *testing.T) {
  cfg := gnparser.NewConfig(gnparser.OptFormat("compact"))
  gnp := gnparser.New(cfg)
  gnps := NewGNparserService(gnp, 0)

  name := url.QueryEscape("Bubo bubo")
  tests := []struct {
    csv, det, startsWith, pattern string
    contains                      bool
  }{
    {"true", "false", "Id", "[", false},
    {"true", "true", "Id", "[", false},
    {"false", "false", "[", "details", false},
    {"false", "true", "[", "details", true},
  }

  for _, v := range tests {
    e := echo.New()
    q := make(url.Values)
    q.Set("csv", v.csv)
    q.Set("with_details", v.det)
    req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetPath("/:names")
    c.SetParamNames("names")
    c.SetParamValues(name)

    assert.Nil(t, parseNamesGET(gnps)(c))

    body := rec.Body.String()
    assert.True(t, strings.HasPrefix(body, v.startsWith))
    if v.contains {
      assert.True(t, strings.Contains(body, v.pattern))
    } else {
      assert.False(t, strings.HasPrefix(body, v.pattern))
    }
  }
}

func TestParsePOST(t *testing.T) {
  cfg := gnparser.NewConfig(gnparser.OptFormat("compact"))
  gnp := gnparser.New(cfg)
  gnps := NewGNparserService(gnp, 0)

  var response []parsed.Parsed
  names := []string{
    "Not name", "Bubo bubo", "Leptochloöpsis virgata", 
    "Pomatomus", "Pardosa moesta",
    "Plantago major var major", 
    "Cytospora ribis mitovirus 2",
    "A-shaped rods", "Alb. alba",
    "Pisonia grandis", "Acacia vestita may",
    "Sarracenia flava 'Maxima'",
  }
  params := inputREST{
    Names:         	    names,
    CSV:           	    false,
    WithDetails:   	    false,
    WithCultivars: 	    true,
    PreserveDiaereses:  true,
  }
  reqBody, err := gnfmt.GNjson{}.Encode(params)
  assert.Nil(t, err)
  r := bytes.NewReader(reqBody)
  req := httptest.NewRequest(http.MethodPost, "/api/v1", r)
  req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
  rec := httptest.NewRecorder()
  e := echo.New()
  c := e.NewContext(req, rec)

  assert.Nil(t, parseNamesPOST(gnps)(c))

  enc := gnfmt.GNjson{}
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
    case 2:
      assert.Equal(t, v.Verbatim, "Leptochloöpsis virgata", v.Verbatim)
      assert.Equal(t, v.Canonical.Simple, "Leptochloöpsis virgata")
    case 11:
      assert.Equal(t, v.Normalized, "Sarracenia flava ‘Maxima’")
      assert.Equal(t, v.Cardinality, 3)
    }

  }

  params = inputREST{
    Names:             names,
    CSV:               true,
    WithDetails:       false,
    WithCultivars:     false,
    PreserveDiaereses: false,
  }
  reqBody, err = gnfmt.GNjson{}.Encode(params)
  r = bytes.NewReader(reqBody)
  req = httptest.NewRequest(http.MethodPost, "/api/v1", r)
  req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
  rec = httptest.NewRecorder()
  c = e.NewContext(req, rec)
  assert.Nil(t, parseNamesPOST(gnps)(c))
  assert.True(t, strings.HasPrefix(rec.Body.String(), "Id"))
}
