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
	cfg := gnparser.NewConfig(gnparser.OptFormat(gnfmt.CompactJSON))
	gnp := gnparser.New(cfg)
	gnps := NewGNparserService(gnp, 0)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	e.Renderer, err = NewTemplate()
	assert.Nil(t, err)

	assert.Nil(t, homePOST(gnps)(c))
	assert.Equal(t, http.StatusFound, rec.Code)
}

// func TestDocAPI(t *testing.T) {
// 	c, rec := handlerGET("/doc/api")
// 	assert.Nil(t, docAPI()(c))
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Contains(t, rec.Body.String(), "Application Programming Interface")
// }

func TestInfo(t *testing.T) {
	c, rec := handlerGET("/")

	assert.Nil(t, info()(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "OpenAPI")
}

func TestPing(t *testing.T) {
	cfg := gnparser.NewConfig(gnparser.OptFormat(gnfmt.CompactJSON))
	gnp := gnparser.New(cfg)
	gnps := NewGNparserService(gnp, 0)
	c, rec := handlerGET("/ping")

	assert.Nil(t, ping(gnps)(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "pong", rec.Body.String())
}

func TestVer(t *testing.T) {
	cfg := gnparser.NewConfig(gnparser.OptFormat(gnfmt.CompactJSON))
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
	cfg := gnparser.NewConfig(gnparser.OptFormat(gnfmt.CompactJSON))
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

	assert.Equal(t, len(names), len(response))
	for i, v := range response {
		switch i {
		case 0:
			assert.Equal(t, "Not name", v.Verbatim, v.Verbatim)
			assert.False(t, v.Parsed, v.Verbatim)
		case 1:
			assert.Equal(t, "Bubo bubo", v.Verbatim, v.Verbatim)
			assert.Equal(t, "Bubo bubo", v.Canonical.Simple)
		}
	}
}

func TestParseParamsGET(t *testing.T) {
	cfg := gnparser.NewConfig(gnparser.OptFormat(gnfmt.CompactJSON))
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
	cfg := gnparser.NewConfig(gnparser.OptFormat(gnfmt.CompactJSON))
	gnp := gnparser.New(cfg)
	gnps := NewGNparserService(gnp, 0)

	var response []parsed.Parsed
	names := []string{
		"Not name", "Bubo bubo", "Leptochloöpsis virgata",
		"Solanum linnaeanum Hepper & P.-M. L. Jaeger",
		"Pomatomus", "Pardosa moesta",
		"Plantago major var major",
		"Cytospora ribis mitovirus 2",
		"A-shaped rods", "Alb. alba",
		"Pisonia grandis", "Acacia vestita may",
		"Sarracenia flava 'Maxima'",
	}
	params := inputREST{
		Names:             names,
		CSV:               false,
		WithDetails:       false,
		WithCultivars:     true,
		PreserveDiaereses: true,
		CompactAuthors:    true,
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

	assert.Equal(t, len(names), len(response))
	for i, v := range response {
		switch i {
		case 0:
			assert.Equal(t, "Not name", v.Verbatim, v.Verbatim)
			assert.False(t, v.Parsed, v.Verbatim)
		case 1:
			assert.Equal(t, "Bubo bubo", v.Verbatim, v.Verbatim)
			assert.Equal(t, "Bubo bubo", v.Canonical.Simple)
		case 2:
			assert.Equal(t, "Leptochloöpsis virgata", v.Verbatim, v.Verbatim)
			assert.Equal(t, "Leptochloöpsis virgata", v.Canonical.Simple)
		case 3:
			assert.Equal(t, "Solanum linnaeanum Hepper & P.-M.L.Jaeger", v.Normalized)
			assert.Equal(t, "Hepper & P.-M.L.Jaeger", v.Authorship.Normalized)
		case 12:
			assert.Equal(t, "Sarracenia flava ‘Maxima’", v.Normalized)
			assert.Equal(t, 3, v.Cardinality)
		}

	}

	params = inputREST{
		Names:             names,
		CSV:               true,
		WithDetails:       false,
		WithCultivars:     false,
		PreserveDiaereses: false,
		CompactAuthors:    false,
	}
	reqBody, err = gnfmt.GNjson{}.Encode(params)
	r = bytes.NewReader(reqBody)
	req = httptest.NewRequest(http.MethodPost, "/api/v1", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	assert.Nil(t, err)
	assert.Nil(t, parseNamesPOST(gnps)(c))
	assert.True(t, strings.HasPrefix(rec.Body.String(), "Id"))
}

func TestParsePOST_FlatOutput(t *testing.T) {
	cfg := gnparser.NewConfig(gnparser.OptFormat(gnfmt.CompactJSON))
	gnp := gnparser.New(cfg)
	gnps := NewGNparserService(gnp, 0)

	var response any
	names := []string{"Homo sapiens Linnaeus, 1758"}

	// Test with FlatOutput = true (flattened JSON)
	params := inputREST{
		Names:       names,
		CSV:         false,
		WithDetails: false,
		FlatOutput:  true,
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

	responseSlice, ok := response.([]any)
	assert.True(t, ok)
	assert.Equal(t, 1, len(responseSlice))

	body := rec.Body.String()
	// Flattened output should not contain nested "canonical" object
	assert.NotContains(t, body, `"canonical":`)
	// Should contain flat fields
	assert.Contains(t, body, `"canonicalSimple"`)
	assert.Contains(t, body, `"authorship"`)

	// Test with FlatOutput = false (nested JSON)
	params = inputREST{
		Names:       names,
		CSV:         false,
		WithDetails: false,
		FlatOutput:  false,
	}
	reqBody, err = gnfmt.GNjson{}.Encode(params)
	assert.Nil(t, err)
	r = bytes.NewReader(reqBody)
	req = httptest.NewRequest(http.MethodPost, "/api/v1", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	assert.Nil(t, parseNamesPOST(gnps)(c))

	body = rec.Body.String()
	// Nested output should contain "canonical" object
	assert.Contains(t, body, `"canonical":`)
	assert.NotContains(t, body, `"canonicalSimple"`)
}

func TestParseGET_FlatOutput(t *testing.T) {
	cfg := gnparser.NewConfig(gnparser.OptFormat(gnfmt.CompactJSON))
	gnp := gnparser.New(cfg)
	gnps := NewGNparserService(gnp, 0)

	name := url.QueryEscape("Homo sapiens")

	// Test with flatten=true
	e := echo.New()
	q := make(url.Values)
	q.Set("flatten", "true")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:names")
	c.SetParamNames("names")
	c.SetParamValues(name)

	assert.Nil(t, parseNamesGET(gnps)(c))

	body := rec.Body.String()
	assert.NotContains(t, body, `"canonical":`)
	assert.Contains(t, body, `"canonicalSimple"`)

	// Test with flatten=false (default)
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/:names")
	c.SetParamNames("names")
	c.SetParamValues(name)

	assert.Nil(t, parseNamesGET(gnps)(c))

	body = rec.Body.String()
	assert.Contains(t, body, `"canonical":`)
	assert.NotContains(t, body, `"canonicalSimple"`)
}

func TestParsePOST_CSV_WithFlatOutput(t *testing.T) {
	cfg := gnparser.NewConfig(gnparser.OptFormat(gnfmt.CompactJSON))
	gnp := gnparser.New(cfg)
	gnps := NewGNparserService(gnp, 0)

	names := []string{"Homo sapiens Linnaeus, 1758", "Bubo bubo"}

	// Test CSV output without details (simple 10 fields)
	params := inputREST{
		Names:       names,
		CSV:         true,
		WithDetails: false,
		FlatOutput:  true,
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

	body := rec.Body.String()
	// CSV without details should have simple header (10 fields)
	assert.Contains(t, body, "Id,Verbatim,Cardinality")
	assert.Contains(t, body, "NomCodeSetting")
	// Should NOT include extended fields when WithDetails=false
	assert.NotContains(t, body, "Parsed,")
	assert.NotContains(t, body, ",Authors,")
	assert.NotContains(t, body, ",Genus,")
	assert.NotContains(t, body, "CultivarEpithet")

	// Test CSV with WithDetails=true (extended fields)
	params = inputREST{
		Names:       names,
		CSV:         true,
		WithDetails: true,
		FlatOutput:  false, // FlatOutput is ignored for CSV
	}
	reqBody, err = gnfmt.GNjson{}.Encode(params)
	assert.Nil(t, err)
	r = bytes.NewReader(reqBody)
	req = httptest.NewRequest(http.MethodPost, "/api/v1", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	assert.Nil(t, parseNamesPOST(gnps)(c))

	body = rec.Body.String()
	// Should include all extended fields when WithDetails=true
	assert.Contains(t, body, "Parsed,")
	assert.Contains(t, body, ",Authors,")
	assert.Contains(t, body, ",Genus,")
	assert.Contains(t, body, ",Species,")
	assert.Contains(t, body, ",Infraspecies")
	assert.Contains(t, body, "CultivarEpithet")
}
