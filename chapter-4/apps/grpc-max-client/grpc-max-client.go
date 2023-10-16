package grpcmaxclient

import (
	"context"
	"flag"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/proto"
	"google.golang.org/grpc"
)

var (
	numbers string
)

func init() {
	flag.StringVar(&numbers, "numbers", "", "provide comma separated numbers to compute max")
}

type App struct {
	clientConn *grpc.ClientConn
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	var err error
	flag.Parse()

	servAddr := "localhost:50053"

	opts := grpc.WithInsecure()
	a.clientConn, err = grpc.Dial(servAddr, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	maxServiceClient := proto.NewMaxServiceClient(a.clientConn)

	numberArr := make([]int, 0)

	strNumbers := strings.Split(strings.TrimSpace(numbers), ",")
	for _, strNum := range strNumbers {
		intNum, _ := strconv.Atoi(strings.TrimSpace(strNum))
		numberArr = append(numberArr, intNum)
	}

	doBidirectionalStreaming(maxServiceClient, numberArr)
}

func (a *App) Shutdown() {
	a.clientConn.Close()
}

func doBidirectionalStreaming(c proto.MaxServiceClient, numberArr []int) {
	requests := make([]*proto.MaxRequest, 0)

	for _, n := range numberArr {
		req := &proto.MaxRequest{Number: int32(n)}
		requests = append(requests, req)
	}

	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("error while invoking FindMax rpc: %v", err)
	}

	waitChan := make(chan struct{})

	go func(reqs []*proto.MaxRequest) {
		for _, req := range reqs {
			log.Printf("sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}(requests)

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving from FindMax rpc response: %v", err)
				break
			}
			log.Printf("response received: %v\n", res.GetMax())
		}
		close(waitChan)
	}()

	<-waitChan
}
