package grpcaverageclient

import (
	"context"
	"flag"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	numbers string
)

func init() {
	flag.StringVar(&numbers, "numbers", "", "provide comma separated numbers to compute average")
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

	servAddr := "localhost:50052"

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	a.clientConn, err = grpc.Dial(servAddr, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	averageServiceClient := proto.NewAverageServiceClient(a.clientConn)

	numberArr := make([]int, 0)

	strNumbers := strings.Split(strings.TrimSpace(numbers), ",")
	for _, strNum := range strNumbers {
		intNum, _ := strconv.Atoi(strings.TrimSpace(strNum))
		numberArr = append(numberArr, intNum)
	}

	doClientStreaming(averageServiceClient, numberArr)
}

func (a *App) Shutdown() {
	a.clientConn.Close()
}

func doClientStreaming(c proto.AverageServiceClient, numberArr []int) {
	requests := make([]*proto.AverageRequest, 0)

	for _, n := range numberArr {
		req := &proto.AverageRequest{Number: int32(n)}
		requests = append(requests, req)
	}

	stream, err := c.FindAverage(context.Background())
	if err != nil {
		log.Fatalf("error while invoking FindAverage rpc: %v", err)
	}

	for _, req := range requests {
		log.Printf("sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while reading FindAverage response: %v", err)
	}
	log.Printf("FindAverage response: %v", res)
}
