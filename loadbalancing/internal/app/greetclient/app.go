package greetclient

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/loadbalancing/internal/pkg/proto"
	lapb "github.com/HiteshRepo/Modern-API-Design-with-gRPC/loadbalancing/internal/pkg/proto/lookaside"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type App struct {
	greetClients       map[string]proto.GreetServiceClient
	conn               map[string]*grpc.ClientConn
	currentGreetClient proto.GreetServiceClient

	lookasideConn   *grpc.ClientConn
	lookasideClient lapb.LookasideClient
}

func NewGreetClient() *App {
	return &App{
		greetClients: make(map[string]proto.GreetServiceClient),
		conn:         make(map[string]*grpc.ClientConn),
	}
}

type GreetingRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (a *App) Start() {
	err := a.setupLookasideClient()
	if err != nil {
		log.Printf("could not setup lookaside client: %v", err)
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

func (a *App) setupGreetClient(host string) error {
	fmt.Println("Starting greet client")

	opts := grpc.WithInsecure()
	serverPort := "50051"

	if port, ok := os.LookupEnv("GRPC_SERVER_PORT"); ok {
		serverPort = port
	}

	if c, ok := a.greetClients[host]; !ok {
		servAddr := fmt.Sprintf("%s:%s", host, serverPort)

		fmt.Println("dialing greet server", servAddr)

		conn, err := grpc.Dial(
			servAddr,
			opts,
		)
		if err != nil {
			log.Printf("could not connect greet server: %v", err)
			return err
		}

		a.conn[host] = conn

		a.currentGreetClient = proto.NewGreetServiceClient(conn)
		a.greetClients[host] = a.currentGreetClient
	} else {
		a.currentGreetClient = c
	}

	return nil
}

func (a *App) setupLookasideClient() error {
	var err error

	fmt.Println("Starting lookaside client")

	opts := grpc.WithInsecure()
	serverPort := "50055"

	if port, ok := os.LookupEnv("LB_PORT"); ok {
		serverPort = port
	}

	servAddr := fmt.Sprintf("%s:%s", os.Getenv("LB_SVC_NAME"), serverPort)

	fmt.Println("dialing lookaside server", servAddr)

	a.lookasideConn, err = grpc.Dial(
		servAddr,
		opts,
	)
	if err != nil {
		log.Printf("could not connect lookaside server: %v", err)
		return err
	}

	a.lookasideClient = lapb.NewLookasideClient(a.lookasideConn)

	return nil
}

func (a *App) Shutdown() {
	for _, conn := range a.conn {
		conn.Close()
	}
}

func (a *App) doUnary(firstName, lastName string) string {
	r, err := a.lookasideClient.Resolve(context.Background(), &lapb.Request{
		Router:    lapb.Request_ROUNDROBIN,
		Service:   os.Getenv("GRPC_SVC"),
		Namespace: os.Getenv("POD_NAMESPACE"),
	})

	if err != nil {
		log.Printf("could not resolve greet server address: %v", err)
		return ""
	}

	err = a.setupGreetClient(r.GetAddress())
	if err != nil {
		log.Printf("could not setup greet client: %v", err)
		return ""
	}

	req := &proto.GreetingRequest{
		Greeting: &proto.Greeting{
			FirstName: firstName,
			LastName:  lastName,
		},
	}
	res, err := a.currentGreetClient.Greet(context.Background(), req)
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
