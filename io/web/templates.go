package web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"path"
	"strings"

	vlib "github.com/gnames/gnlib/ent/verifier"
	"github.com/labstack/echo/v4"
)

//go:embed templates
var tmpls embed.FS

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

func NewTemplate() (*echoTempl, error) {
	t, err := parseFiles()
	if err != nil {
		return nil, fmt.Errorf("cannot parse file %w", err)
	}
	return &echoTempl{t}, nil
}

func parseFiles() (*template.Template, error) {
	var err error
	var t *template.Template

	var filenames []string
	dir := "templates"
	entries, _ := tmpls.ReadDir(dir)
	for i := range entries {
		if entries[i].Type().IsRegular() {
			filenames = append(
				filenames,
				fmt.Sprintf("%s/%s", dir, entries[i].Name()),
			)
		}
	}

	for _, filename := range filenames {
		name := path.Base(filename)
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		addFuncs(tmpl)
		_, err = tmpl.ParseFS(tmpls, filename)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func addFuncs(tmpl *template.Template) {
	tmpl.Funcs(template.FuncMap{
		"isEven": func(i int) bool {
			return i%2 == 0
		},
		"classification": func(pathStr, rankStr string) string {
			if pathStr == "" {
				return ""
			}
			paths := strings.Split(pathStr, "|")
			var ranks []string
			if rankStr != "" {
				ranks = strings.Split(rankStr, "|")
			}

			res := make([]string, len(paths))
			for i := range paths {
				path := strings.TrimSpace(paths[i])
				if len(ranks) == len(paths) {
					rank := strings.TrimSpace(ranks[i])
					if rank != "" {
						path = fmt.Sprintf("%s (%s)", path, rank)
					}
				}
				res[i] = path
			}
			return strings.Join(res, " >> ")
		},
		"matchType": func(mt vlib.MatchTypeValue, ed int) template.HTML {
			var res string
			clr := map[string]string{
				"green":  "#080",
				"yellow": "#a80",
				"red":    "#800",
			}
			switch mt {
			case vlib.Exact:
				res = fmt.Sprintf("<span style='color: %s'>%s match by canonical form</span>", clr["green"], mt)
			case vlib.NoMatch:
				res = fmt.Sprintf("<span style='color: %s'>%s</span>", clr["red"], mt)
			case vlib.Fuzzy, vlib.PartialFuzzy:
				res = fmt.Sprintf("<span style='color: %s'>%s match, edit distance: %d</span>", clr["yellow"], mt, ed)
			default:
				res = fmt.Sprintf("<span style='color: %s'>%s match</span>", clr["yellow"], mt)
			}
			return template.HTML(res)
		},
	})
}
