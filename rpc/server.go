package rpc

import (
	"fmt"
	"log"
	"net"
	"sync"

	"gitlab.com/gogna/gnparser"
	"gitlab.com/gogna/gnparser/dict"
	"gitlab.com/gogna/gnparser/output"
	"gitlab.com/gogna/gnparser/pb"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type gnparserServer struct {
	MaxWorkersNum int
}

// Ver takes an empty argument and returns the version of the gnparser as well
// as its build timestamp.
func (gnparserServer) Ver(ctx context.Context,
	v *pb.Void) (*pb.Version, error) {
	ver := output.Version
	build := output.Build
	return &pb.Version{Value: ver, BuildTime: build}, nil
}

// ParseArray takes an input with an array of name-strings and returns an
// output with an array of protobuf objects with parsing results.
// The order of elements in output is preserved the same as in input.
// The name-strings are getting stripped from html elements, and the
// 'Verbatim' value in such cases would differ fromo the 'raw' input.
func (gnps gnparserServer) ParseArray(ctx context.Context,
	ia *pb.InputArray) (*pb.OutputArray, error) {
	arrayMax := 10000
	if len(ia.Names) > arrayMax {
		err := fmt.Errorf("keep input smaller than %d entries", arrayMax)
		return nil, err
	}
	if len(ia.Names) == 0 {
		err := fmt.Errorf("empty input")
		return nil, err
	}

	parsed := gnps.parseArray(ia)
	oa := &pb.OutputArray{Output: parsed}
	return oa, nil
}

var dictionary *dict.Dictionary

// Run takes a port number to run as well as the number of workers to support.
func Run(port int, workersNum int) {
	gnps := gnparserServer{MaxWorkersNum: workersNum}
	srv := grpc.NewServer()
	dictionary = dict.LoadDictionary()
	pb.RegisterGNparserServer(srv, gnps)
	portVal := fmt.Sprintf(":%d", port)
	l, err := net.Listen("tcp", portVal)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", portVal, err)
	}
	log.Fatal(srv.Serve(l))
}

type parseArrayOutput struct {
	inputName    string
	outputParsed *pb.Parsed
}

func (gnps gnparserServer) parseArray(ia *pb.InputArray) []*pb.Parsed {
	jobs := int(ia.JobsNumber)
	if jobs == 0 || gnps.MaxWorkersNum < jobs {
		jobs = gnps.MaxWorkersNum
	}
	skipClean := ia.SkipCleaning
	log.Printf("Processing %d names using %d jobs", len(ia.Names), jobs)
	resMap := make(map[string]*pb.Parsed)
	inCh := make(chan string)
	outCh := make(chan *parseArrayOutput)
	var parseWg sync.WaitGroup
	var processWg sync.WaitGroup
	processWg.Add(1)
	parseWg.Add(jobs)

	for i := 0; i < jobs; i++ {
		go parseWorker(inCh, outCh, skipClean, &parseWg)
	}
	go processParseArray(outCh, &processWg, resMap)
	for _, v := range ia.Names {
		inCh <- v
	}
	close(inCh)
	parseWg.Wait()
	close(outCh)
	processWg.Wait()

	res := make([]*pb.Parsed, len(ia.Names))
	for i, v := range ia.Names {
		if pv, ok := resMap[v]; ok {
			res[i] = pv
		}
	}
	return res
}

func parseWorker(inCh <-chan string, outCh chan<- *parseArrayOutput,
	skipClean bool, wg *sync.WaitGroup) {
	defer wg.Done()
	opts := []gnparser.Option{gnparser.OptRemoveHTML(!skipClean)}
	gnp := gnparser.NewGNparser(opts...)
	for v := range inCh {
		res := gnp.ParseToObject(v)
		outCh <- &parseArrayOutput{inputName: v, outputParsed: res}
	}
}

func processParseArray(ch <-chan *parseArrayOutput, wg *sync.WaitGroup,
	resMap map[string]*pb.Parsed) {
	defer wg.Done()
	for v := range ch {
		resMap[v.inputName] = v.outputParsed
	}
}
