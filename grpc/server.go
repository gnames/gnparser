package grpc

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"gitlab.com/gogna/gnparser"

	"gitlab.com/gogna/gnparser/dict"
	"gitlab.com/gogna/gnparser/output"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

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
	opts := []gnparser.Option{gnparser.WorkersNum(gnps.WorkersNum)}
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

func (gnps gnparserServer) ParseInOrder(stream GNparser_ParseInOrderServer) error {
	gnp := gnparser.NewGNparser()
	firstRecord := true

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		switch c := in.Content.(type) {
		case *Input_Name:
			if firstRecord {
				firstRecord = false
			}
			res, err := gnp.ParseAndFormat(c.Name)
			strError := ""
			if err != nil {
				strError = err.Error()
			}
			out := &Output{Value: res, Error: strError}
			err = stream.Send(out)
			if err != nil {
				return err
			}
		case *Input_Format:
			if firstRecord {
				firstRecord = false
				f := c.Format
				gnp = gnparser.NewGNparser(gnparser.Format(strFormat(f)))
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

func processStream(stream GNparser_ParseServer,
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
