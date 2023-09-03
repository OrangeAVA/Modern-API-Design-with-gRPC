package book_test

import (
	"testing"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/examples/book"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewBook(t *testing.T) {
	b := book.NewBook("1234", "Ava Orange")

	assert.Equal(t, "1234", b.GetIsbn())
	assert.Equal(t, "Ava Orange", b.GetPublisher())
}

func Test_SeriallizeBookInfoInHexaDecimal(t *testing.T) {
	hex, err := book.SeriallizeBookInfoInHexaDecimal("1234", "Ava Orange")

	require.NoError(t, err)
	assert.Equal(t, "0a0431323334120a417661204f72616e6765", hex)
}

func Test_NewBookRequest(t *testing.T) {
	b := book.NewBookRequest(5)

	assert.Equal(t, int32(5), b.GetId())
}

func Test_SeriallizeBookRequestInHexaDecimal(t *testing.T) {
	hex, err := book.SeriallizeBookRequestInHexaDecimal(5)

	require.NoError(t, err)
	assert.Equal(t, "0805", hex)
}

func Test_NewBookWithTitleAndAvailability(t *testing.T) {
	b := book.NewBookWithTitleAndAvailability(123, "The Great Gatsby", true)

	assert.Equal(t, int32(123), b.GetId())
	assert.Equal(t, "The Great Gatsby", b.GetTitle())
	assert.True(t, b.GetAvailable())
}
