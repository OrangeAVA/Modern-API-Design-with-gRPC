package bookserver_test

import (
	"context"
	"log"
	"net"
	"testing"

	bookserver "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-8/book-server"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-8/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type BookServerTestSuite struct {
	*suite.Suite

	listener *bufconn.Listener
	server   *grpc.Server

	app *bookserver.App
}

func (suite *BookServerTestSuite) SetupSuite() {
	suite.listener = bufconn.Listen(1024 * 1024)
	suite.server = grpc.NewServer()

	suite.app = bookserver.NewApp()
	proto.RegisterBookServiceServer(suite.server, suite.app)
}

func TestBookServerTestSuite(t *testing.T) {
	suite.Run(t, &BookServerTestSuite{
		Suite: new(suite.Suite),
	})
}

func (suite *BookServerTestSuite) TeardownSuite() {
	if suite.listener != nil {
		suite.listener.Close()
	}

	if suite.server != nil {
		suite.server.Stop()
	}
}

func (suite *BookServerTestSuite) TestBookService_ListBooks() {
	// Start the gRPC server in a goroutine

	go func() {
		if err := suite.server.Serve(suite.listener); err != nil {
			log.Fatalf("server failed to start %v", err)
		}
	}()

	t := suite.T()

	// Set up a gRPC client connection to the in-memory listener
	conn, clientCleanup := suite.setupGRPCClient()
	defer clientCleanup()

	// Make a ListBooks gRPC call
	resp, err := suite.callListBooksRPC(conn)
	require.NoError(t, err, "failed to ListBooks")

	// Assert the response
	expectedBooks := suite.app.GetAllBooks()
	assert.Equal(t, expectedBooks, resp.Books, "unexpected list of books")

}

// setupGRPCClient creates a gRPC client connection and returns a cleanup function
func (suite *BookServerTestSuite) setupGRPCClient() (*grpc.ClientConn, func()) {
	t := suite.T()

	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return suite.listener.Dial()
		}),
		grpc.WithInsecure(),
	)
	require.NoError(t, err, "failed to dial bufnet")

	// Return a cleanup function to close the client connection
	return conn, func() {
		conn.Close()
	}
}

// callListBooksRPC makes a ListBooks gRPC call and returns the response
func (suite *BookServerTestSuite) callListBooksRPC(conn *grpc.ClientConn) (*proto.ListBooksRespose, error) {
	client := proto.NewBookServiceClient(conn)
	empty := &proto.Empty{}
	return client.ListBooks(context.Background(), empty)
}
