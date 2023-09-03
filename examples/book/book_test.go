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
