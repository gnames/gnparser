package web

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/shurcooL/httpfs/html/vfstemplate"
	"gitlab.com/gogna/gnparser"
	"gitlab.com/gogna/gnparser/fs"
	"gitlab.com/gogna/gnparser/output"
)

type Data struct {
	Input    string
	Parsed   []string
	HomePage bool
	Version  string
}

func NewData() *Data {
	return &Data{Version: output.Version}
}

func home(w http.ResponseWriter, r *http.Request) {
	data := NewData()
	data.HomePage = true
	params := r.URL.Query()
	if txt, ok := params["q"]; ok && len(txt) > 0 {
		fmt.Println(txt)
		names := namesFromText(txt[0])
		data.Input = txt[0]
		data.HomePage = true
		data.Parsed = parseForWeb(names)
	}
	var t *template.Template
	t, err := vfstemplate.ParseFiles(fs.Files, t, "templates/layout.html",
		"templates/home.html")
	if err != nil {
		fmt.Fprintf(w, "<html><body><h1>Error</h2><p>%s</p></body></html>",
			err.Error())
		return
	}
	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		fmt.Fprintf(w, "<html><body><h1>Error</h2><p>%s</p></body></html>",
			err.Error())
		return
	}
}

func docAPI(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	t, err := vfstemplate.ParseFiles(fs.Files, t, "templates/layout.html",
		"templates/doc_api.html")
	if err != nil {
		fmt.Fprintf(w, "<html><body><h1>Error</h2><p>%s</p></body></html>",
			err.Error())
		return
	}
	data := NewData()
	t.ExecuteTemplate(w, "layout", data)
}

func parseForWeb(names []string) []string {
	parsed := make([]string, len(names))
	opts := []gnparser.Option{gnparser.OptFormat("pretty")}
	gnp := gnparser.NewGNparser(opts...)
	for i, v := range names {
		json, err := gnp.ParseAndFormat(v)
		if err != nil {
			parsed[i] = err.Error()
			continue
		}
		parsed[i] = json
	}
	return parsed
}

func namesFromText(txt string) []string {
	var names []string
	names = strings.Split(txt, "|")
	if len(names) > 1 {
		return names
	}
	names = names[0:0]
	strs := strings.Split(txt, "\n")
	for _, v := range strs {
		v = strings.TrimRight(v, "\r ")
		fmt.Printf("'%s'", v)
		if v != "" {
			names = append(names, v)
		}
	}
	return names
}
