package grpc

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"

	"gitlab.com/gogna/gnparser"

	"gitlab.com/gogna/gnparser/dict"
	"gitlab.com/gogna/gnparser/output"
	"gitlab.com/gogna/gnparser/preprocess"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type parseStream interface {
	Send(*Output) error
	Recv() (*Input, error)
	grpc.ServerStream
}

type cleanStream interface {
	Send(*Cleaned) error
	Recv() (*Input, error)
	grpc.ServerStream
}

type gnparserServer struct {
	WorkersNum int
}

func (gnparserServer) Ver(ctx context.Context,
	v *Void) (*Version, error) {
	ver := output.Version
	build := output.Build
	return &Version{Value: ver, BuildTime: build}, nil
}

func (gnps gnparserServer) Parse(stream GNparser_ParseServer) error {
	wn := gnps.WorkersNum
	return gnps.parse(stream, wn)
}

func (gnps gnparserServer) ParseInOrder(stream GNparser_ParseInOrderServer) error {
	return gnps.parse(stream, 1)
}

func (gnps gnparserServer) parse(stream parseStream, wn int) error {
	opts := []gnparser.Option{gnparser.WorkersNum(wn)}
	gnp := gnparser.NewGNparser(opts...)
	inCh := make(chan string)
	outCh := make(chan *gnparser.ParseResult)
	var wg sync.WaitGroup
	wg.Add(1)
	firstRecord := true

	go processStream(stream, outCh, &wg)

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			close(inCh)
			wg.Wait()
			return nil
		}
		if err != nil {
			return err
		}
		switch c := in.Content.(type) {
		case *Input_Name:
			if firstRecord {
				firstRecord = false
				go gnp.ParseStream(inCh, outCh, opts...)
			}
			inCh <- c.Name
		case *Input_Format:
			if firstRecord {
				firstRecord = false
				f := c.Format
				opts = append(opts, gnparser.Format(strFormat(f)))
				go gnp.ParseStream(inCh, outCh, opts...)
			}
		}
	}
}

func strFormat(f Format) string {
	switch f {
	case Format_Compact:
		return "compact"
	case Format_Pretty:
		return "pretty"
	case Format_Simple:
		return "simple"
	case Format_Debug:
		return "debug"
	}
	return "compact"
}

func processStream(stream parseStream,
	outCh <-chan *gnparser.ParseResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for r := range outCh {
		errStr := ""
		if r.Error != nil {
			errStr = r.Error.Error()
		}
		out := &Output{
			Value: r.Output,
			Error: errStr,
		}
		stream.Send(out)
	}
}

func (gnps gnparserServer) Clean(stream GNparser_CleanServer) error {
	return gnps.clean(stream, gnps.WorkersNum)
}

func (gnps gnparserServer) CleanInOrder(stream GNparser_CleanInOrderServer) error {
	return gnps.clean(stream, 1)
}

func (gnps gnparserServer) clean(stream cleanStream, wn int) error {
	inCh := make(chan string)
	outCh := make(chan string)
	var wg sync.WaitGroup
	firstRecord := true
	wg.Add(1)
	go processCleaningStream(stream, outCh, &wg)

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			close(inCh)
			wg.Wait()
			return nil
		}
		if err != nil {
			return err
		}
		switch c := in.Content.(type) {
		case *Input_Name:
			if firstRecord {
				firstRecord = false
				go preprocess.CleanupStream(inCh, outCh, wn)
			}
			inCh <- c.Name
		}
	}
}

func processCleaningStream(stream cleanStream,
	outCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for r := range outCh {
		res := strings.Split(r, "|")
		out := &Cleaned{
			Input:  res[0],
			Output: res[1],
		}
		stream.Send(out)
	}
}

var dictionary *dict.Dictionary

func Run(port int, workersNum int) {
	gnps := gnparserServer{WorkersNum: workersNum}
	srv := grpc.NewServer()
	dictionary = dict.LoadDictionary()
	RegisterGNparserServer(srv, gnps)
	portVal := fmt.Sprintf(":%d", port)
	l, err := net.Listen("tcp", portVal)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", portVal, err)
	}
	log.Fatal(srv.Serve(l))
}
