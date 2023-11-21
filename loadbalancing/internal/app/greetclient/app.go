package greetclient

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/loadbalancing/internal/pkg/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/xds"
)

type App struct {
	client proto.GreetServiceClient
	conn   *grpc.ClientConn
}

type GreetingRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (a *App) Start() {
	err := a.setupGreetClient()
	if err != nil {
		log.Printf("could not setup greet client: %v", err)
		return
	}

	fmt.Println("Starting gateway")

	router := gin.Default()
	router.POST("/greet", a.greet)

	port := "9091"
	if len(os.Getenv("GATEWAY_PORT")) > 0 {
		port = os.Getenv("GATEWAY_PORT")
	}

	fmt.Println("Starting REST Gateway")

	router.Run(fmt.Sprintf("0.0.0.0:%s", port))
}

func (a *App) setupGreetClient() error {
	var err error

	fmt.Println("Starting greet client")

	opts := grpc.WithInsecure()
	serverHost := ""
	serverPort := "50051"

	if port, ok := os.LookupEnv("GRPC_SERVER_PORT"); ok {
		serverPort = port
	}

	if host, ok := os.LookupEnv("GRPC_SERVER_HOST"); ok {
		serverHost = host
	} else {
		log.Println("GRPC_SERVER_HOST not provided")
		return errors.New("GRPC_SERVER_HOST not provided")
	}

	servAddr := fmt.Sprintf("xds:///%s:%s", serverHost, serverPort)

	fmt.Println("dialing", servAddr)

	a.conn, err = grpc.Dial(
		servAddr,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithBlock(),
		opts,
	)
	if err != nil {
		log.Printf("could not connect: %v", err)
		return err
	}

	a.client = proto.NewGreetServiceClient(a.conn)

	return nil
}

func (a *App) Shutdown() {
	a.conn.Close()
}

func (a *App) doUnary(firstName, lastName string) string {
	req := &proto.GreetingRequest{
		Greeting: &proto.Greeting{
			FirstName: firstName,
			LastName:  lastName,
		},
	}
	res, err := a.client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet rpc : %v", err)
	}
	return fmt.Sprintf("reponse from Greet rpc: %v", res.Result)
}

func (a *App) greet(c *gin.Context) {
	fmt.Println("got request - REST Gateway")

	var gr GreetingRequest
	if err := c.BindJSON(&gr); err != nil {
		return
	}

	resp := a.doUnary(gr.FirstName, gr.LastName)
	c.String(http.StatusOK, resp)
}
