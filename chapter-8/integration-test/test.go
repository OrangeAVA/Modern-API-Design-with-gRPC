package integration_test

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	bookclient "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-8/book-client"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-8/proto"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestIntegration(t *testing.T) {
	ctx := context.Background()

	// Start a gRPC server container
	serverContainer := startGRPCServerContainer(ctx, t)

	// Cleanup the container after the test
	defer func() {
		assert.NoError(t, serverContainer.Terminate(ctx))
	}()

	// Create a gRPC client connected to the server container
	client := createGRPCClient(ctx, t, serverContainer)

	// Test your gRPC client against the server
	resp, err := client.ListBooks(ctx, &proto.Empty{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

// startGRPCServerContainer starts a gRPC server container using testcontainers
func startGRPCServerContainer(ctx context.Context, t *testing.T) testcontainers.Container {
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    filepath.Join("."),
			Dockerfile: "args.Dockerfile",
		},
		ExposedPorts: []string{"50091"},
		WaitingFor:   wait.ForListeningPort("50091").WithStartupTimeout(30 * time.Second),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)

	return container
}

// createGRPCClient creates a gRPC client connected to the specified server container
func createGRPCClient(ctx context.Context, t *testing.T, serverContainer testcontainers.Container) proto.BookServiceClient {
	port, err := serverContainer.MappedPort(ctx, "50091")
	assert.NoError(t, err)

	// Use the mapped port to create a gRPC client
	client := bookclient.NewApp(fmt.Sprintf("localhost:%s", port.Port()))
	return client
}
