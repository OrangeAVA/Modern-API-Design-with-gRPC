package bookclient

import (
	"context"
	"testing"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-8/proto"
	mock_proto "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-8/proto/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BookClientTestSuite struct {
	*suite.Suite

	controller        *gomock.Controller
	bookServiceClient *mock_proto.MockBookServiceClient
}

func TestBookServerTestSuite(t *testing.T) {
	suite.Run(t, &BookClientTestSuite{
		Suite: new(suite.Suite),
	})
}

func (suite *BookClientTestSuite) SetupSuite() {
	suite.controller = gomock.NewController(suite.T())
	suite.bookServiceClient = mock_proto.NewMockBookServiceClient(suite.controller)
}

func (suite *BookClientTestSuite) TeardownSuite() {
	if suite.controller != nil {
		suite.controller.Finish()
	}
}

func (suite *BookClientTestSuite) TestBookClient_ListBooks() {
	t := suite.T()

	client := App{
		client: suite.bookServiceClient,
	}

	// Expected data
	expectedBooks := []*proto.Book{
		{
			Isbn:      1234,
			Name:      "book-1",
			Publisher: "publisher-1",
		},
		{
			Isbn:      1235,
			Name:      "book-2",
			Publisher: "publisher-2",
		},
	}

	// Set up the mock
	suite.bookServiceClient.EXPECT().
		ListBooks(gomock.Any(), gomock.Any()).
		Return(&proto.ListBooksRespose{Books: expectedBooks}, nil)

	// Call the ListBooks method on the client
	resp, err := client.ListBooks(context.Background(), &proto.Empty{})

	// Assert that the response and error match the expected values
	assert.NoError(t, err)
	assert.Equal(t, expectedBooks, resp.Books)
}
