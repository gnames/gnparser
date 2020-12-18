package web

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/mux"

	"github.com/gnames/gnparser"
	jsoniter "github.com/json-iterator/go"
)

func apiEmptyRequest(w http.ResponseWriter, r *http.Request) {
	res := `{"error": "Unrecognized request"}`
	fmt.Fprint(w, res)
}

func apiGetParse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	namesPipe, ok := params["q"]
	if !ok {
		fmt.Fprint(w, "[]\n")
		return
	}
	names := strings.Split(namesPipe, "|")
	parseSlice(w, names)
}

func apiPostParse(w http.ResponseWriter, r *http.Request) {
	var names []string
	_ = jsoniter.NewDecoder(r.Body).Decode(&names)
	if names == nil || len(names) == 0 {
		fmt.Fprint(w, "[]\n")
		return
	}
	parseSlice(w, names)
}

func parseSlice(w http.ResponseWriter, ns []string) {
	in := make(chan string)
	out := make(chan *gnparser.ParseResult)
	var wg sync.WaitGroup
	wg.Add(1)
	opts := []gnparser.Option{gnparser.OptFormat("compact")}
	go gnparser.ParseStream(8, in, out, opts...)
	go processResults(w, out, &wg)
	for _, v := range ns {
		in <- v
	}
	close(in)
	wg.Wait()
}

func processResults(w http.ResponseWriter, out <-chan *gnparser.ParseResult, wg *sync.WaitGroup) {
	defer wg.Done()
	var res []string
	for r := range out {
		if r.Error == nil {
			res = append(res, r.Output)
		}
	}
	fmt.Fprint(w, "[\n"+strings.Join(res, ",\n")+"]\n")
}
