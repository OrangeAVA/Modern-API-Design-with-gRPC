package lookaside

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/loadbalancing/internal/app/lookaside/router"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/loadbalancing/internal/pkg/kubernetes"
	proto "github.com/HiteshRepo/Modern-API-Design-with-gRPC/loadbalancing/internal/pkg/proto/lookaside"
	"google.golang.org/grpc"
	k8s "k8s.io/client-go/kubernetes"
)

const RefreshInterval = 60

type App struct {
	routers         map[string]*router.Router
	refreshInterval float64
	client          *k8s.Clientset
	proto.UnimplementedLookasideServer
}

func NewLookasideApp() *App {
	return &App{
		routers:         make(map[string]*router.Router),
		refreshInterval: RefreshInterval,
	}
}

func (a *App) Start() {
	fmt.Println("Starting look aside load balancer")

	port := "50055"
	if len(os.Getenv("LB_PORT")) > 0 {
		port = os.Getenv("LB_PORT")
	}

	servAddr := fmt.Sprintf("0.0.0.0:%s", port)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	proto.RegisterLookasideServer(s, a)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *App) Shutdown() {

}

func (a *App) Resolve(ctx context.Context, input *proto.Request) (*proto.Response, error) {
	fmt.Println("got request")

	var (
		address string
		err     error
	)

	if r, ok := a.routers[input.Service]; !ok || r.NeedsRefresh() {
		addresses, err := a.refreshAddresses(input.Service, input.Namespace)
		if err != nil {
			return nil, err
		}

		a.routers[input.Service] = &router.Router{Addresses: addresses, LastRefresh: time.Now(), RefreshIntervalInSeconds: a.refreshInterval}
	}

	switch input.Router {
	case proto.Request_RANDOM:
		address, err = a.routers[input.Service].ResolveRandom()
	case proto.Request_ROUNDROBIN:
		address, err = a.routers[input.Service].ResolveRoundRobin()
	}

	if err != nil {
		return &proto.Response{Address: ""}, err
	}

	return &proto.Response{Address: address}, nil
}

func (a *App) refreshAddresses(service, namespace string) ([]string, error) {
	var err error

	if a.client == nil {
		a.client, err = kubernetes.GetClusterClient()
		if err != nil {
			fmt.Printf("error while creating client: %s\n", err.Error())
			return nil, err
		}
	}

	addresses := kubernetes.GetEndpointsOfService(a.client, service, namespace)

	return addresses, nil
}
